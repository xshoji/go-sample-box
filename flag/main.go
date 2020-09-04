package main

import (
	"flag"
	"fmt"
	"os"
)

//> go - Flag command line parsing in golang - Stack Overflow
//> https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

const Separator = ""

var (
	// Define long parameters ( don't set default value ).
	helpFlag    = flag.Bool("help", false, Separator)
	titleFlag   = flag.String("title", "", Separator)
	countFlag   = flag.Int("count", 0, Separator)
	isDebugFlag = flag.Bool("debug", false, Separator)
)


func init() {
	// Define short parameters and description ( set default value here if you need ).
	flag.BoolVar(helpFlag, "h", false, "help")
	flag.StringVar(titleFlag, "t", "", "[required] title")
	flag.IntVar(countFlag, "c", 1, "[optional] count")
	flag.BoolVar(isDebugFlag, "d", false, "[optional] debug")
}

// << Execution sample >>
//
// $ go run main.go -t test
// title:  test
// count:  1
// debug:  false
//
// $ go run main.go -title test2
// title:  test2
// count:  1
// debug:  false
//
// $ go run main.go -h
// Usage of /var/folders/2y/fcx63zfs20g9f_7ktsdhbts0bsslmr/T/go-build235323428/b001/exe/main:
//   -c int
//     	[optional] count (default 1)
//   -count int
//
//   -d	[optional] debug
//   -debug
//
//   -h	help
//   -help
//
//   -t string
//     	[required] title
//   -title string
//
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *helpFlag || *titleFlag == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("title: ", *titleFlag)
	fmt.Println("count: ", *countFlag)
	fmt.Println("debug: ", *isDebugFlag)

}
