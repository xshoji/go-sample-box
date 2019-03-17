package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
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
	parser.Name = "chromedp-auto-screenshot"
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
	defer cancel()

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

	// run task list
	var res string
	err = c.Run(ctxt, text(opts.Url, opts.QuerySelector, &res))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	defer func() {
		// shutdown chrome
		err = c.Shutdown(ctxt)
		if err != nil {
			log.Fatal(err)
		}

		// wait for chrome to finish
		err = c.Wait()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("\n\n\result: \n%s\n\n\n", res)
}

func text(url string, selector string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.yahoo.co.jp/`),
		chromedp.InnerHTML(`ul.emphasis`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
