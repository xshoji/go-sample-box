package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"os/signal"
)

var (
	arguments = struct {
		url           *string
		querySelector *string
		debug         *bool
		noHeadless    *bool
		help          *bool
	}{
		flag.String("u", "" /*      */, "[Required] URL"),
		flag.String("q", "" /*      */, "[Required] QuerySelector used to output as string"),
		flag.Bool("d", false /*   */, "\nEnable debug mode"),
		flag.Bool("n", false /*   */, "\nDisable Headless mode"),
		flag.Bool("h", false /*   */, "\nShow help"),
	}
)

// [ Usage ]
// go run cmd/gethtml/main.go -u="https://news.yahoo.co.jp/" -q="#liveStream"
// go run cmd/gethtml/main.go -u="https://news.yahoo.co.jp/" -q="section.toptopics"
func main() {

	flag.Parse()
	if *arguments.help || *arguments.url == "" || *arguments.querySelector == "" {
		flag.Usage()
		os.Exit(0)
	}
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

	// run task list
	var res string
	err = chromedp.Run(ctxt, createTasks(*arguments.url, *arguments.querySelector, &res))
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
