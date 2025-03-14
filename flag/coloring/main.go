package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
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
	// ( the -h, --help option is defined by default in the flag package )
	//
	paramsAdd      = flag.Int("a" /*    */, 0 /*       */, ColorPrinter.Colorize(ColorPrinter.Yellow, "[required]")+" add")
	paramsItemName = flag.String("i" /* */, "" /*      */, ColorPrinter.Colorize(ColorPrinter.Yellow, "[required]")+" item name")
	paramsFilesize = flag.Int("f" /*    */, 10 /*      */, ColorPrinter.Colorize(ColorPrinter.Green, "[optional]")+" filesize")
	paramsCount    = flag.Int("c" /*    */, 1 /*       */, ColorPrinter.Colorize(ColorPrinter.Green, "[optional]")+" count")
	paramsGlobal   = flag.Bool("g" /*   */, false /*   */, "\n"+ColorPrinter.Colorize(ColorPrinter.Blue, "global"))
	paramsDebug    = flag.Bool("d" /*   */, false /*   */, "\n"+ColorPrinter.Colorize(ColorPrinter.Cyan, "debug"))
)

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsAdd == 0 || *paramsItemName == "" {
		fmt.Printf("\n%s Missing required option\n\n", ColorPrinter.Colorize(ColorPrinter.Red, "[ERROR]"))
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-10s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}
