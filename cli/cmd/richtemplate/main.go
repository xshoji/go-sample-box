package main

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Req        = "(REQ)"
	UsageDummy = "########"
	TimeFormat = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	commandDescription     = "A sample command demonstrating rich template usage in Go CLI applications."
	commandOptionMaxLength = 0
	// Command options (the -h and --help flags are provided by default in the flag package)	commandDescription     = "Here is the command description."
	optionFilePath        = defineFlagValue("f", "file-path" /*    */, Color.Yellow(Req)+" File path" /*  */, "" /*                         */, flag.String, flag.StringVar)
	optionUrl             = defineFlagValue("u", "url" /*          */, "URL" /*                           */, "https://httpbin.org/get" /*  */, flag.String, flag.StringVar)
	optionLineIndex       = defineFlagValue("l", "line-index" /*   */, "Index of line" /*                 */, 10 /*                         */, flag.Int, flag.IntVar)
	optionDurationWaitSec = defineFlagValue("w", "wait-seconds" /* */, "Duration of wait seconds" /*      */, 1*time.Second /*              */, flag.Duration, flag.DurationVar)
	optionPrintSrc        = defineFlagValue("p", "print-src" /*    */, "Print source code" /*             */, false /*                      */, flag.Bool, flag.BoolVar)

	// Set environment variable
	environmentValueLoopCount, _ = strconv.Atoi(cmp.Or(os.Getenv("LOOP_COUNT"), "10"))

	// Color colorize string
	Color = struct {
		Red    func(string) string
		Green  func(string) string
		Yellow func(string) string
	}{
		// "\033[31m", "\033[32m", "\033[33m" = ANSI color codes, "\033[0m" = reset color code
		Red:    func(text string) string { return "\033[31m" + text + "\033[0m" },
		Green:  func(text string) string { return "\033[32m" + text + "\033[0m" },
		Yellow: func(text string) string { return "\033[33m" + text + "\033[0m" },
	}
)

func init() {
	// Customize the usage message
	flag.Usage = customUsage(os.Stdout, commandDescription, strconv.Itoa(commandOptionMaxLength))
}

// Build:
// $ GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /tmp/tool main.go
// $ go build -ldflags="-s -w" -trimpath -o /tmp/$(basename "${PWD}") main.go
func main() {

	flag.Parse()
	if *optionPrintSrc {
		fmt.Printf("%s", srcBytes)
		os.Exit(0)
	}
	if *optionFilePath == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("[ Environment variable ]\nLOOP_COUNT: %d\n\n", environmentValueLoopCount)
	fmt.Printf("[ Command options ]\n%s\n", getOptionsUsage(strconv.Itoa(commandOptionMaxLength), true))

	time.Sleep(*optionDurationWaitSec)
	contents := ReadAllFileContents(optionFilePath)
	fmt.Println(strings.Split(contents, "\n")[*optionLineIndex])

	res := HttpGetJson(*optionUrl + "?" + fmt.Sprintf("optionLineIndex=%d", *optionLineIndex))
	fmt.Println(res)
	fmt.Println(Get(res, "args.optionLineIndex"))

	fmt.Println(Color.Red("Sample: Error"))
	fmt.Println(Color.Yellow("Sample: Warning"))
	fmt.Println(Color.Green("Sample: Success"))

}

// =======================================
// HTTP Utils
// =======================================

func HttpGetJson(url string) interface{} {
	res, err := http.Get(url)
	handleError(err, "http.Get(url)")
	responseBody, err := io.ReadAll(res.Body)
	handleError(err, "io.ReadAll(res.Body)")
	return ToJsonObject(responseBody)
}

// =======================================
// Json Utils
// =======================================

// ToJsonObject json bytes to interface{} object
func ToJsonObject(body []byte) interface{} {
	var jsonObject interface{}
	err := json.Unmarshal(body, &jsonObject)
	handleError(err, "json.Unmarshal")
	return jsonObject
}

// Get get value in interface{} object [ example : object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb" ]
func Get(object interface{}, keyChain string) interface{} {
	var result interface{}
	var exists bool
	for _, key := range strings.Split(keyChain, ".") {
		exists = false
		if _, ok := object.(map[string]interface{}); ok {
			exists = true
			object = object.(map[string]interface{})[key]
			result = object
			continue
		}
		if values, ok := object.([]interface{}); ok {
			for i, v := range values {
				if strconv.FormatInt(int64(i), 10) == key {
					exists = true
					object = v
					result = object
					continue
				}
			}
		}
	}
	if exists {
		return result
	}
	return nil
}

// =======================================
// File Utils
// =======================================

func ReadAllFileContents(filePath *string) string {
	file, err := os.Open(*filePath)
	handleError(err, "os.Open(*filePath)")
	defer func() { handleError(file.Close(), "file.Close()") }()

	contents, err := io.ReadAll(file)
	handleError(err, "io.ReadAll(file)")
	return string(contents)
}

// =======================================
// Common Utils
// =======================================

func handleError(err error, prefixErrMessage string) {
	if err != nil {
		fmt.Printf("%s [ERROR %s]: %v\n", time.Now().Format(TimeFormat), prefixErrMessage, err)
	}
}

// =======================================
// flag Utils
// =======================================

// Helper function for flag
func defineFlagValue[T comparable](short, long, description string, defaultValue T, flagFunc func(name string, value T, usage string) *T, flagVarFunc func(p *T, name string, value T, usage string)) *T {
	flagUsage := short + UsageDummy + description
	var zero T
	if defaultValue != zero {
		flagUsage = flagUsage + fmt.Sprintf(" (default %v)", defaultValue)
	}
	commandOptionMaxLength = max(commandOptionMaxLength, len(fmt.Sprintf("  -%s", short))+len(fmt.Sprintf(", --%s", long))+6)
	f := flagFunc(long, defaultValue, flagUsage)
	flagVarFunc(f, short, defaultValue, UsageDummy)
	return f
}

// Custom usage message
func customUsage(output io.Writer, description, fieldWidth string) func() {
	return func() {
		fmt.Fprintf(output, "Usage: %s [OPTIONS] [-h, --help]\n\n", func() string { e, _ := os.Executable(); return filepath.Base(e) }())
		fmt.Fprintf(output, "Description:\n  %s\n\n", description)
		fmt.Fprintf(output, "Options:\n%s", getOptionsUsage(fieldWidth, false))
	}
}

// Get options usage message
func getOptionsUsage(fieldWidth string, currentValue bool) string {
	optionUsages := make([]string, 0)
	flag.VisitAll(func(f *flag.Flag) {
		if f.Usage == UsageDummy {
			return
		}
		value := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%T", f.Value), "*flag.", ""), "Value", ""), "bool", "")
		if currentValue {
			value = f.Value.String()
		}
		format := "  -%-1s, --%-" + fieldWidth + "s %s\n"
		short := strings.Split(f.Usage, UsageDummy)[0]
		mainUsage := strings.Split(f.Usage, UsageDummy)[1]
		optionUsages = append(optionUsages, fmt.Sprintf(format, short, f.Name+" "+value, mainUsage))
	})
	sort.SliceStable(optionUsages, func(i, j int) bool {
		return strings.Count(optionUsages[i], Req) > strings.Count(optionUsages[j], Req)
	})
	return strings.Join(optionUsages, "")
}
