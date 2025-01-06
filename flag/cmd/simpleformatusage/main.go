package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	UsageRequiredPrefix = "\u001B[33m[req]\u001B[0m "
)

var (

	// Define options and description
	commandDescription      = "Here is the command description."
	commandOptionFieldWidth = 12
	optionLevel             = flag.Int("l" /*     */, 0 /*       */, UsageRequiredPrefix+"level")
	optionName              = flag.String("n" /*  */, "" /*      */, UsageRequiredPrefix+"name")
	optionBirthday          = flag.String("b" /*  */, "" /*      */, "birthday (format: 1900/01/02)")
	optionWeight            = flag.Float64("f" /* */, 60.0 /*    */, "weight")
	optionHelp              = flag.Bool("h" /*    */, false /*   */, "\nhelp")
	optionDebug             = flag.Bool("d" /*    */, false /*   */, "\ndebug")
)

// Set formatted usage
func init() {
	formatUsage(commandDescription, commandOptionFieldWidth)
}

// << Execution sample >>
//
// $ go run cmd/simpleformatusage/main.go -l 2 -n John -d
// [ Command options ]
// -b           birthday (format: 1900/01/02)
// -d true      debug
// -f 60        weight
// -h false     help
// -l 2         [req] level
// -n John      [req] name
//
// $ go run cmd/simpleformatusage/main.go -h
// Usage: /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build3001085942/b001/exe/main [OPTIONS]
//
// Description:
//   Here is the command description.
//
// Options:
//   -b string  birthday (format: 1900/01/02)
//   -d         debug
//   -f float   weight (default 60)
//   -h         help
//   -l int     [req] level
//   -n string  [req] name

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *optionHelp || *optionLevel == 0 || *optionName == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-9s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}

// formatUsage optionFieldWidth [ general: 12, bool only: 5 ]
func formatUsage(description string, optionFieldWidth int) {
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	usageLines := strings.Split(b.String(), "\n")
	usage := strings.Replace(strings.Replace(usageLines[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	re := regexp.MustCompile(" +(-\\S+)( *\\S*|\t)*\n(\\s+)(.*)\n")
	usage += re.ReplaceAllStringFunc(strings.Join(usageLines[1:], "\n"), func(m string) string {
		parts := re.FindStringSubmatch(m)
		return fmt.Sprintf("  %-"+strconv.Itoa(optionFieldWidth)+"s %s\n", parts[1]+" "+strings.TrimSpace(parts[2]), parts[4])
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
