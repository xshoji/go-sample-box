package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	arguments = struct {
		// Required parameters
		aInt    *int
		iString *string
		// Optional parameters
		fInt    *int
		bString *string
		debug   *bool
		help    *bool
	}{
		flag.Int("a", 0 /*          */, "[Required] apple count"),
		flag.String("i", "" /*      */, "[Required] iterator class name"),
		flag.Int("f", 10 /*         */, "[Optional] flagship"),
		flag.String("b", "BIG" /*   */, "[Optional] bridge"),
		flag.Bool("d", false /*   */, "\n[Optional] debug"),
		flag.Bool("h", false /*   */, "\nhelp"),
	}
)

// << Execution sample >>
//
// $ go run cmd/struct/main.go -a=100 -i="iterator"
// apple:  100
// iterator:  iterator
// flagship:  10
// bridge:  BIG
// debug:  false
//
// $ go run cmd/struct/main.go -a=100 -i="iterator" -b="SMALL" -d
// apple:  100
// iterator:  iterator
// flagship:  10
// bridge:  SMALL
// debug:  true
//
// $ go run cmd/struct/main.go -a=100 -i="iterator" -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build006578447/b001/exe/main:
//   -a int
//     	[Required] apple count
//   -b string
//     	[Optional] bridge (default "BIG")
//   -d
//     	[Optional] debug
//   -f int
//     	[Optional] flagship (default 10)
//   -h
//     	help
//   -i string
//     	[Required] iterator class name
//
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *arguments.help || *arguments.aInt == 0 || *arguments.iString == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("apple: ", *arguments.aInt)
	fmt.Println("iterator: ", *arguments.iString)
	fmt.Println("flagship: ", *arguments.fInt)
	fmt.Println("bridge: ", *arguments.bString)
	fmt.Println("debug: ", *arguments.debug)
}
