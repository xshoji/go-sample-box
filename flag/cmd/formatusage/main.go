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
)

var (
	// Command options
	commandDescription = "Here is the command description."
	optionAdd          = flag.Int("a" /*    */, 0 /*     */, UsageRequiredPrefix+"add")
	optionItemName     = flag.String("i" /* */, "" /*    */, UsageRequiredPrefix+"item-name")
	optionFilesize     = flag.Int("f" /*    */, 10 /*    */, "filesize")
	optionCount        = flag.Int("c" /*    */, 1 /*     */, "count")
	optionHelp         = flag.Bool("h" /*   */, false /* */, "help")
)

func init() {
	// Adjust Usage
	formatUsage(commandDescription)
}

// << Execution sample >>
// $ go run cmd/formatusage/main.go -a 10 -i test
// [ Command options ]
// -a 10        [required] add
// -c 1         count
// -f 10        filesize
// -h false     help
// -i test      [required] item-name
//
// $ go run cmd/formatusage/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build3188841424/b001/exe/main:
//
// Description:
//
//	Here is the command description.
//
// Options:
//
//	-a int      [required] add
//	-i string   [required] item-name
//	-c int      count (default 1)
//	-f int      filesize (default 10)
//	-h          help
func main() {

	flag.Parse()
	if *optionHelp || *optionAdd == 0 || *optionItemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-10s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}

func formatUsage(description string) {
	// Get default flags usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	// Get default flags usage
	re := regexp.MustCompile("\\s+(-\\S+ *\\S*)+\n*\\s+(.+)")
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
	usage := strings.Replace(strings.Replace(strings.Split(b.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	for _, v := range usageOptions {
		usage += fmt.Sprintf("%-"+strconv.Itoa(int(maxLength+2.0))+"s", re.ReplaceAllString(v, "  $1")) + re.ReplaceAllString(v, "$2\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
