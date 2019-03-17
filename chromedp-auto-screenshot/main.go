package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type options struct {
	Url           string `short:"u" long:"url" description:"URL" required:"true"`
	QuerySelector string `short:"q" long:"queryselector" description:"Queryselector used to capture a element" required:"true"`
	Output        string `short:"o" long:"output" description:"Output file path" default:"/tmp/img.png"`
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

	fmt.Printf("url: %v\n", opts.Url)
	fmt.Printf("query: %v\n", opts.QuerySelector)
	fmt.Printf("output: %v\n", opts.Output)

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

	// > screenshot from a wrong page · Issue #205 · chromedp/chromedp
	// > https://github.com/chromedp/chromedp/issues/205
	param := emulation.SetDeviceMetricsOverride(1920, 1080, 1, false)
	// set param
	err = c.Run(ctxt, param)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, screenshot(opts.Url, opts.QuerySelector, opts.Output))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	defer func() {
		err = c.Shutdown(ctxt)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// headlessモードの時は不要
	// wait for chrome to finish
	//err = c.Wait()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func screenshot(urlstr, sel string, filePath string) chromedp.Tasks {
	var buf []byte
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, &buf, chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return ioutil.WriteFile(filePath, buf, 0644)
		}),
	}
}
