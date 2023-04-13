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

const UsageDummy = "########"
const UsageRequiredPrefix = "[required] "

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
	flag.CommandLine.SetOutput(b)
	flag.Usage()
	flag.CommandLine.SetOutput(os.Stderr)
	// Sort params and description ( order by UsageRequiredPrefix )
	regex := "(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + "\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)"
	re := regexp.MustCompile(regex)
	usageParams := re.FindAllString(b.String(), -1)
	sort.Slice(usageParams, func(i, j int) bool {
		isRequired1 := strings.Index(usageParams[i], UsageRequiredPrefix) >= 0
		isRequired2 := strings.Index(usageParams[j], UsageRequiredPrefix) >= 0
		if isRequired1 && isRequired2 {
			return strings.Compare(usageParams[i], usageParams[j]) == -1
		} else if isRequired1 {
			return true
		} else {
			return false
		}
	})
	// Calculate max param name
	maxLengthParam := ""
	for _, v := range usageParams {
		paramName := re.ReplaceAllString(v, "$1, -$3$4")
		if len(maxLengthParam) < len(paramName) {
			maxLengthParam = paramName
		}
	}
	// Adjust usage
	usage := strings.Split(b.String(), "\n")[0] + "\n"
	for _, v := range usageParams {
		usage = usage + fmt.Sprintf("%-"+strconv.Itoa(len(maxLengthParam)+4)+"s", re.ReplaceAllString(v, "  $1, -$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}
}
