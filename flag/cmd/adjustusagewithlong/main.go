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
	// Define short options ( don't set default value ).
	optionAdd      = flag.Int("a", 0, UsageDummy)
	optionItemName = flag.String("i", "", UsageDummy)
	optionFilesize = flag.Int("f", 0, UsageDummy)
	optionCount    = flag.Int("c", 0, UsageDummy)
	optionHelp     = flag.Bool("h", false, UsageDummy)
)

func init() {
	// Define long parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.IntVar(optionAdd /*         */, "add" /*       */, 0 /*     */, UsageRequiredPrefix+"add")
	flag.StringVar(optionItemName /* */, "item-name" /* */, "" /*    */, UsageRequiredPrefix+"item-name")
	// Optional parameters
	flag.IntVar(optionFilesize /*    */, "filesize" /*  */, 10 /*    */, "filesize")
	flag.IntVar(optionCount /*       */, "count" /*     */, 1 /*     */, "count")
	flag.BoolVar(optionHelp /*       */, "help" /*      */, false /* */, "help")

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
	if *optionHelp || *optionAdd == 0 || *optionItemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("add:", *optionAdd)
	fmt.Println("filesize:", *optionFilesize)
	fmt.Println("item-name:", *optionItemName)
	fmt.Println("count:", *optionCount)
}

func adjustUsage() {
	// Get default flags usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	// Get default flags usage
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + ".*\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageOptions := re.FindAllString(b.String(), -1)
	maxLength := 0.0
	sort.Slice(usageOptions, func(i, j int) bool {
		maxLength = math.Max(maxLength, math.Max(float64(len(re.ReplaceAllString(usageOptions[i], "$1, -$3$4"))), float64(len(re.ReplaceAllString(usageOptions[j], "$1, -$3$4")))))
		if len(strings.Split(usageOptions[i]+usageOptions[j], UsageRequiredPrefix))%2 == 1 {
			return strings.Compare(usageOptions[i], usageOptions[j]) == -1
		} else {
			return strings.Index(usageOptions[i], UsageRequiredPrefix) >= 0
		}
	})
	usage := strings.Replace(strings.Replace(strings.Split(b.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + CommandDescription + "\n\nOptions:\n"
	for _, v := range usageOptions {
		usage += fmt.Sprintf("%-6s%-"+strconv.Itoa(int(maxLength))+"s", re.ReplaceAllString(v, "  $1,"), re.ReplaceAllString(v, "-$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
