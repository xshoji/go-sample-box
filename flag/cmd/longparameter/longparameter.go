package main

import (
	"flag"
	"fmt"
	"os"
)

const Usage = ""

//> go - Flag command line parsing in golang - Stack Overflow
//> https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	//
	// Define long parameters ( don't set default value ).
	//
	// Required parameters
	paramsApple    = flag.Int("apple" /*          */, 0, Usage)
	paramsIterator = flag.String("iterator" /*    */, "", Usage)
	paramsTitle    = flag.String("title" /*       */, "", Usage)
	// Optional parameters
	paramsFlagship   = flag.Int("flagship" /*       */, 0, Usage)
	paramsCount      = flag.Int("count" /*          */, 0, Usage)
	paramsBridge     = flag.String("bridge" /*      */, "", Usage)
	paramsEstimation = flag.String("estimation" /*  */, "", Usage)
	paramsGulp       = flag.Bool("gulp" /*          */, false, Usage)
	paramsHelp       = flag.Bool("help" /*          */, false, Usage)
	paramsDebug      = flag.Bool("debug" /*         */, false, Usage)
)

func init() {
	//
	// Define short parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(paramsApple /*         */, "a", 0 /*     */, "[required] apple")
	flag.StringVar(paramsIterator /*   */, "i", "" /*    */, "[required] iterator")
	flag.StringVar(paramsTitle /*      */, "t", "" /*    */, "[required] title")
	// Optional parameters
	flag.IntVar(paramsFlagship /*      */, "f", 10 /*    */, "[optional] flagship")
	flag.IntVar(paramsCount /*         */, "c", 1 /*     */, "[optional] count")
	flag.StringVar(paramsBridge /*     */, "b", "BIG" /* */, "[optional] bridge")
	flag.StringVar(paramsEstimation /* */, "e", "NO" /*  */, "[optional] estimation")
	flag.BoolVar(paramsGulp /*         */, "g", false /* */, "\ngulp")
	flag.BoolVar(paramsHelp /*         */, "h", false /* */, "\nhelp")
	flag.BoolVar(paramsDebug /*        */, "d", false /* */, "\ndebug")
}

// << Execution sample >>
//
// $ go run cmd/longparameter/longparameter.go -i "param1" -t "param2" -a 100
// apple:  100
// bridge:  BIG
// estimation:  NO
// flagship:  10
// gulp:  false
// iterator:  param1
// title:  param2
// count:  1
// debug:  false
//
// $ go run cmd/longparameter/longparameter.go -iterator "param1" -title "param2" -apple 100
// apple:  100
// bridge:  BIG
// estimation:  NO
// flagship:  10
// gulp:  false
// iterator:  param1
// title:  param2
// count:  1
// debug:  false
//
// $ go run cmd/longparameter/longparameter.go -h
// Usage of /var/folders/2c/_9j92fnj5z3754dw8_h345zc0000gn/T/go-build576006264/b001/exe/longparameter:
//  -a int
//    	[required] apple
//  -apple int
//
//  -b string
//    	[optional] bridge (default "BIG")
//  -bridge string
//
//  -c int
//    	[optional] count (default 1)
//  -count int
//
//  -d
//    	debug
//  -debug
//
//  -e string
//    	[optional] estimation (default "NO")
//  -estimation string
//
//  -f int
//    	[optional] flagship (default 10)
//  -flagship int
//
//  -g
//    	gulp
//  -gulp
//
//  -h
//    	help
//  -help
//
//  -i string
//    	[required] iterator
//  -iterator string
//
//  -t string
//    	[required] title
//  -title string
//
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsHelp || *paramsApple == 0 || *paramsIterator == "" || *paramsTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("apple: ", *paramsApple)
	fmt.Println("bridge: ", *paramsBridge)
	fmt.Println("estimation: ", *paramsEstimation)
	fmt.Println("flagship: ", *paramsFlagship)
	fmt.Println("gulp: ", *paramsGulp)
	fmt.Println("iterator: ", *paramsIterator)
	fmt.Println("title: ", *paramsTitle)
	fmt.Println("count: ", *paramsCount)
	fmt.Println("debug: ", *paramsDebug)
}
