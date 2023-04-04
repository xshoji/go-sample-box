package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

const Usage = "########"

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Define long parameters ( don't set default value ).
	//
	// Required parameters
	paramsAdd      = flag.Int("a", 0, Usage)
	paramsIncrease = flag.String("i", "", Usage)
	paramsTitle    = flag.String("t", "", Usage)
	// Optional parameters
	paramsFilesize    = flag.Int("f", 0, Usage)
	paramsCount       = flag.Int("c", 0, Usage)
	paramsBinary      = flag.String("b", "", Usage)
	paramsEnvironment = flag.String("e", "", Usage)
	paramsGlobal      = flag.Bool("g", false, Usage)
	paramsHelp        = flag.Bool("h", false, Usage)
	paramsDebug       = flag.Bool("d", false, Usage)
)

func init() {
	// Define short parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(paramsAdd /*            */, "add" /*         */, 0 /*          */, "[required] add")
	flag.StringVar(paramsIncrease /*    */, "increase" /*    */, "" /*         */, "[required] increase")
	flag.StringVar(paramsTitle /*       */, "title" /*       */, "" /*         */, "[required] title")
	// Optional parameters
	flag.IntVar(paramsFilesize /*       */, "filesize" /*    */, 10 /*         */, "[optional] filesize")
	flag.IntVar(paramsCount /*          */, "count" /*       */, 1 /*          */, "[optional] count")
	flag.StringVar(paramsBinary /*      */, "binary" /*      */, "00010101" /* */, "[optional] binary")
	flag.StringVar(paramsEnvironment /* */, "environment" /* */, "DEV" /*      */, "[optional] environment")
	flag.BoolVar(paramsGlobal /*        */, "global" /*      */, false /*      */, "global")
	flag.BoolVar(paramsHelp /*          */, "help" /*        */, false /*      */, "help")
	flag.BoolVar(paramsDebug /*         */, "debug" /*       */, false /*      */, "debug")
}

// << Execution sample >>
//
// $ go run cmd/longparameter/longparameter.go -i "param1" -t "param2" -a 100
// add: 100
// binary: 00010101
// environment: DEV
// filesize: 10
// global: false
// increase: param1
// title: param2
// count: 1
// debug: false
//
// $ $ go run cmd/longparameter/longparameter.go -increase "param1" -title "param2" -add 100
// add: 100
// binary: 00010101
// environment: DEV
// filesize: 10
// global: false
// increase: param1
// title: param2
// count: 1
// debug: false
//
// $ go run cmd/longparameter/longparameter.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build1111939584/b001/exe/longparameter:
//   -a int
//   -add int
//     	[required] add
//   -b string
//   -binary string
//     	[optional] binary (default "00010101")
//   -c int
//   -count int
//     	[optional] count (default 1)
//   -d
//   -debug
//     	debug
//   -e string
//   -environment string
//     	[optional] environment (default "DEV")
//   -f int
//   -filesize int
//     	[optional] filesize (default 10)
//   -g
//   -global
//     	global
//   -h
//   -help
//     	help
//   -i string
//   -increase string
//     	[required] increase
//   -t string
//   -title string
//     	[required] title

func main() {

	// Set Usage
	var b bytes.Buffer
	flag.CommandLine.SetOutput(&b)
	flag.Usage()
	usage := strings.Replace(b.String(), "\n    \t"+Usage, "", -1)
	usage = strings.Replace(usage, "\t"+Usage, "", -1)
	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}

	flag.Parse()
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
