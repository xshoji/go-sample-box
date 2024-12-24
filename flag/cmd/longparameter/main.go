package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
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
// [ Command options ]
// -a 100                ########
// -add 100              [required] add
// -b 00010101           ########
// -binary 00010101      [optional] binary
// -c 1                  ########
// -count 1              [optional] count
// -d false              ########
// -debug false          debug
// -e DEV                ########
// -environment DEV      [optional] environment
// -f 10                 ########
// -filesize 10          [optional] filesize
// -g false              ########
// -global false         global
// -h false              ########
// -help false           help
// -i param1             ########
// -item-name param1     [required] item name
// -t param2             ########
// -title param2         [required] title
//
// $ go run cmd/longparameter/main.go --item-name "param11" --title "param22" --add 200
// [ Command options ]
// -a 200                ########
// -add 200              [required] add
// -b 00010101           ########
// -binary 00010101      [optional] binary
// -c 1                  ########
// -count 1              [optional] count
// -d false              ########
// -debug false          debug
// -e DEV                ########
// -environment DEV      [optional] environment
// -f 10                 ########
// -filesize 10          [optional] filesize
// -g false              ########
// -global false         global
// -h false              ########
// -help false           help
// -i param11            ########
// -item-name param11    [required] item name
// -t param22            ########
// -title param22        [required] title
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

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("-%-20s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}
