package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

const UsageRequiredPrefix = "[required] "
const UsageParamNameWidth = "15"

var (
	// Define short parameters
	paramsAdd      = flag.Int("a" /*    */, 0 /*     */, UsageRequiredPrefix+"add")
	paramsItemName = flag.String("i" /* */, "" /*    */, UsageRequiredPrefix+"item-name")
	paramsFilesize = flag.Int("f" /*    */, 10 /*    */, "filesize")
	paramsCount    = flag.Int("c" /*    */, 1 /*     */, "count")
	paramsHelp     = flag.Bool("h" /*   */, false /* */, "help")
)

func init() {
	// Adjust Usage
	adjustUsage()
}

// << Execution sample >>
// $ go run cmd/adjustusage/main.go -a 10 -i test
// add: 10
// filesize: 10
// item-name: test
// count: 1
//
// $ go run cmd/adjustusage/main.go --add 12 --item-name test
// add: 12
// filesize: 10
// item-name: test
// count: 1
//
// $ go run cmd/adjustusage/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build3125881208/b001/exe/main:
//   -a int       [required] add
//   -i string    [required] item-name
//   -c int       count (default 1)
//   -f int       filesize (default 10)
//   -h           help

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
	regex := "\\s+(-\\S+ *\\S*)+\n*\\s+(.+)"
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
	// Adjust usage
	usage := strings.Split(b.String(), "\n")[0] + "\n"
	for _, v := range usageParams {
		usage = usage + fmt.Sprintf("%-"+UsageParamNameWidth+"s", re.ReplaceAllString(v, "  $1")) + re.ReplaceAllString(v, "$2\n")
	}
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}
}
