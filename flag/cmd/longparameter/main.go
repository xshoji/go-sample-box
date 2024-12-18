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
	// Define short options ( don't set default value ).
	optionAdd         = flag.Int("a", 0, DummyUsage)
	optionItemName    = flag.String("i", "", DummyUsage)
	optionTitle       = flag.String("t", "", DummyUsage)
	optionFilesize    = flag.Int("f", 0, DummyUsage)
	optionCount       = flag.Int("c", 0, DummyUsage)
	optionBinary      = flag.String("b", "", DummyUsage)
	optionEnvironment = flag.String("e", "", DummyUsage)
	optionGlobal      = flag.Bool("g", false, DummyUsage)
	optionHelp        = flag.Bool("h", false, DummyUsage)
	optionDebug       = flag.Bool("d", false, DummyUsage)
)

func init() {
	// Define long options and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(optionAdd /*            */, "add" /*         */, 0 /*          */, "[required] add")
	flag.StringVar(optionItemName /*    */, "item-name" /*   */, "" /*         */, "[required] item name")
	flag.StringVar(optionTitle /*       */, "title" /*       */, "" /*         */, "[required] title")
	// Optional parameters
	flag.IntVar(optionFilesize /*       */, "filesize" /*    */, 10 /*         */, "[optional] filesize")
	flag.IntVar(optionCount /*          */, "count" /*       */, 1 /*          */, "[optional] count")
	flag.StringVar(optionBinary /*      */, "binary" /*      */, "00010101" /* */, "[optional] binary")
	flag.StringVar(optionEnvironment /* */, "environment" /* */, "DEV" /*      */, "[optional] environment")
	flag.BoolVar(optionGlobal /*        */, "global" /*      */, false /*      */, "global")
	flag.BoolVar(optionHelp /*          */, "help" /*        */, false /*      */, "help")
	flag.BoolVar(optionDebug /*         */, "debug" /*       */, false /*      */, "debug")
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
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + DummyUsage + ".*\n*\\s+(-\\S+)( *\\S*)+\n")
	usage := re.ReplaceAllString(b.String(), "  $1, -$3$4\n")
	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}

	flag.Parse()
	if *optionHelp || *optionAdd == 0 || *optionItemName == "" || *optionTitle == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *optionAdd)
	fmt.Println("binary:", *optionBinary)
	fmt.Println("environment:", *optionEnvironment)
	fmt.Println("filesize:", *optionFilesize)
	fmt.Println("global:", *optionGlobal)
	fmt.Println("item-name:", *optionItemName)
	fmt.Println("title:", *optionTitle)
	fmt.Println("count:", *optionCount)
	fmt.Println("debug:", *optionDebug)
}
