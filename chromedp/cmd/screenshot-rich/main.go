package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var (
	ColorPrinter = struct {
		Red      string
		Green    string
		Yellow   string
		Colorize func(string, string) string
	}{
		Red:    "\033[31m",
		Green:  "\033[32m",
		Yellow: "\033[33m",
		Colorize: func(color string, text string) string {
			if runtime.GOOS == "windows" {
				return text
			}
			colorReset := "\033[0m"
			return color + text + colorReset
		},
	}
	arguments = struct {
		url               *string
		querySelector     *string
		outputPath        *string
		profileDir        *string
		windowWidth       *int64
		windowHeight      *int64
		deviceScaleFactor *float64
		fullScreenshot    *bool
		debug             *bool
		noHeadless        *bool
		saveProfile       *bool
	}{
		flag.String("u", "" /*              */, ColorPrinter.Colorize(ColorPrinter.Yellow, "[Required]")+" URL"),
		flag.String("q", "" /*              */, "Query selector. Screenshot the first matching element. ( e.g. -q=\".className#id\" )"),
		flag.String("o", `/tmp/img.png` /*  */, "Output path of screenshot"),
		flag.String("pd", `` /*             */, "Chrome profile directory to copy and cache. (e.g. -pd=\"/Users/you/Library/Application Support/Google/Chrome/Default\"). Use -s to keep cache after execution."),
		flag.Int64("wi", 1280 /*            */, "Viewport width (affects page layout, e.g. responsive design). Without -q, this is the output image width"),
		flag.Int64("he", 860 /*             */, "Viewport height (affects page layout, e.g. responsive design). Without -q, this is the output image height"),
		flag.Float64("ds", 2.0 /*          */, "Device scale factor (2.0 = Retina)"),
		flag.Bool("f", false /*             */, "\nEnable full screenshot mode"),
		flag.Bool("d", false /*             */, "\nEnable debug mode"),
		flag.Bool("n", false /*             */, "\nDisable headless mode"),
		flag.Bool("s", false /*             */, "\nSave cached profile (do not delete after execution)"),
	}
)

// [ Usage ]
// go run cmd/screenshot/main.go -u="https://news.yahoo.co.jp/" -q="#liveStream" -o="/tmp/yahoo_news_livestream.png"
// go run cmd/screenshot/main.go -u="https://news.yahoo.co.jp/" -q="section.toptopics" -o="/tmp/yahoo_news_toptopics.png"
// # full screenshot
// go run cmd/screenshot/main.go -u="https://www.yahoo.co.jp/" -q="html" -wi=1280 -he=800 -o=/tmp/s.png
// # viewport screenshot (no selector)
// go run cmd/screenshot/main.go -u="https://www.yahoo.co.jp/" -wi=1280 -he=800 -o=/tmp/s.png
//
// [ References ]
// querySelector()を使うとjQueryみたいにセレクターで要素を取得できるよ。（DOMおれおれAdvent Calendar 2015 – 02日目） ｜ Ginpen.com
// https://ginpen.com/2015/12/02/queryselector-api-like-jquery/
func main() {
	flag.Parse()
	if *arguments.url == "" {
		flag.Usage()
		os.Exit(0)
	}

	// --- 1. Profile cache setup ---
	profileCacheDir := setupProfileCache()

	// --- 2. Browser context ---
	browserCtx, shutdownBrowser := newBrowserContext()
	defer shutdownBrowser()

	// --- 3. Log settings ---
	logSettings(profileCacheDir)

	// --- 4. Take screenshot ---
	buf, err := takeScreenshot(browserCtx)
	if err != nil {
		log.Fatal(err)
	}

	// --- 5. Write output ---
	if err := os.WriteFile(*arguments.outputPath, buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("saved screenshot: %s", *arguments.outputPath)

	// --- 6. Cleanup ---
	// Shut down Chrome before deleting profile to release file locks
	shutdownBrowser()
	cleanupProfileCache(profileCacheDir)
}

// setupProfileCache copies the specified Chrome profile to a cache directory.
// Returns the cache directory path to clean up later (empty string if no profile specified).
//
// Cache structure:
//
//	~/.chromedpscreenshot/userdata-<profileName>/             <- user-data-dir
//	~/.chromedpscreenshot/userdata-<profileName>/<profileName>/ <- copied profile data
func setupProfileCache() string {
	if *arguments.profileDir == "" {
		return ""
	}

	profileName := filepath.Base(*arguments.profileDir)
	userDataDir := filepath.Join(chromeProfileCacheRoot(), "userdata-"+profileName)
	profileSubDir := filepath.Join(userDataDir, profileName)

	if _, err := os.Stat(profileSubDir); os.IsNotExist(err) {
		if err := os.MkdirAll(userDataDir, 0700); err != nil {
			log.Fatalf("failed to create cache dir: %v", err)
		}
		if err := os.CopyFS(profileSubDir, os.DirFS(*arguments.profileDir)); err != nil {
			log.Fatalf("failed to copy profile: %v", err)
		}
		log.Printf("copied profile: %s -> %s", *arguments.profileDir, profileSubDir)
	} else {
		log.Printf("reuse cached profile: %s", profileSubDir)
	}

	return userDataDir
}

// newBrowserContext creates a chromedp browser context with the configured options.
// Returns the context and a shutdown function that can be called multiple times safely.
func newBrowserContext() (context.Context, func()) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", !*arguments.noHeadless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
	)

	if *arguments.profileDir != "" {
		profileName := filepath.Base(*arguments.profileDir)
		userDataDir := filepath.Join(chromeProfileCacheRoot(), "userdata-"+profileName)
		opts = append(opts,
			chromedp.Flag("user-data-dir", userDataDir),
			chromedp.Flag("profile-directory", profileName),
			chromedp.Flag("use-mock-keychain", false),
			chromedp.Flag("password-store", "keychain"),
		)
	}

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)

	var ctxOpts []chromedp.ContextOption
	if *arguments.debug {
		ctxOpts = append(ctxOpts, chromedp.WithDebugf(log.Printf))
	} else {
		ctxOpts = append(ctxOpts, chromedp.WithLogf(log.Printf))
	}

	browserCtx, browserCancel := chromedp.NewContext(allocCtx, ctxOpts...)

	// Handle interrupt signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)
	go func() {
		<-signals
		browserCancel()
		allocCancel()
		os.Exit(0)
	}()

	var once sync.Once
	shutdown := func() {
		once.Do(func() {
			browserCancel()
			allocCancel()
		})
	}

	return browserCtx, shutdown
}

