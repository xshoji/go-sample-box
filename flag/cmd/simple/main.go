package main

import (
	"flag"
	"fmt"
	"os"
)

//> go - Flag command line parsing in golang - Stack Overflow
//> https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	//
	// Define short parameters and description ( set default value here if you need ).
	//
	// Required parameters
	argumentApple    = flag.Int("a", 0 /*        */, "[required] apple")
	argumentIterator = flag.String("i", "" /*    */, "[required] iterator")
	argumentTitle    = flag.String("t", "" /*    */, "[required] title")
	// Optional parameters
	argumentFlagship   = flag.Int("f", 10 /*       */, "[optional] flagship")
	argumentCount      = flag.Int("c", 1 /*        */, "[optional] count")
	argumentBridge     = flag.String("b", "BIG" /* */, "[optional] bridge")
	argumentEstimation = flag.String("e", "NO" /*  */, "[optional] estimation")
	argumentGulp       = flag.Bool("g", false /*   */, "\ngulp")
	argumentHelp       = flag.Bool("h", false /*   */, "\nhelp")
	argumentDebug      = flag.Bool("d", false /*   */, "\ndebug")
)

// << Execution sample >>
//
// $ go run cmd/simple/main.go -i "param1" -t "param2" -a 100
// apple:  100
// bridge:  BIG
// estimation:  NO
// forget:  10
// gulp:  false
// iterator:  param1
// title:  param2
// count:  1
// debug:  false
//
// $ go run cmd/simple/main.go -iterator "param1" -title "param2" -apple 100
// apple:  100
// bridge:  BIG
// estimation:  NO
// forget:  10
// gulp:  false
// iterator:  param1
// title:  param2
// count:  1
// debug:  false
//
// $ go run cmd/simple/main.go -h
// Usage of /var/folders/2c/_9j92fnj5z3754dw8_h345zc0000gn/T/go-build548286484/b001/exe/main:
//   -a int
//         [required] apple
//   -b string
//         [optional] bridge (default "BIG")
//   -c int
//         [optional] count (default 1)
//   -d
//         debug
//   -e string
//         [optional] estimation (default "NO")
//   -f int
//         [optional] forget (default 10)
//   -g
//         gulp
//   -h
//         help
//   -i string
//         [required] iterator
//   -t string
//         [required] title
//
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *argumentHelp || *argumentApple == 0 || *argumentIterator == "" || *argumentTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("apple: ", *argumentApple)
	fmt.Println("bridge: ", *argumentBridge)
	fmt.Println("estimation: ", *argumentEstimation)
	fmt.Println("flagship: ", *argumentFlagship)
	fmt.Println("gulp: ", *argumentGulp)
	fmt.Println("iterator: ", *argumentIterator)
	fmt.Println("title: ", *argumentTitle)
	fmt.Println("count: ", *argumentCount)
	fmt.Println("debug: ", *argumentDebug)
}
