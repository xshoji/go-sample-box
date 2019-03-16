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
	"time"
)

type options struct {
	Url    string `short:"u" long:"url" description:"url" required:"true"`
	Id     string `short:"i" long:"id" description:"id" required:"true"`
	Output string `short:"o" long:"output" description:"Output file path" default:"/tmp/img.png"`
}

// go run main.go --url="https://www.amazon.co.jp/" --id="navFooter" --output="/tmp/test.png"
func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}
	fmt.Printf("url: %v\n", opts.Url)
	fmt.Printf("id: %v\n", opts.Id)
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

	param := emulation.SetDeviceMetricsOverride(1920, 1080, 1, false)
	// set param
	err = c.Run(ctxt, param)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, screenshot(opts.Url, opts.Id, opts.Output))
	if err != nil {
		log.Fatal(err)
	}

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
}

func screenshot(urlstr, sel string, filePath string) chromedp.Tasks {
	var buf []byte
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, &buf, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return ioutil.WriteFile(filePath, buf, 0644)
		}),
	}
}
