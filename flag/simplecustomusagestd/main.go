package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Command options (the -h and --help flags are provided by default in the flag package)
	optionFilePath        = flag.String("f" /*   */, "" /*                      */, "\x1b[33m(required)\x1b[0m File path")
	optionUrl             = flag.String("u" /*   */, "https://example/path" /*  */, "URL")
	optionLineIndex       = flag.Int("l" /*      */, 10 /*                      */, "Index of line")
	optionDebug           = flag.Bool("d" /*     */, false /*                   */, "\nEnable debug mode")
	optionDurationWaitSec = flag.Duration("w" /* */, 1*time.Second /*           */, "Duration of wait seconds (e.g., 1s, 500ms, 2m)")
)

func init() {
	// Set custom usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n\n", func() string { e, _ := os.Executable(); return filepath.Base(e) }())
		fmt.Fprintf(flag.CommandLine.Output(), "Description:\n  %s\n\n", "Simple flag usage example.")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
}

// << Execution sample >>
//
// $ go run cmd/simple/main.go -h
// Usage: main [OPTIONS]
//
// Description:
//   Simple flag usage example.
//
// Options:
//   -d
//     	Enable debug mode
//   -f string
//     	(required) File path
//   -l int
//     	Index of line (default 10)
//   -u string
//     	URL (default "https://example/path")
//   -w duration
//     	Duration of wait seconds (e.g., 1s, 500ms, 2m) (default 1s)
//
// $ go run main.go -f main.go
// [ Command options ]
// -d false         Enable debug mode
// -f main.go       (required) File path
// -l 10            Index of line
// -u https://example/path URL
// -w 1s            Duration of wait seconds (e.g., 1s, 500ms, 2m)

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *optionFilePath == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-15s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}
