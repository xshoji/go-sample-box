package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"github.com/yosssi/gohtml"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	arguments = struct {
		keyword    *string
		service    *string
		debug      *bool
		noHeadless *bool
		help       *bool
	}{
		flag.String("k", "" /*         */, "[Required] Search keyword"),
		flag.String("s", "yahoo" /*    */, "Service type [ yahoo | tabelog ]"),
		flag.Bool("d", false /*      */, "\nEnable debug mode"),
		flag.Bool("n", false /*      */, "\nDisable Headless mode"),
		flag.Bool("h", false /*      */, "\nShow help"),
	}
)

// [ Usage ]
// go run cmd/searchwords/main.go -w="yahoo japan"
// go run cmd/searchwords/main.go -w="オムライス" -s="tabelog"
func main() {

	flag.Parse()
	if *arguments.help || *arguments.keyword == "" {
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
	)...)
	defer cancel()
	loggingContextOption := chromedp.WithLogf(log.Printf)
	if *arguments.debug {
		// debug log mode
		loggingContextOption = chromedp.WithDebugf(log.Printf)
	}
	ctxt, _ = chromedp.NewContext(ctxt, loggingContextOption)
	// Signal handling
	go func() {
		// 1. Create channel for os.Signal and Notify interrupt signal to channel.
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM) // syscall.SIGINT, syscall.SIGTERM ( [!] Cannot handle os.Kill )

		// 2. Wait signal
		s := <-signalChannel
		log.Printf("Signal is received: %v\n", s)
		cancel()
		signalValue := int(s.(syscall.Signal))
		os.Exit(signalValue)
	}()

	// run task list
	var res string
	var tasks chromedp.Tasks
	if *arguments.service == "tabelog" {
		tasks = searchTasksTabelog(*arguments.keyword, &res)
	} else {
		tasks = searchTasksYahoo(*arguments.keyword, &res)
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
		chromedp.SendKeys(`.SearchBox__searchInput`, word, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click(`.SearchBox__searchButton`, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.InnerHTML(`#contents`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}

func searchTasksTabelog(word string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://tabelog.com/`),
		chromedp.SendKeys(`#sk`, word, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click(`#js-global-search-btn`, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Sleep(2 * time.Second),
		chromedp.InnerHTML(`.flexible-rstlst`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
