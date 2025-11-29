package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	UsageRequiredPrefix = "\u001B[33m(REQ)\u001B[0m "
)

var (
	// Command options ( the -h, --help option is defined by default in the flag package )
	commandDescription      = "Here is the command description."
	commandOptionFieldWidth = "12" // recommended width = general: 12, bool only: 5
	optionLevel             = flag.Int("l" /*     */, 0 /*       */, UsageRequiredPrefix+"Level")
	optionName              = flag.String("n" /*  */, "" /*      */, UsageRequiredPrefix+"Name")
	optionBirthday          = flag.String("b" /*  */, "" /*      */, "Birthday (format: 1900/01/02)")
	optionWeight            = flag.Float64("f" /* */, 60.0 /*    */, "Weight")
	optionDebug             = flag.Bool("d" /*    */, false /*   */, "Debug")
)

// << Execution sample >>
//
// $ go run flag/simpleformatusage/main.go -h
// Usage: /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build2691711322/b001/exe/main [OPTIONS]
//
// Description:
//   Here is the command description.
//
// Options:
//   -b string    Birthday (format: 1900/01/02)
//   -d           Debug
//   -f float     Weight (default 60)
//   -l int       (REQ) Level
//   -n string    (REQ) Name
//
// $ $ go run flag/simpleformatusage/main.go -l 10 -n John
// [ Command options ]
//   -b         Birthday (format: 1900/01/02)
//   -d false   Debug
//   -f 60      Weight
//   -l 10      (REQ) Level
//   -n John    (REQ) Name
//

func main() {

	// Set custom usage (flag.VisitsAll を使ったUsageの加工方法、型やデフォルト値の表示ができなくなるので採用していない)
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(nil) }()
	usage := strings.Replace(strings.Replace(b.String(), ":", " [OPTIONS]\n\nDescription:\n  "+commandDescription+"\n\nOptions:\n", 1), "Usage of", "Usage:", 1)
	re := regexp.MustCompile(`[^,] +(-\S+)(?: (\S+))?\n*(\s+)(.*)\n`)
	flag.Usage = func() {
		_, _ = fmt.Fprint(flag.CommandLine.Output(), re.ReplaceAllStringFunc(usage, func(m string) string {
			return fmt.Sprintf("  %-"+commandOptionFieldWidth+"s %s\n", re.FindStringSubmatch(m)[1]+" "+strings.TrimSpace(re.FindStringSubmatch(m)[2]), re.FindStringSubmatch(m)[4])
		}))
	}

	flag.Parse()
	// Required parameter
	if *optionLevel == 0 || *optionName == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-9s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}
