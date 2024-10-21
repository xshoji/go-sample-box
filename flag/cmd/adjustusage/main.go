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

const UsageRequiredPrefix = "\u001B[33m[required]\u001B[0m "

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
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	re := regexp.MustCompile("\\s+(-\\S+ *\\S*)+\n*\\s+(.+)")
	usageParams := re.FindAllString(b.String(), -1)
	maxLengthParam := 0.0
	sort.Slice(usageParams, func(i, j int) bool {
		maxLengthParam = math.Max(maxLengthParam, math.Max(float64(len(re.ReplaceAllString(usageParams[i], "$1, -$3$4"))), float64(len(re.ReplaceAllString(usageParams[j], "$1, -$3$4")))))
		return strings.Index(usageParams[i], UsageRequiredPrefix) >= 0 || strings.Compare(usageParams[i], usageParams[j]) == -1
	})
	usage := strings.Split(b.String(), "\n")[0] + "\n"
	for _, v := range usageParams {
		usage = usage + fmt.Sprintf("%-"+strconv.Itoa(int(maxLengthParam+2.0))+"s", re.ReplaceAllString(v, "  $1")) + re.ReplaceAllString(v, "$2\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
