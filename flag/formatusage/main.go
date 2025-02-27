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
)

var (
	// Command options
	commandDescription     = "Here is the command description."
	commandOptionMaxLength = 0
	optionAdd              = flag.Int("a" /*    */, 0 /*     */, UsageRequiredPrefix+"add")
	optionItemName         = flag.String("i" /* */, "" /*    */, UsageRequiredPrefix+"item-name")
	optionFilesize         = flag.Int("f" /*    */, 10 /*    */, "filesize")
	optionCount            = flag.Int("c" /*    */, 1 /*     */, "count")
	optionHelp             = flag.Bool("h" /*   */, false /* */, "help")
)

func init() {
	// Format usage
	formatUsage(commandDescription, &commandOptionMaxLength, new(bytes.Buffer))
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
// $ go run main.go
// Usage: main [OPTIONS]
//
// Description:
//	Here is the command description.
//
// Options:
//	-a int     [RQD] add
//	-i string  [RQD] item-name
//	-c int     count (default 1)
//	-f int     filesize (default 10)
//	-h         help

func main() {

	flag.Parse()
	if *optionHelp || *optionAdd == 0 || *optionItemName == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-"+strconv.Itoa(commandOptionMaxLength)+"s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}

func formatUsage(description string, maxLength *int, buffer *bytes.Buffer) {
	// Get default flags usage
	func() { flag.CommandLine.SetOutput(buffer); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	re := regexp.MustCompile("\\s+(-\\S+ *\\S*)+\n*\\s+(.+)")
	usageFirst := strings.Replace(strings.Replace(strings.Split(buffer.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	usageOptions := re.FindAllString(buffer.String(), -1)
	for _, v := range usageOptions {
		*maxLength = max(*maxLength, len(re.ReplaceAllString(v, "  $1")))
	}
	usageOptionsRep := make([]string, 0)
	for _, v := range usageOptions {
		usageOptionsRep = append(usageOptionsRep, fmt.Sprintf("%-"+strconv.Itoa(*maxLength+2)+"s", re.ReplaceAllString(v, "  $1"))+re.ReplaceAllString(v, "$2\n"))
	}
	sort.SliceStable(usageOptionsRep, func(i, j int) bool {
		return strings.Count(usageOptionsRep[i], UsageRequiredPrefix) > strings.Count(usageOptionsRep[j], UsageRequiredPrefix)
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usageFirst+strings.Join(usageOptionsRep, "")) }
}
