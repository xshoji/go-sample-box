package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

type options struct {
	Url           string `short:"u" long:"url" description:"URL" required:"true"`
	QuerySelector string `short:"q" long:"query-selector" description:"QuerySelector used to capture a element" required:"true"`
	Output        string `short:"o" long:"output" description:"Output file path" default:"/tmp/img.png"`
	Debug         bool   `short:"d" long:"debug" description:"Debug mode"`
	NoHeadless    bool   `short:"n" long:"no-headless" description:"No Headless mode"`
}

// [ Usage ]
// go run main.go -u="https://news.yahoo.co.jp/" -q="#liveStream" -o="/tmp/yahoo_news_livestream.png"
// go run main.go -u="https://news.yahoo.co.jp/" -q="section.toptopics" -o="/tmp/yahoo_news_toptopics.png"
//
// [ References ]
// > querySelector()を使うとjQueryみたいにセレクターで要素を取得できるよ。（DOMおれおれAdvent Calendar 2015 – 02日目） ｜ Ginpen.com
// > https://ginpen.com/2015/12/02/queryselector-api-like-jquery/
func main() {
	opts := *new(options)
	parser := flags.NewParser(&opts, flags.Default)
	// set name
	parser.Name = "chromedp-screenshot"
	if _, err := parser.Parse(); err != nil {
		flagsError, _ := err.(*flags.Error)
		// help時は何もしない
		if flagsError.Type == flags.ErrHelp {
			return
		}
		fmt.Println()
		parser.WriteHelp(os.Stdout)
		fmt.Println()
		return
	}

	fmt.Printf("url: %v\n", opts.Url)
	fmt.Printf("query: %v\n", opts.QuerySelector)
	fmt.Printf("output: %v\n", opts.Output)

	var err error

	// create context
	// > go - How to use Chrome headless with chromedp? - Stack Overflow
	// > https://stackoverflow.com/questions/44067030/how-to-use-chrome-headless-with-chromedp
	// > How to run chromedp on foreground? · Issue #495 · chromedp/chromedp
	// > https://github.com/chromedp/chromedp/issues/495
	ctxt, cancel := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", !opts.NoHeadless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
	)...,
	)
	defer cancel()
	loggingContextOption := chromedp.WithLogf(log.Printf)
	if opts.Debug {
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
	err = chromedp.Run(ctxt, emulation.SetDeviceMetricsOverride(3840, 2160, 1, false))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var buf []byte
	err = chromedp.Run(ctxt, screenshot(opts.Url, opts.QuerySelector, &buf))
	if err != nil {
		log.Fatal(err)
	}

	// write to file
	if err := ioutil.WriteFile(opts.Output, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func screenshot(url, sel string, buf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, buf, chromedp.ByQuery),
	}
}
