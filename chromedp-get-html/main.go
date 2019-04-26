package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"os/signal"
)

type options struct {
	Url           string `short:"u" long:"url" description:"URL" required:"true"`
	QuerySelector string `short:"q" long:"queryselector" description:"Queryselector used to output as string" required:"true"`
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
	ctxt, cancel := context.WithCancel(context.Background())

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf), chromedp.WithRunnerOptions(
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("no-first-run", true),
		runner.Flag("no-default-browser-check", true),
		runner.RemoteDebuggingPort(9222),
	))
	if err != nil {
		log.Fatal(err)
	}

	// defer handling
	defer cancel()
	shutdownFunc := func() {
		err = c.Shutdown(ctxt)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer shutdownFunc()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)
	go func() {
		<-signals
		cancel()
		shutdownFunc()
		os.Exit(0)
	}()

	// run task list
	var res string
	err = c.Run(ctxt, text(opts.Url, opts.QuerySelector, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\nresult: \n%s\n\n\n", res)
}

func text(url string, selector string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.InnerHTML(selector, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
