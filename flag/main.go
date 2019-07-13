package main

import (
	"flag"
	"fmt"
	"os"
)

//> go - Flag command line parsing in golang - Stack Overflow
//> https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var helpFlag = flag.Bool("help", false, "help")
var titleFlag = flag.String("title", "", "[require] title")
var countFlag = flag.Int("count", 1, "[option] count")
var isDebugFlag = flag.Bool("debug", false, "[option] debug")

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(titleFlag, "t", "", "= -title")
	flag.IntVar(countFlag, "c", 1, "= -count")
	flag.BoolVar(isDebugFlag, "d", false, "= -debug")
}

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *helpFlag == true || *titleFlag == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("title: ", *titleFlag)
	fmt.Println("count: ", *countFlag)
	fmt.Println("debug: ", *isDebugFlag)

}
