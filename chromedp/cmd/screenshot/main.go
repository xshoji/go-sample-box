package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
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
		profieDir         *string
		windowWidth       *int64
		windowHeight      *int64
		deviceScaleFactor *float64
		fullscreenshot    *bool
		debug             *bool
		noHeadless        *bool
		saveProfile       *bool
	}{
		flag.String("u", "" /*              */, ColorPrinter.Colorize(ColorPrinter.Yellow, "[Required]")+" URL"),
		flag.String("q", "" /*              */, "Query selector. A screenshot target is the first element node matching the selector. ( e.g. -q=\".className#id\" )"),
		flag.String("o", `/tmp/img.png` /*  */, "Output path of screenshot"),
		flag.String("pd", `` /*             */, "Set Google Chrome user profile directory (e.g. (macOS): -pd=\"/Users/you/Library/Application Support/Google/Chrome/Default\"). The profile is cached and reused if it exists. By default the cache is deleted after execution; use -s to keep it."),
		flag.Int64("wi", 1280 /*            */, "Set viewport width (affects page layout, e.g. responsive design). Without -q, this is the output image width"),
		flag.Int64("he", 860 /*             */, "Set viewport height (affects page layout, e.g. responsive design). Without -q, this is the output image height"),
		flag.Float64("ds", 2.0 /*          */, "Set deviceScaleFactor (2.0 = Retina)"),
		flag.Bool("f", false /*             */, "\nEnable full screenshot mode (if true, the -wi, -he flags are ignored)"),
		flag.Bool("d", false /*             */, "\nEnable debug mode"),
		flag.Bool("n", false /*             */, "\nDisable Headless mode"),
		flag.Bool("s", false /*             */, "\nSave cached profile (do not delete after execution)"),
	}
)

