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
	Req        = "(required)"
	UsageDummy = "########"
	TimeFormat = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	commandDescription = "A sample command demonstrating rich template usage in Go CLI applications."
	// Command options (the -h and --help flags are provided by default in the flag package)
	optionFilePath        = defineFlagValue("f", "file-path" /*    */, Color.Yellow(Req)+" File path" /*                   */, "" /*                         */, flag.String, flag.StringVar)
	optionUrl             = defineFlagValue("u", "url" /*          */, "URL" /*                                            */, "https://httpbin.org/get" /*  */, flag.String, flag.StringVar)
	optionLineIndex       = defineFlagValue("l", "line-index" /*   */, "Index of line" /*                                  */, 10 /*                         */, flag.Int, flag.IntVar)
	optionDurationWaitSec = defineFlagValue("w", "wait-seconds" /* */, "Duration of wait seconds (e.g., 500ms, 3s, 2m)" /* */, 1*time.Second /*              */, flag.Duration, flag.DurationVar)
	optionPrintSrc        = defineFlagValue("p", "print-src" /*    */, "Print source code" /*                              */, false /*                      */, flag.Bool, flag.BoolVar)
	optionDebug           = defineFlagValue("d", "debug" /*        */, "Debug mode" /*                                     */, false /*                      */, flag.Bool, flag.BoolVar)

	// Set environment variable
	environmentValueLoopCount, _ = strconv.Atoi(cmp.Or(os.Getenv("LOOP_COUNT"), "10"))

	// Color colorize string
	Color = struct {
		Red    func(string) string
		Green  func(string) string
		Yellow func(string) string
	}{
		// "\x1b[31m", "\x1b[32m", "\x1b[33m" = ANSI color codes, "\x1b[0m" = reset color code
		Red:    func(text string) string { return "\x1b[31m" + text + "\x1b[0m" },
		Green:  func(text string) string { return "\x1b[32m" + text + "\x1b[0m" },
		Yellow: func(text string) string { return "\x1b[33m" + text + "\x1b[0m" },
	}
)

func init() {
	// Customize the usage message
	flag.Usage = customUsage(commandDescription)
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
	optionsUsage, _ := getOptionsUsage(true)
	fmt.Printf("[ Command options ]\n%s\n", optionsUsage)

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

func HttpGetJson(url string) any {
	res, err := http.Get(url)
	handleError(err, "http.Get(url)")
	responseBody, err := io.ReadAll(res.Body)
	handleError(err, "io.ReadAll(res.Body)")
	return ToJsonObject(responseBody)
}

// =======================================
// Json Utils
// =======================================

// ToJsonObject json bytes to any object
func ToJsonObject(body []byte) any {
	var jsonObject any
	err := json.Unmarshal(body, &jsonObject)
	handleError(err, "json.Unmarshal")
	return jsonObject
}

// Get get value in any object [ example : object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb" ]
func Get(object any, keyChain string) any {
	var result any
	var exists bool
	for _, key := range strings.Split(keyChain, ".") {
		exists = false
		if _, ok := object.(map[string]any); ok {
			exists = true
			object = object.(map[string]any)[key]
			result = object
			continue
		}
		if values, ok := object.([]any); ok {
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
	f := flagFunc(long, defaultValue, flagUsage)
	flagVarFunc(f, short, defaultValue, UsageDummy)
	return f
}

// Custom usage message
func customUsage(description string) func() {
	return func() {
		optionsUsage, requiredOptionExample := getOptionsUsage(false)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s %s[OPTIONS]\n\n", func() string { e, _ := os.Executable(); return filepath.Base(e) }(), requiredOptionExample)
		fmt.Fprintf(flag.CommandLine.Output(), "Description:\n  %s\n\n", description)
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n%s", optionsUsage)
	}
}

// Get options usage message
func getOptionsUsage(currentValue bool) (string, string) {
	requiredOptionExample := ""
	optionNameWidth := 0
	usages := make([]string, 0)
	getType := func(v string) string {
		return strings.NewReplacer("*flag.boolValue", "", "*flag.", "<", "Value", ">").Replace(v)
		//return strings.NewReplacer("*flag.boolValue", "", "*flag.", "", "Value", "").Replace(v)
	}
	flag.VisitAll(func(f *flag.Flag) {
		optionNameWidth = max(optionNameWidth, len(fmt.Sprintf("%s %s", f.Name, getType(fmt.Sprintf("%T", f.Value))))+4)
	})
	flag.VisitAll(func(f *flag.Flag) {
		if f.Usage == UsageDummy {
			return
		}
		value := getType(fmt.Sprintf("%T", f.Value))
		if currentValue {
			value = f.Value.String()
		}
		short := strings.Split(f.Usage, UsageDummy)[0]
		mainUsage := strings.Split(f.Usage, UsageDummy)[1]
		if strings.Contains(mainUsage, Req) {
			requiredOptionExample += fmt.Sprintf("--%s %s ", f.Name, value)
		}
		usages = append(usages, fmt.Sprintf("  -%-1s, --%-"+strconv.Itoa(optionNameWidth)+"s %s\n", short, f.Name+" "+value, mainUsage))
	})
	sort.SliceStable(usages, func(i, j int) bool {
		return strings.Count(usages[i], Req) > strings.Count(usages[j], Req)
	})
	return strings.Join(usages, ""), requiredOptionExample
}