// takeScreenshot navigates to the URL and captures a screenshot.
func takeScreenshot(ctx context.Context) ([]byte, error) {
	var buf []byte

	// Set viewport
	if err := chromedp.Run(ctx, emulation.SetDeviceMetricsOverride(
		*arguments.windowWidth,
		*arguments.windowHeight,
		*arguments.deviceScaleFactor,
		false,
	)); err != nil {
		return nil, err
	}

	// Build tasks: navigate -> wait -> capture
	tasks := chromedp.Tasks{
		// Use page.Navigate directly to avoid hanging on pages
		// that never fire the load event within a constrained viewport.
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, _, _, _, err := page.Navigate(*arguments.url).Do(ctx)
			return err
		}),
		chromedp.Sleep(3 * time.Second),
	}

	// Wait & capture
	switch {
	case *arguments.fullScreenshot:
		tasks = append(tasks, chromedp.FullScreenshot(&buf, 100))
	case *arguments.querySelector != "":
		tasks = append(tasks,
			chromedp.WaitVisible(*arguments.querySelector, chromedp.ByQuery),
			chromedp.Screenshot(*arguments.querySelector, &buf, chromedp.ByQuery),
		)
	default:
		// No selector: capture the entire viewport
		tasks = append(tasks, chromedp.ActionFunc(func(ctx context.Context) error {
			data, err := page.CaptureScreenshot().WithFormat(page.CaptureScreenshotFormatPng).Do(ctx)
			if err != nil {
				return err
			}
			buf = data
			return nil
		}))
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		return nil, err
	}
	return buf, nil
}

// cleanupProfileCache deletes the cached profile directory unless -s is specified.
func cleanupProfileCache(cacheDir string) {
	if cacheDir == "" || *arguments.saveProfile {
		return
	}
	log.Printf("delete cached profile: %s", cacheDir)
	if err := os.RemoveAll(cacheDir); err != nil {
		log.Printf("failed to delete cached profile: %v", err)
	}
}

// chromeProfileCacheRoot returns the root directory for cached Chrome profiles.
// Can be overridden by the CHROMEDP_SCREENSHOT_CACHE_DIR environment variable.
func chromeProfileCacheRoot() string {
	if dir := os.Getenv("CHROMEDP_SCREENSHOT_CACHE_DIR"); dir != "" {
		return dir
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home directory: %v", err)
	}
	return filepath.Join(homeDir, ".chromedpscreenshot")
}

func logSettings(profileCacheDir string) {
	log.Printf("            url: %s", *arguments.url)
	log.Printf("          query: %s", *arguments.querySelector)
	log.Printf("         output: %s", *arguments.outputPath)
	log.Printf("    profile dir: %s", *arguments.profileDir)
	log.Printf("       viewport: %dx%d", *arguments.windowWidth, *arguments.windowHeight)
	log.Printf("   scale factor: %.1f", *arguments.deviceScaleFactor)
	log.Printf("full screenshot: %v", *arguments.fullScreenshot)
	log.Printf("       headless: %v", !*arguments.noHeadless)
	if profileCacheDir != "" {
		log.Printf("  profile cache: %s", profileCacheDir)
		log.Printf("   save profile: %v", *arguments.saveProfile)
	}
}
