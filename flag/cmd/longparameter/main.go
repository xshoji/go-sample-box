package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
)

const DummyUsage = "########"

// go - Flag command line parsing in golang - Stack Overflow
// https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var (
	// Define short parameters ( don't set default value ).
	paramsAdd         = flag.Int("a", 0, DummyUsage)
	paramsItemName    = flag.String("i", "", DummyUsage)
	paramsTitle       = flag.String("t", "", DummyUsage)
	paramsFilesize    = flag.Int("f", 0, DummyUsage)
	paramsCount       = flag.Int("c", 0, DummyUsage)
	paramsBinary      = flag.String("b", "", DummyUsage)
	paramsEnvironment = flag.String("e", "", DummyUsage)
	paramsGlobal      = flag.Bool("g", false, DummyUsage)
	paramsHelp        = flag.Bool("h", false, DummyUsage)
	paramsDebug       = flag.Bool("d", false, DummyUsage)
)

func init() {
	// Define long parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(paramsAdd /*            */, "add" /*         */, 0 /*          */, "[required] add")
	flag.StringVar(paramsItemName /*    */, "item-name" /*   */, "" /*         */, "[required] item name")
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
// $ go run cmd/longparameter/main.go -i "param1" -t "param2" -a 100
// add: 100
// binary: 00010101
// environment: DEV
// filesize: 10
// global: false
// item-name: param1
// title: param2
// count: 1
// debug: false
//
// $ go run cmd/longparameter/main.go --item-name "param1" --title "param2" --add 100
// add: 100
// binary: 00010101
// environment: DEV
// filesize: 10
// global: false
// item-name: param1
// title: param2
// count: 1
// debug: false
//
// $ go run cmd/longparameter/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build4248489645/b001/exe/main:
//     -a, --add int
//     	[required] add
//     -b, --binary string
//     	[optional] binary (default "00010101")
//     -c, --count int
//     	[optional] count (default 1)
//     -d, --debug
//     	debug
//     -e, --environment string
//     	[optional] environment (default "DEV")
//     -f, --filesize int
//     	[optional] filesize (default 10)
//     -g, --global
//     	global
//     -h, --help
//     	help
//     -i, --item-name string
//     	[required] item name
//     -t, --title string
//     	[required] title

func main() {

	// Set adjusted usage
	b := new(bytes.Buffer)
	flag.CommandLine.SetOutput(b)
	flag.Usage()
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + DummyUsage + "\n*\\s+(-\\S+)( *\\S*)+\n")
	usage := re.ReplaceAllString(b.String(), "  $1, -$3$4\n")
	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}

	flag.Parse()
	if *paramsHelp || *paramsAdd == 0 || *paramsItemName == "" || *paramsTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *paramsAdd)
	fmt.Println("binary:", *paramsBinary)
	fmt.Println("environment:", *paramsEnvironment)
	fmt.Println("filesize:", *paramsFilesize)
	fmt.Println("global:", *paramsGlobal)
	fmt.Println("item-name:", *paramsItemName)
	fmt.Println("title:", *paramsTitle)
	fmt.Println("count:", *paramsCount)
	fmt.Println("debug:", *paramsDebug)
}
