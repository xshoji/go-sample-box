package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Title string `short:"t" long:"title" description:"Description of title" required:"true"`
	Count int    `short:"c" long:"count" description:"Description of count"`
	Debug bool   `short:"d" long:"debug" description:"Description of debug"`
}

// - [Goでコマンドラインオプションを処理する](https://blog.web-apps.tech/go-cmdline-option-jessevdk-go-flags/)
// - [jessevdk/go-flags： go command line option parser](https://github.com/jessevdk/go-flags)
func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}
	fmt.Printf("title: %s\n", opts.Title)
	fmt.Printf("count: %d\n", opts.Count)
	fmt.Printf("debug: %t\n", opts.Debug)
}
