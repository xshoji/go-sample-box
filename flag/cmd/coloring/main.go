package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	ColorPrinter = struct {
		Red      string
		Green    string
		Yellow   string
		Blue     string
		Purple   string
		Cyan     string
		Gray     string
		White    string
		Colorize func(string, string) string
	}{
		Red:    "\033[31m",
		Green:  "\033[32m",
		Yellow: "\033[33m",
		Blue:   "\033[34m",
		Purple: "\033[35m",
		Cyan:   "\033[36m",
		Gray:   "\033[37m",
		White:  "\033[97m",
		Colorize: func(color string, text string) string {
			if runtime.GOOS == "windows" {
				return text
			}
			colorReset := "\033[0m"
			return color + text + colorReset
		},
	}
)

// How to add colors to your console/terminal output in Go | TwiN
// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
var (
	// Define short parameters and description ( set default value here if you need ).
	//
	paramsAdd      = flag.Int("a" /*    */, 0 /*       */, ColorPrinter.Colorize(ColorPrinter.Yellow, "[required]")+" add")
	paramsItemName = flag.String("i" /* */, "" /*      */, ColorPrinter.Colorize(ColorPrinter.Yellow, "[required]")+" item name")
	paramsFilesize = flag.Int("f" /*    */, 10 /*      */, ColorPrinter.Colorize(ColorPrinter.Green, "[optional]")+" filesize")
	paramsCount    = flag.Int("c" /*    */, 1 /*       */, ColorPrinter.Colorize(ColorPrinter.Green, "[optional]")+" count")
	paramsGlobal   = flag.Bool("g" /*   */, false /*   */, "\n"+ColorPrinter.Colorize(ColorPrinter.Blue, "global"))
	paramsHelp     = flag.Bool("h" /*   */, false /*   */, "\n"+ColorPrinter.Colorize(ColorPrinter.Red, "help"))
	paramsDebug    = flag.Bool("d" /*   */, false /*   */, "\n"+ColorPrinter.Colorize(ColorPrinter.Cyan, "debug"))
)

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsHelp || *paramsAdd == 0 || *paramsItemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *paramsAdd)
	fmt.Println("filesize:", *paramsFilesize)
	fmt.Println("global:", *paramsGlobal)
	fmt.Println("item-name:", *paramsItemName)
	fmt.Println("count:", *paramsCount)
	fmt.Println("debug:", *paramsDebug)
}
