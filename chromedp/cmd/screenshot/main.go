package main

import (
	"context"
	"flag"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	arguments = struct {
		url           *string
		querySelector *string
		outputPath    *string
		windowWidth   *int64
		windowHeight  *int64
		debug         *bool
		noHeadless    *bool
		help          *bool
	}{
		flag.String("u", "" /*              */, "[Required] URL"),
		flag.String("q", "" /*              */, "[Required] Query selector. A screenshot target is the first element node matching the selector. ( e.g. -q=\".className#id\" )"),
		flag.String("o", `/tmp/img.png` /*  */, "Output path of screenshot"),
		flag.Int64("wi", 3840 /*            */, "Set width of window"),
		flag.Int64("he", 2160 /*            */, "Set height of window"),
		flag.Bool("d", false /*             */, "\nEnable debug mode"),
		flag.Bool("n", false /*             */, "\nDisable Headless mode"),
		flag.Bool("h", false /*             */, "\nShow help"),
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
	if *arguments.help || *arguments.url == "" || *arguments.querySelector == "" {
		flag.Usage()
		os.Exit(0)
	}

	log.Printf("url: %v\n", *arguments.url)
	log.Printf("query: %v\n", *arguments.querySelector)
	log.Printf("output: %v\n", *arguments.outputPath)
	log.Printf("width: %v\n", *arguments.windowWidth)
	log.Printf("height: %v\n", *arguments.windowHeight)

	var err error

	// create context
	// > go - How to use Chrome headless with chromedp? - Stack Overflow
	// > https://stackoverflow.com/questions/44067030/how-to-use-chrome-headless-with-chromedp
	// > How to run chromedp on foreground? · Issue #495 · chromedp/chromedp
	// > https://github.com/chromedp/chromedp/issues/495
	ctxt, cancel := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", !*arguments.noHeadless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
	)...,
	)
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
	err = chromedp.Run(ctxt, emulation.SetDeviceMetricsOverride(*arguments.windowWidth, *arguments.windowHeight, 1, false))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var buf []byte
	err = chromedp.Run(ctxt, screenshot(*arguments.url, *arguments.querySelector, &buf))
	if err != nil {
		log.Fatal(err)
	}

	// write to file
	if err := ioutil.WriteFile(*arguments.outputPath, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func screenshot(url, sel string, buf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(1 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, buf, chromedp.ByQuery),
	}
}