// [ Usage ]
// go run cmd/screenshot/main.go -u="https://news.yahoo.co.jp/" -q="#liveStream" -o="/tmp/yahoo_news_livestream.png"
// go run cmd/screenshot/main.go -u="https://news.yahoo.co.jp/" -q="section.toptopics" -o="/tmp/yahoo_news_toptopics.png"
// # full screenshot
// go run cmd/screenshot/main.go -u="https://www.yahoo.co.jp/" -q="html" -wi=1280 -he=800 -o=/tmp/s.png
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

	log.Printf("            url: %v\n", *arguments.url)
	log.Printf("          query: %v\n", *arguments.querySelector)
	log.Printf("         output: %v\n", *arguments.outputPath)
	log.Printf("    profile dir: %v\n", *arguments.profieDir)
	log.Printf("          width: %v\n", *arguments.windowWidth)
	log.Printf("         height: %v\n", *arguments.windowHeight)
	log.Printf("full screenshot: %v\n", *arguments.fullscreenshot)

	var err error

	// Copy profile to a cache directory to avoid lock conflicts with running Chrome.
	// If the cache already exists, it will be reused without copying.
	// By default, the cache is deleted after execution. Use -s to keep it.
	execAllocatorOptions := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", !*arguments.noHeadless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
	)
	var cachedProfileDir string
	if *arguments.profieDir != "" {
		// When using a real Chrome profile, disable mock keychain so that
		// Chrome can decrypt session cookies via the macOS Keychain.
		execAllocatorOptions = append(execAllocatorOptions,
			chromedp.Flag("use-mock-keychain", false),
			chromedp.Flag("password-store", "keychain"),
		)
		profileBase := filepath.Base(*arguments.profieDir)
		cachedProfileUserDataDir := filepath.Join(getChromeProfileCacheDir(), "userdata-"+profileBase)
		cachedProfileDir = cachedProfileUserDataDir
		cachedProfileSubDir := filepath.Join(cachedProfileUserDataDir, profileBase)
		if _, err := os.Stat(cachedProfileSubDir); os.IsNotExist(err) {
			if err := os.MkdirAll(cachedProfileUserDataDir, 0700); err != nil {
				log.Fatalf("failed to create cache dir: %v", err)
			}
			if err := os.CopyFS(cachedProfileSubDir, os.DirFS(*arguments.profieDir)); err != nil {
				log.Fatalf("failed to copy profile: %v", err)
			}
			log.Printf(" copied profile: %v -> %v\n", *arguments.profieDir, cachedProfileSubDir)
		} else {
			log.Printf("  reuse profile: %v\n", cachedProfileSubDir)
		}
		execAllocatorOptions = append(execAllocatorOptions,
			chromedp.Flag("user-data-dir", cachedProfileUserDataDir),
			chromedp.Flag("profile-directory", profileBase),
		)
	}

	// create context
	// > go - How to use Chrome headless with chromedp? - Stack Overflow
	// > https://stackoverflow.com/questions/44067030/how-to-use-chrome-headless-with-chromedp
	// > How to run chromedp on foreground? · Issue #495 · chromedp/chromedp
	// > https://github.com/chromedp/chromedp/issues/495
	ctxt, cancel := chromedp.NewExecAllocator(context.Background(), execAllocatorOptions...)
	defer cancel()
	loggingContextOption := chromedp.WithLogf(log.Printf)
	if *arguments.debug {
		// debug log mode
		loggingContextOption = chromedp.WithDebugf(log.Printf)
	}
	ctxt, cancel = chromedp.NewContext(ctxt, loggingContextOption)
	defer cancel()
	// handle kill signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)
	go func() {
		<-signals
		cancel()
		os.Exit(0)
	}()

	// > screenshot from a wrong page · Issue #205 · chromedp/chromedp
	// > https://github.com/chromedp/chromedp/issues/205
	// set param
	err = chromedp.Run(ctxt, emulation.SetDeviceMetricsOverride(*arguments.windowWidth, *arguments.windowHeight, *arguments.deviceScaleFactor, false))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var buf []byte
	var screenshotFunc chromedp.Action
	if *arguments.fullscreenshot {
		screenshotFunc = chromedp.FullScreenshot(&buf, 100)
	} else if *arguments.querySelector != "" {
		screenshotFunc = chromedp.Screenshot(*arguments.querySelector, &buf, chromedp.ByQuery)
	} else {
		// No selector: capture the entire viewport
		screenshotFunc = chromedp.ActionFunc(func(ctx context.Context) error {
			data, err := page.CaptureScreenshot().WithFormat(page.CaptureScreenshotFormatPng).Do(ctx)
			if err != nil {
				return err
			}
			buf = data
			return nil
		})
	}
	err = chromedp.Run(ctxt, screenshot(*arguments.url, *arguments.querySelector, screenshotFunc))
	if err != nil {
		log.Fatal(err)
	}

	// write to file
	if err := os.WriteFile(*arguments.outputPath, buf, 0644); err != nil {
		log.Fatal(err)
	}

	// Shut down Chrome before cleaning up the profile directory
	cancel()

	// Clean up cached profile directory unless -s is specified
	if cachedProfileDir != "" && !*arguments.saveProfile {
		log.Printf("delete cached profile: %v\n", cachedProfileDir)
		if err := os.RemoveAll(cachedProfileDir); err != nil {
			log.Printf("failed to delete cached profile: %v", err)
		}
	}
}

func screenshot(url, sel string, screenShotAction chromedp.Action) chromedp.Tasks {
	tasks := chromedp.Tasks{
		// Use ActionFunc to navigate without waiting for full page load,
		// which can hang when viewport height is constrained.
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, _, _, _, err := page.Navigate(url).Do(ctx)
			return err
		}),
		chromedp.Sleep(3 * time.Second),
	}
	if sel != "" {
		tasks = append(tasks, chromedp.WaitVisible(sel, chromedp.ByQuery))
	}
	tasks = append(tasks, screenShotAction)
	return tasks
}

func getChromeProfileCacheDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home directory: %v", err)
	}
	return filepath.Join(homeDir, ".chromedpscreenshot", "chrome-profile-cache")
}
