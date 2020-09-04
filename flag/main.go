package main

import (
	"flag"
	"fmt"
	"os"
)

const Separator = ""

//> go - Flag command line parsing in golang - Stack Overflow
//> https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	//
	// Define long parameters ( don't set default value ).
	//
	aIntFlag    = flag.Int("apple", /*          */ 0, Separator)
	bStringFlag = flag.String("bridge", /*      */ "", Separator)
	eStringFlag = flag.String("estimation", /*  */ "", Separator)
	fIntFlag    = flag.Int("forget", /*         */ 0, Separator)
	gBoolFlag   = flag.Bool("gulp", /*          */ false, Separator)
	iStringFlag = flag.String("iterator", /*    */ "", Separator)
	helpFlag    = flag.Bool("help", /*          */ false, Separator)
	titleFlag   = flag.String("title", /*       */ "", Separator)
	countFlag   = flag.Int("count", /*          */ 0, Separator)
	isDebugFlag = flag.Bool("debug", /*         */ false, Separator)
)

func init() {
	//
	// Define short parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(aIntFlag, /*         */ "a", 0, /*         */ "[required] apple")
	flag.StringVar(iStringFlag, /*   */ "i", "", /*        */ "[required] iterator")
	flag.StringVar(titleFlag, /*     */ "t", "", /*        */ "[required] title")
	// Optional parameters
	flag.StringVar(bStringFlag, /*   */ "b", "BIG", /*     */ "[optional] bridge")
	flag.StringVar(eStringFlag, /*   */ "e", "NO", /*      */ "[optional] estimation")
	flag.IntVar(fIntFlag, /*         */ "f", 10, /*        */ "[optional] forget")
	flag.BoolVar(gBoolFlag, /*       */ "g", false, /*     */ "[optional] gulp")
	flag.BoolVar(helpFlag, /*        */ "h", false, /*     */ "help")
	flag.IntVar(countFlag, /*        */ "c", 1, /*         */ "[optional] count")
	flag.BoolVar(isDebugFlag, /*     */ "d", false, /*     */ "[optional] debug")
}

// << Execution sample >>
//
// $ go run main.go -i "param1" -t "param2" -a 100
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
// $ go run main.go -iterator "param1" -title "param2" -apple 100
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
// $ go run main.go -h
// Usage of /var/folders/2y/fcx63zfs20g9f_7ktsdhbts0bsslmr/T/go-build676215051/b001/exe/main:
//   -a int
//     	[required] apple
//   -apple int
// 
//   -b string
//     	[optional] bridge (default "BIG")
//   -bridge string
// 
//   -c int
//     	[optional] count (default 1)
//   -count int
// 
//   -d	[optional] debug
//   -debug
// 
//   -e string
//     	[optional] estimation (default "NO")
//   -estimation string
// 
//   -f int
//     	[optional] forget (default 10)
//   -forget int
// 
//   -g	[optional] gulp
//   -gulp
// 
//   -h	help
//   -help
// 
//   -i string
//     	[required] iterator
//   -iterator string
// 
//   -t string
//     	[required] title
//   -title string
//
func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *helpFlag || *aIntFlag == 0 || *iStringFlag == "" || *titleFlag == ""{
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("apple: ", *aIntFlag)
	fmt.Println("bridge: ", *bStringFlag)
	fmt.Println("estimation: ", *eStringFlag)
	fmt.Println("forget: ", *fIntFlag)
	fmt.Println("gulp: ", *gBoolFlag)
	fmt.Println("iterator: ", *iStringFlag)
	fmt.Println("title: ", *titleFlag)
	fmt.Println("count: ", *countFlag)
	fmt.Println("debug: ", *isDebugFlag)
}
