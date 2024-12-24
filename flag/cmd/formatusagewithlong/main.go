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
	CommandDescription  = "Here is the command description."
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
	formatUsage()
}

// << Execution sample >>
// $ go run cmd/formatusagewithlong/main.go -a 12 -i test
// [ Command options ]
// -a 12                 ########
// -add 12               [required] add
// -c 1                  ########
// -count 1              count
// -f 10                 ########
// -filesize 10          filesize
// -h false              ########
// -help false           help
// -i test               ########
// -item-name test       [required] item-name
//
// $ go run cmd/formatusagewithlong/main.go --add 14 --item-name test2
// [ Command options ]
// -a 14                 ########
// -add 14               [required] add
// -c 1                  ########
// -count 1              count
// -f 10                 ########
// -filesize 10          filesize
// -h false              ########
// -help false           help
// -i test2              ########
// -item-name test2      [required] item-name
//
// $ go run cmd/formatusagewithlong/main.go -h
// Usage: /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build624316317/b001/exe/main [OPTIONS]
//
// Description:
//   Here is the command description.
//
// Options:
//   -a, --add int             [required] add
//   -i, --item-name string    [required] item-name
//   -c, --count int           count (default 1)
//   -f, --filesize int        filesize (default 10)
//   -h, --help                help

func main() {

	flag.Parse()
	if *optionHelp || *optionAdd == 0 || *optionItemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("-%-20s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}

func formatUsage() {
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
