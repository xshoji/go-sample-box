package main

import (
	"flag"
	"fmt"
	"os"
)

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Define short parameters and description ( set default value here if you need ).
	//
	paramsAdd         = flag.Int("a" /*    */, 0 /*       */, "[required] add")
	paramsIncrease    = flag.String("i" /* */, "" /*      */, "[required] increase")
	paramsTitle       = flag.String("t" /* */, "" /*      */, "[required] title")
	paramsFilesize    = flag.Int("f" /*    */, 10 /*      */, "[optional] filesize")
	paramsCount       = flag.Int("c" /*    */, 1 /*       */, "[optional] count")
	paramsBinary      = flag.String("b" /* */, "00101" /* */, "[optional] binary")
	paramsEnvironment = flag.String("e" /* */, "DEV" /*   */, "[optional] environment")
	paramsGlobal      = flag.Bool("g" /*   */, false /*   */, "\nglobal")
	paramsHelp        = flag.Bool("h" /*   */, false /*   */, "\nhelp")
	paramsDebug       = flag.Bool("d" /*   */, false /*   */, "\ndebug")
)

// << Execution sample >>
//
// $ go run cmd/simple/main.go -i "param1" -t "param2" -a 100
// add: 100
// binary: 00101
// environment: DEV
// filesize: 10
// global: false
// increase: param1
// title: param2
// count: 1
// debug: false
//
// $ go run cmd/simple/main.go -h
// Usage of /var/folders/2c/_9j92fnj5z3754dw8_h345zc0000gn/T/go-build548286484/b001/exe/simple:
//	-a int
//	      [required] apple
//	-b string
//	      [optional] bridge (default "BIG")
//	-c int
//	      [optional] count (default 1)
//	-d
//	      debug
//	-e string
//	      [optional] estimation (default "NO")
//	-f int
//	      [optional] forget (default 10)
//	-g
//	      gulp
//	-h
//	      help
//	-i string
//	      [required] iterator
//	-t string
//	      [required] title

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsHelp || *paramsAdd == 0 || *paramsIncrease == "" || *paramsTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *paramsAdd)
	fmt.Println("binary:", *paramsBinary)
	fmt.Println("environment:", *paramsEnvironment)
	fmt.Println("filesize:", *paramsFilesize)
	fmt.Println("global:", *paramsGlobal)
	fmt.Println("increase:", *paramsIncrease)
	fmt.Println("title:", *paramsTitle)
	fmt.Println("count:", *paramsCount)
	fmt.Println("debug:", *paramsDebug)
}
