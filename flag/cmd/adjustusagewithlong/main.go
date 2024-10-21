package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	UsageRequiredPrefix = "\u001B[33m[required]\u001B[0m "
	UsageDummy          = "########"
	CommandDescription  = "Command description is here."
)

var (
	// Define short parameters ( don't set default value ).
	paramsAdd      = flag.Int("a", 0, UsageDummy)
	paramsItemName = flag.String("i", "", UsageDummy)
	paramsFilesize = flag.Int("f", 0, UsageDummy)
	paramsCount    = flag.Int("c", 0, UsageDummy)
	paramsHelp     = flag.Bool("h", false, UsageDummy)
)

func init() {
	// Define long parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(paramsAdd /*         */, "add" /*       */, 0 /*     */, UsageRequiredPrefix+"add")
	flag.StringVar(paramsItemName /* */, "item-name" /* */, "" /*    */, UsageRequiredPrefix+"item-name")
	// Optional parameters
	flag.IntVar(paramsFilesize /*    */, "filesize" /*  */, 10 /*    */, "filesize")
	flag.IntVar(paramsCount /*       */, "count" /*     */, 1 /*     */, "count")
	flag.BoolVar(paramsHelp /*       */, "help" /*      */, false /* */, "help")

	// Adjust Usage
	adjustUsage()
}

// << Execution sample >>
// $ go run cmd/adjustusagewithlong/main.go -a 12 -i test
// add: 12
// filesize: 10
// item-name: test
// count: 1
//
// $ go run cmd/adjustusagewithlong/main.go --add 12 --item-name test
// add: 12
// filesize: 10
// item-name: test
// count: 1
//
// $ go run cmd/adjustusagewithlong/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build749980003/b001/exe/main:
//   -a, --add int           [required] add
//   -i, --item-name string  [required] item-name
//   -c, --count int         count (default 1)
//   -f, --filesize int      filesize (default 10)
//   -h, --help              help

func main() {

	flag.Parse()
	if *paramsHelp || *paramsAdd == 0 || *paramsItemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *paramsAdd)
	fmt.Println("filesize:", *paramsFilesize)
	fmt.Println("item-name:", *paramsItemName)
	fmt.Println("count:", *paramsCount)
}

func adjustUsage() {
	// Get default flags usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	// Get default flags usage
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + "\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageParams := re.FindAllString(b.String(), -1)
	maxLengthParam := 0.0
	sort.Slice(usageParams, func(i, j int) bool {
		maxLengthParam = math.Max(maxLengthParam, math.Max(float64(len(re.ReplaceAllString(usageParams[i], "$1, -$3$4"))), float64(len(re.ReplaceAllString(usageParams[j], "$1, -$3$4")))))
		if len(strings.Split(usageParams[i]+usageParams[j], UsageRequiredPrefix))%2 == 1 {
			return strings.Compare(usageParams[i], usageParams[j]) == -1
		} else {
			return strings.Index(usageParams[i], UsageRequiredPrefix) >= 0
		}
	})
	usage := strings.Split(b.String(), "\n")[0] + "\n\nDescription:\n  " + CommandDescription + "\n\nOptions:\n"
	for _, v := range usageParams {
		usage += fmt.Sprintf("%-"+strconv.Itoa(int(maxLengthParam+4.0))+"s", re.ReplaceAllString(v, "  $1, -$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
