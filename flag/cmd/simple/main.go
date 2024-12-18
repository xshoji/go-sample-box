package main

import (
	"flag"
	"fmt"
	"os"
)

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Define short options and description ( set default value here if you need ).
	//
	optionAdd         = flag.Int("a" /*    */, 0 /*       */, "[required] add")
	optionItemName    = flag.String("i" /* */, "" /*      */, "[required] item name")
	optionTitle       = flag.String("t" /* */, "" /*      */, "[required] title")
	optionFilesize    = flag.Int("f" /*    */, 10 /*      */, "[optional] filesize")
	optionCount       = flag.Int("c" /*    */, 1 /*       */, "[optional] count")
	optionBinary      = flag.String("b" /* */, "00101" /* */, "[optional] binary")
	optionEnvironment = flag.String("e" /* */, "DEV" /*   */, "[optional] environment")
	optionGlobal      = flag.Bool("g" /*   */, false /*   */, "\nglobal")
	optionHelp        = flag.Bool("h" /*   */, false /*   */, "\nhelp")
	optionDebug       = flag.Bool("d" /*   */, false /*   */, "\ndebug")
)

// << Execution sample >>
//
// $ go run cmd/simple/main.go -i "param1" -t "param2" -a 100
// add: 100
// binary: 00101
// environment: DEV
// filesize: 10
// global: false
// item-name: param1
// title: param2
// count: 1
// debug: false
//
// $ go run cmd/simple/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build131759646/b001/exe/main:
//  -a int
//    	[required] add
//  -b string
//    	[optional] binary (default "00101")
//  -c int
//    	[optional] count (default 1)
//  -d
//    	debug
//  -e string
//    	[optional] environment (default "DEV")
//  -f int
//    	[optional] filesize (default 10)
//  -g
//    	global
//  -h
//    	help
//  -i string
//    	[required] item name
//  -t string
//    	[required] title

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *optionHelp || *optionAdd == 0 || *optionItemName == "" || *optionTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *optionAdd)
	fmt.Println("binary:", *optionBinary)
	fmt.Println("environment:", *optionEnvironment)
	fmt.Println("filesize:", *optionFilesize)
	fmt.Println("global:", *optionGlobal)
	fmt.Println("item-name:", *optionItemName)
	fmt.Println("title:", *optionTitle)
	fmt.Println("count:", *optionCount)
	fmt.Println("debug:", *optionDebug)
}
