package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	arguments = struct {
		url        *string
		debug      *bool
		noHeadless *bool
		help       *bool
	}{
		flag.String("u", "" /*      */, "[Required] URL"),
		flag.Bool("d", false /*   */, "\nEnable debug mode"),
		flag.Bool("n", false /*   */, "\nDisable Headless mode"),
		flag.Bool("h", false /*   */, "\nShow help"),
	}
)

// [ Usage ]
// go run cmd/evaluatejs/main.go -u="https://news.yahoo.co.jp/"
func main() {

	flag.Parse()
	if *arguments.help || *arguments.url == "" {
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
	var res1 string
	var res2 string
	err = chromedp.Run(ctxt, createTasks(*arguments.url, &res1, &res2))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nres1: \n%s\n\n", res1)
	log.Printf("\nres2: \n%s\n\n", res2)
}

func createTasks(url string, res1 *string, res2 *string) chromedp.Tasks {
	js1 := `
(() => {
  const bodyElement = document.querySelector("body");
  const textContent = bodyElement.children[0].textContent;
  return textContent;
})()
`
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.Evaluate(js1, res1, chromedp.EvalIgnoreExceptions),
		chromedp.Evaluate(`document.title;`, res2, chromedp.EvalIgnoreExceptions),
	}
}
