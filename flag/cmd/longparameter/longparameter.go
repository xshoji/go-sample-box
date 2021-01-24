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
	argsApple    = flag.Int("apple" /*          */, 0, Usage)
	argsIterator = flag.String("iterator" /*    */, "", Usage)
	argsTitle    = flag.String("title" /*       */, "", Usage)
	// Optional parameters
	argsFlagship   = flag.Int("flagship" /*       */, 0, Usage)
	argsCount      = flag.Int("count" /*          */, 0, Usage)
	argsBridge     = flag.String("bridge" /*      */, "", Usage)
	argsEstimation = flag.String("estimation" /*  */, "", Usage)
	argsGulp       = flag.Bool("gulp" /*          */, false, Usage)
	argsHelp       = flag.Bool("help" /*          */, false, Usage)
	argsDebug      = flag.Bool("debug" /*         */, false, Usage)
)

func init() {
	//
	// Define short parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(argsApple /*         */, "a", 0 /*     */, "[required] apple")
	flag.StringVar(argsIterator /*   */, "i", "" /*    */, "[required] iterator")
	flag.StringVar(argsTitle /*      */, "t", "" /*    */, "[required] title")
	// Optional parameters
	flag.IntVar(argsFlagship /*      */, "f", 10 /*    */, "[optional] flagship")
	flag.IntVar(argsCount /*         */, "c", 1 /*     */, "[optional] count")
	flag.StringVar(argsBridge /*     */, "b", "BIG" /* */, "[optional] bridge")
	flag.StringVar(argsEstimation /* */, "e", "NO" /*  */, "[optional] estimation")
	flag.BoolVar(argsGulp /*         */, "g", false /* */, "\ngulp")
	flag.BoolVar(argsHelp /*         */, "h", false /* */, "\nhelp")
	flag.BoolVar(argsDebug /*        */, "d", false /* */, "\ndebug")
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
	if *argsHelp || *argsApple == 0 || *argsIterator == "" || *argsTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("apple: ", *argsApple)
	fmt.Println("bridge: ", *argsBridge)
	fmt.Println("estimation: ", *argsEstimation)
	fmt.Println("flagship: ", *argsFlagship)
	fmt.Println("gulp: ", *argsGulp)
	fmt.Println("iterator: ", *argsIterator)
	fmt.Println("title: ", *argsTitle)
	fmt.Println("count: ", *argsCount)
	fmt.Println("debug: ", *argsDebug)
}
