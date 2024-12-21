package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	// Define options and description
	//
	optionLevel    = flag.Int("l" /*     */, 0 /*       */, "[req] level")
	optionName     = flag.String("n" /*  */, "" /*      */, "[req] name")
	optionBirthday = flag.String("b" /*  */, "" /*      */, "birthday (format: 1900/01/02)")
	optionWeight   = flag.Float64("f" /* */, 60.0 /*    */, "weight")
	optionHelp     = flag.Bool("h" /*    */, false /*   */, "\nhelp")
	optionDebug    = flag.Bool("d" /*    */, false /*   */, "\ndebug")
)

// << Execution sample >>
//
// $ go run main.go -l 2 -n John
// [req] Level :  2
// [req] Name  :  John
// Birthday    :
// Weight      :  60
// Debug       :  false
//
// $ go run cmd/simpleformatusage/main.go -h
// Usage: /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build3001085942/b001/exe/main [OPTIONS]
//
// Description:
//   Command description is here.
//
// Options:
//   -b string  birthday (format: 1900/01/02)
//   -d         debug
//   -f float   weight (default 60)
//   -h         help
//   -l int     [req] level
//   -n string  [req] name

func main() {

	// Set formatted usage
	description := "Command description is here."
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	usageLines := strings.Split(b.String(), "\n")
	usage := strings.Replace(strings.Replace(usageLines[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	re := regexp.MustCompile(" +(-\\S+)( *\\S*|\t)*\n(\\s+)(.*)\n")
	usage += re.ReplaceAllStringFunc(strings.Join(usageLines[1:], "\n"), func(m string) string {
		parts := re.FindStringSubmatch(m)
		return fmt.Sprintf("  %-10s %s\n", parts[1]+" "+strings.TrimSpace(parts[2]), parts[4])
	})
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage)
	}

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *optionHelp || *optionLevel == 0 || *optionName == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("[req] Level : ", *optionLevel)
	fmt.Println("[req] Name  : ", *optionName)
	fmt.Println("Birthday    : ", *optionBirthday)
	fmt.Println("Weight      : ", *optionWeight)
	fmt.Println("Debug       : ", *optionDebug)
}
