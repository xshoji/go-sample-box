package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"os/signal"
)

type options struct {
	Url           string `short:"u" long:"url" description:"URL" required:"true"`
	QuerySelector string `short:"q" long:"query-selector" description:"QuerySelector used to output as string" required:"true"`
	Debug         bool   `short:"d" long:"debug" description:"Debug mode"`
	NoHeadless    bool   `short:"n" long:"no-headless" description:"No Headless mode"`
}

// [ Usage ]
// go run main.go -u="https://news.yahoo.co.jp/" -q="#liveStream"
// go run main.go -u="https://news.yahoo.co.jp/" -q="section.toptopics"
func main() {

	opts := *new(options)
	parser := flags.NewParser(&opts, flags.Default)
	// set name
	parser.Name = "chromedp-get-html"
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

	// run task list
	var res string
	err = chromedp.Run(ctxt, createTasks(opts.Url, opts.QuerySelector, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\nresult: \n%s\n\n\n", res)
}

func createTasks(url string, selector string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.OuterHTML(selector, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
