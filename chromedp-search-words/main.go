package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/jessevdk/go-flags"
	"github.com/yosssi/gohtml"
	"log"
	"os"
	"os/signal"
	"time"
)

type options struct {
	Words      string `short:"w" long:"words" description:"Search words" required:"true"`
	Service    string `short:"s" long:"service" description:"Service type [ yahoo | tabelog ]" default:"yahoo"`
	Debug      bool   `short:"d" long:"debug" description:"Debug mode"`
	NoHeadless bool   `short:"n" long:"no-headless" description:"No Headless mode"`
}

// [ Usage ]
// go run main.go -w="yahoo japan" -s="yahoo"
// go run main.go -w="オムライス" -s="tabelog"
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
	)...)
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
	var tasks chromedp.Tasks
	if opts.Service == "tabelog" {
		tasks = searchTasksTabelog(opts.Words, &res)
	} else {
		tasks = searchTasksYahoo(opts.Words, &res)
	}

	err = chromedp.Run(ctxt, tasks)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\nresult: \n%s\n\n\n", gohtml.Format(res))
}

func searchTasksYahoo(word string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://search.yahoo.co.jp/`),
		chromedp.WaitVisible(`#yschsp`, chromedp.ByQuery),
		chromedp.SendKeys(`#yschsp`, word, chromedp.ByQuery),
		chromedp.WaitVisible(`.sbox_1 .b`, chromedp.ByQuery),
		chromedp.Click(`.sbox_1 .b`, chromedp.ByQuery),
		chromedp.WaitVisible(`#contents`, chromedp.ByQuery),
		chromedp.InnerHTML(`#contents`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}

func searchTasksTabelog(word string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://tabelog.com/`),
		chromedp.WaitVisible(`#sk`, chromedp.ByQuery),
		chromedp.SendKeys(`#sk`, word, chromedp.ByQuery),
		chromedp.WaitVisible(`#js-global-search-btn`, chromedp.ByQuery),
		chromedp.Click(`#js-global-search-btn`, chromedp.ByQuery),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#column-main`, chromedp.ByQuery),
		chromedp.InnerHTML(`#column-main`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
