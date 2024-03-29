package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	params = struct {
		// Required parameters
		add      *int
		itemName *string
		// Optional parameters
		filesize *int
		binary   *string
		debug    *bool
		help     *bool
	}{
		flag.Int("a", 0 /*           */, "[Required] add"),
		flag.String("i", "" /*       */, "[Required] item-name"),
		flag.Int("f", 10 /*          */, "[Optional] filesize"),
		flag.String("b", "000101" /* */, "[Optional] binary"),
		flag.Bool("d", false /*      */, "\ndebug"),
		flag.Bool("h", false /*      */, "\nhelp"),
	}
)

// << Execution sample >>
//
// $ go run cmd/struct/main.go -a 100 -i "itemName"
// add: 100
// item-name: itemName
// filesize: 10
// binary: 000101
// debug: false
//
// $ go run cmd/struct/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build46637023/b001/exe/struct:
//	-a int
//	  	[Required] add
//	-b string
//	  	[Optional] binary (default "000101")
//	-d
//	  	debug
//	-f int
//	  	[Optional] filesize (default 10)
//	-h
//	  	help
//	-i string
//	  	[Required] item-name

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *params.help || *params.add == 0 || *params.itemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *params.add)
	fmt.Println("item-name:", *params.itemName)
	fmt.Println("filesize:", *params.filesize)
	fmt.Println("binary:", *params.binary)
	fmt.Println("debug:", *params.debug)
}
