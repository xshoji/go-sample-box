package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/jessevdk/go-flags"
	"github.com/yosssi/gohtml"
	"log"
	"os"
)

type options struct {
	Words   string `short:"w" long:"words" description:"Search words" required:"true"`
	Service string `short:"s" long:"service" description:"Service type [ yahoo | tabelog ]" default:"yahoo"`
}

// [ Usage ]
// go run main.go -u="https://news.yahoo.co.jp/" -q="#liveStream"
// go run main.go -u="https://news.yahoo.co.jp/" -q="section.toptopics"
func main() {

	opts := *new(options)
	parser := flags.NewParser(&opts, flags.Default)
	// set name
	parser.Name = "chromedp-search-words"
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
	//	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var res string
	var tasks chromedp.Tasks
	if opts.Service == "tabelog" {
		tasks = searchTasksTabelog(opts.Words, &res)
	} else {
		tasks = searchTasksYahoo(opts.Words, &res)
	}

	err = c.Run(ctxt, tasks)
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
	}()

	log.Printf("\n\nresult: \n%s\n\n\n", gohtml.Format(res))
}

func searchTasksYahoo(word string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.yahoo.co.jp/`),
		chromedp.WaitVisible(`#srchtxtBg`, chromedp.ByQuery),
		chromedp.SendKeys(`#srchtxtBg > input`, word, chromedp.ByQuery),
		chromedp.WaitVisible(`#srchbtn`, chromedp.ByQuery),
		chromedp.Click(`#srchbtn`, chromedp.ByQuery),
		chromedp.WaitVisible(`#mIn`, chromedp.ByQuery),
		chromedp.InnerHTML(`#mIn`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}

func searchTasksTabelog(word string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://tabelog.com/`),
		chromedp.WaitVisible(`#sk`, chromedp.ByQuery),
		chromedp.SendKeys(`#sk`, word, chromedp.ByQuery),
		chromedp.WaitVisible(`#js-global-search-btn`, chromedp.ByQuery),
		chromedp.Click(`#js-global-search-btn`, chromedp.ByQuery),
		chromedp.WaitVisible(`#column-main`, chromedp.ByQuery),
		chromedp.InnerHTML(`#column-main`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
