package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const DummyUsage = "########"

var (
	// Define long parameters ( don't set default value ).
	paramsAdd      = flag.Int("a", 0, DummyUsage)
	paramsIncrease = flag.String("i", "", DummyUsage)
	paramsFilesize = flag.Int("f", 0, DummyUsage)
	paramsCount    = flag.Int("c", 0, DummyUsage)
	paramsHelp     = flag.Bool("h", false, DummyUsage)
)

func init() {
	// Define short parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(paramsAdd /*            */, "add" /*         */, 0 /*          */, "[required] add")
	flag.StringVar(paramsIncrease /*    */, "increase" /*    */, "" /*         */, "[required] increase")
	// Optional parameters
	flag.IntVar(paramsFilesize /*       */, "filesize" /*    */, 10 /*         */, "[optional] filesize")
	flag.IntVar(paramsCount /*          */, "count" /*       */, 1 /*          */, "[optional] count")
	flag.BoolVar(paramsHelp /*          */, "help" /*        */, false /*      */, "help")
}

// << Execution sample >>
// $ go run cmd/adjustusage/main.go -a 12 -i test
// add: 12
// filesize: 10
// increase: test
// count: 1
//
// $ go run cmd/adjustusage/main.go --add 12 --increase test
// add: 12
// filesize: 10
// increase: test
// count: 1
//
// $ go run cmd/longparameter/longparameter.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build565554301/b001/exe/main:
//   -a, --add int          [required] add
//   -i, --increase string  [required] increase
//   -c, --count int        [optional] count (default 1)
//   -f, --filesize int     [optional] filesize (default 10)
//   -h, --help             help

func main() {

	// Set Usage
	adjustUsage1()

	flag.Parse()
	if *paramsHelp || *paramsAdd == 0 || *paramsIncrease == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *paramsAdd)
	fmt.Println("filesize:", *paramsFilesize)
	fmt.Println("increase:", *paramsIncrease)
	fmt.Println("count:", *paramsCount)
}

func adjustUsage1() {
	b := new(bytes.Buffer)
	flag.CommandLine.SetOutput(b)
	flag.Usage()
	regex := "(-\\S+)( *\\S*)+\n*\\s+" + DummyUsage + "\n*\\s+(-\\S+)( *\\S*)+\n\\s+(\\[required\\].+|\\[optional\\].+|\\S+)+"
	re := regexp.MustCompile(regex)
	usage := strings.Split(b.String(), "\n")[0]
	usageParams := re.FindAllString(b.String(), -1)
	sort.Slice(usageParams, func(i, j int) bool {
		reqMatch1, _ := regexp.MatchString(`\[required`, usageParams[i])
		reqMatch2, _ := regexp.MatchString(`\[required`, usageParams[j])
		optMatch1, _ := regexp.MatchString(`\[optional`, usageParams[i])
		optMatch2, _ := regexp.MatchString(`\[optional`, usageParams[j])
		if reqMatch1 && reqMatch2 {
			return strings.Compare(usageParams[i], usageParams[j]) == -1
		} else if reqMatch1 {
			return true
		} else if reqMatch2 {
			return false
		} else if reqMatch2 && optMatch1 {
			return true
		} else if optMatch1 && optMatch2 {
			return strings.Compare(usageParams[i], usageParams[j]) == -1
		} else if optMatch1 {
			return true
		} else {
			return false
		}
	})
	re2 := regexp.MustCompile(regex)
	usage = usage + "\n"
	maxLengthParam := ""
	for _, v := range usageParams {
		paramName := re2.ReplaceAllString(v, "$1, -$3$4")
		if len(maxLengthParam) < len(paramName) {
			maxLengthParam = paramName
		}
	}
	for _, v := range usageParams {
		usage = usage + fmt.Sprintf("%-"+strconv.Itoa(len(maxLengthParam)+4)+"s", re2.ReplaceAllString(v, "	  $1, -$3$4")) + re2.ReplaceAllString(v, "$5\n")
	}
	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}
}
