package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Define options and description
	//
	optionLevel    = flag.Int("l" /*     */, 0 /*       */, "[req] level")
	optionName     = flag.String("n" /*  */, "" /*      */, "[req] name")
	optionBirthday = flag.String("b" /*  */, "" /*      */, "birthday (format: 1900/01/02)")
	optionWeight   = flag.Float64("f" /* */, 60.0 /*    */, "weight")
	optionHelp     = flag.Bool("h" /*    */, false /*   */, "\nhelp")
	optionDebug    = flag.Bool("d" /*    */, false /*   */, "\ndebug")
)

// << Execution sample >>
//
// $ go run cmd/simple/main.go -l 2 -n John -d
// [ Command options ]
// -b           birthday (format: 1900/01/02)
// -d true      debug
// -f 60        weight
// -h false     help
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
//	-h
//	  	help
//	-l int
//	  	[req] level
//	-n string
//	  	[req] name
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *optionHelp || *optionLevel == 0 || *optionName == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("-%s %-7v   %s\n", a.Name, a.Value, strings.Trim(a.Usage, "\n"))
	})
}
