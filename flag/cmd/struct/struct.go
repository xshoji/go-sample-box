package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	args = struct {
		// Required parameters
		apple    *int
		iterator *string
		// Optional parameters
		flagship *int
		bridge   *string
		debug    *bool
		help     *bool
	}{
		flag.Int("a", 0 /*          */, "[Required] apple count"),
		flag.String("i", "" /*      */, "[Required] iterator class name"),
		flag.Int("f", 10 /*         */, "[Optional] flagship"),
		flag.String("b", "BIG" /*   */, "[Optional] bridge"),
		flag.Bool("d", false /*   */, "\ndebug"),
		flag.Bool("h", false /*   */, "\nhelp"),
	}
)

// << Execution sample >>
//
// $ go run cmd/struct/struct.go -a=100 -i="iterator"
// apple:  100
// iterator:  iterator
// flagship:  10
// bridge:  BIG
// debug:  false
//
// $ go run cmd/struct/struct.go -a=100 -i="iterator" -b="SMALL" -d
// apple:  100
// iterator:  iterator
// flagship:  10
// bridge:  SMALL
// debug:  true
//
// $ go run cmd/struct/struct.go -a=100 -i="iterator" -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build006578447/b001/exe/struct:
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
	if *args.help || *args.apple == 0 || *args.iterator == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("apple: ", *args.apple)
	fmt.Println("iterator: ", *args.iterator)
	fmt.Println("flagship: ", *args.flagship)
	fmt.Println("bridge: ", *args.bridge)
	fmt.Println("debug: ", *args.debug)
}
