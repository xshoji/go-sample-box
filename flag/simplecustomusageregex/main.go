package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	Req = "\x1b[33m(required)\x1b[0m "
)

var (
	// Command options ( the -h, --help option is defined by default in the flag package )
	commandDescription      = "Here is the command description."
	commandOptionFieldWidth = "12" // recommended width = general: 12, bool only: 5
	optionFilePath          = flag.String("f" /*   */, "" /*                      */, Req+"File path")
	optionUrl               = flag.String("u" /*   */, "https://example/path" /*  */, "URL")
	optionLineIndex         = flag.Int("l" /*      */, 10 /*                      */, "Index of line")
	optionDebug             = flag.Bool("d" /*     */, false /*                   */, "\nEnable debug mode")
	optionDurationWaitSec   = flag.Duration("w" /* */, 1*time.Second /*           */, "Duration of wait seconds (e.g., 1s, 500ms, 2m)")
)

// << Execution sample >>
//
// Usage: /var/folders/vv/fc_khb3904d2d0_rz86jxhv00000gn/T/go-build3849862186/b001/exe/main [OPTIONS]
//
// Description:
//   Here is the command description.
//
// Options:
//   -d           Enable debug mode
//   -f string    (required) File path
//   -l int       Index of line (default 10)
//   -u string    URL (default "https://example/path")
//   -w duration  Duration of wait seconds (e.g., 1s, 500ms, 2m) (default 1s)
//
// $ go run main.go -f main.go
// [ Command options ]
//   -d false      Enable debug mode
//   -f main.go    (required) File path
//   -l 10         Index of line
//   -u https://example/path URL
//   -w 1s         Duration of wait seconds (e.g., 1s, 500ms, 2m)
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
	if *optionFilePath == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-"+commandOptionFieldWidth+"s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}
