package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Define options and description
	// ( the -h, --help option is defined by default in the flag package )
	//
	optionLevel    = flag.Int("l" /*     */, 0 /*       */, "(REQ) level")
	optionName     = flag.String("n" /*  */, "" /*      */, "(REQ) name")
	optionBirthday = flag.String("b" /*  */, "" /*      */, "birthday (format: 1900/01/02)")
	optionWeight   = flag.Float64("f" /* */, 60.0 /*    */, "weight")
	optionDebug    = flag.Bool("d" /*    */, false /*   */, "\ndebug")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n\n", func() string { e, _ := os.Executable(); return filepath.Base(e) }())
		fmt.Fprintf(flag.CommandLine.Output(), "Description:\n  %s\n\n", "Simple flag usage example.")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
}

// << Execution sample >>
//
// $ go run cmd/simple/main.go -l 2 -n John -d
// [ Command options ]
// -b           birthday (format: 1900/01/02)
// -d true      debug
// -f 60        weight
// -l 2         [req] level
// -n John      [req] name
//
// $ go run cmd/simple/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build3551952392/b001/exe/main:
//
//	-b string
//	  	birthday (format: 1900/01/02)
//	-d
//	  	debug
//	-f float
//	  	weight (default 60)
//	-l int
//	  	[req] level
//	-n string
//	  	[req] name
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *optionLevel == 0 || *optionName == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-9s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}
