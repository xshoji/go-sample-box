package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

// > Make default values clearer · Issue #145 · jessevdk/go-flags
// > https://github.com/jessevdk/go-flags/issues/145
type options struct {
	Title string `short:"t" long:"title" description:"Description of title" required:"true"`
	Count int    `short:"c" long:"count" description:"Description of count" default:"3"`
	Debug bool   `short:"d" long:"debug" description:"Description of debug"`
}

// - [Goでコマンドラインオプションを処理する](https://blog.web-apps.tech/go-cmdline-option-jessevdk-go-flags/)
// - [jessevdk/go-flags： go command line option parser](https://github.com/jessevdk/go-flags)
// > Calling a program with -h/--help sets error in Parse(opts) ? · Issue #45 · jessevdk/go-flags
// > https://github.com/jessevdk/go-flags/issues/45
func main() {

	// [ simple ]
	//var opts options
	//if _, err := flags.Parse(&opts); err != nil {
	//	// some error handling
	//	return
	//}

	// [ custom name ]
	//opts := *new(options)
	//parser := flags.NewParser(&opts, flags.Default)
	//// set name
	//parser.Name = "argument-go-flags"
	//if _, err := parser.Parse(); err != nil {
	//	// some error handling
	//	return
	//}

	// [ --help and error help ]
	opts := *new(options)
	parser := flags.NewParser(&opts, flags.Default)
	// set name
	parser.Name = "argument-go-flags"
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

	fmt.Printf("title: %s\n", opts.Title)
	fmt.Printf("count: %d\n", opts.Count)
	fmt.Printf("debug: %t\n", opts.Debug)
}
