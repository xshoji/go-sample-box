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

const (
	UsageRequiredPrefix = "\u001B[33m[RQD]\u001B[0m "
	UsageDummy          = "########"
)

var (
	// Define short options ( don't set default value ).
	commandDescription     = "Here is the command description."
	commandOptionMaxLength = 0
	optionAdd              = flag.Int("a", 0, UsageDummy)
	optionItemName         = flag.String("i", "", UsageDummy)
	optionFilesize         = flag.Int("f", 0, UsageDummy)
	optionCount            = flag.Int("c", 0, UsageDummy)
)

func init() {
	// Define long parameters and description ( set default value here if you need ).
	// ( the -h, --help option is defined by default in the flag package )
	//
	// Required parameters
	flag.IntVar(optionAdd /*         */, "add" /*       */, 0 /*     */, UsageRequiredPrefix+"add")
	flag.StringVar(optionItemName /* */, "item-name" /* */, "" /*    */, UsageRequiredPrefix+"item-name")
	// Optional parameters
	flag.IntVar(optionFilesize /*    */, "filesize" /*  */, 10 /*    */, "filesize")
	flag.IntVar(optionCount /*       */, "count" /*     */, 1 /*     */, "count")

	// Adjust Usage
	formatUsage(commandDescription, &commandOptionMaxLength, new(bytes.Buffer))
}

// << Execution sample >>
// $ go run cmd/formatusagewithlong/main.go -a 12 -i test
// [ Command options ]
//  --add 12                 [RQD] add
//  --count 1                count
//  --filesize 10            filesize
//  --item-name test         [RQD] item-name
//
// $ go run cmd/formatusagewithlong/main.go --add 14 --item-name test2
// [ Command options ]
//  --add 14                 [RQD] add
//  --count 1                count
//  --filesize 10            filesize
//  --item-name test2        [RQD] item-name
//
// $ go run cmd/formatusagewithlong/main.go -h
// Usage: /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build624316317/b001/exe/main [OPTIONS] [-h, --help]
//
// Description:
//   Here is the command description.
//
// Options:
//   -a, --add int             [RQD] add
//   -i, --item-name string    [RQD] item-name
//   -c, --count int           count (default 1)
//   -f, --filesize int        filesize (default 10)

func main() {

	flag.Parse()
	if *optionAdd == 0 || *optionItemName == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		if a.Usage == UsageDummy {
			return
		}
		fmt.Printf("  --%-"+strconv.Itoa(commandOptionMaxLength)+"s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}

func formatUsage(description string, maxLength *int, buffer *bytes.Buffer) {
	// Get default flags usage
	func() { flag.CommandLine.SetOutput(buffer); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + ".*\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageFirst := strings.Replace(strings.Replace(strings.Split(buffer.String(), "\n")[0], ":", " [OPTIONS] [-h, --help]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	usageOptions := re.FindAllString(buffer.String(), -1)
	for _, v := range usageOptions {
		*maxLength = max(*maxLength, len(re.ReplaceAllString(v, "$1, -$3$4")))
	}
	usageOptionsRep := make([]string, 0)
	for _, v := range usageOptions {
		usageOptionsRep = append(usageOptionsRep, fmt.Sprintf("%-6s%-"+strconv.Itoa(*maxLength)+"s", re.ReplaceAllString(v, "  $1,"), re.ReplaceAllString(v, "-$3$4"))+re.ReplaceAllString(v, "$5\n"))
	}
	sort.SliceStable(usageOptionsRep, func(i, j int) bool {
		return strings.Count(usageOptionsRep[i], UsageRequiredPrefix) > strings.Count(usageOptionsRep[j], UsageRequiredPrefix)
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usageFirst+strings.Join(usageOptionsRep, "")) }
}
