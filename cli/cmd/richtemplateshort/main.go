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
	"strconv"
	"strings"
	"time"
)

const (
	Req        = "(required)"
	TimeFormat = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	commandDescription           = "A sample command demonstrating rich template (short only) usage in Go CLI applications."
	commandOptionMaxLength       = "12" // Recommended width = general: 12, bool only: 5
	commandRequiredOptionExample = ""   // Auto-adjusted in defineFlagValue
	// Command options (the -h and --help flags are provided by default in the flag package)
	optionFilePath        = defineFlagValue("f", Color.Yellow(Req)+" File path" /*                    */, "" /*                         */, flag.String)
	optionUrl             = defineFlagValue("u", "URL" /*                                             */, "https://httpbin.org/get" /*  */, flag.String)
	optionLineIndex       = defineFlagValue("l", "Index of line" /*                                   */, 10 /*                         */, flag.Int)
	optionDurationWaitSec = defineFlagValue("w", "Duration of wait seconds (e.g., 1s, 500ms, 2m)" /*  */, 1*time.Second /*              */, flag.Duration)
	optionPrintSrc        = defineFlagValue("p", "Print source code" /*                               */, false /*                      */, flag.Bool)

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
	flag.Usage = customUsage(os.Stdout, commandDescription, commandOptionMaxLength)
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
	fmt.Printf("[ Command options ]\n")
	printOptionsUsage(commandOptionMaxLength, true)

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
func defineFlagValue[T comparable](flagName, description string, defaultValue T, flagFunc func(name string, value T, usage string) *T) *T {
	var zero T
	if defaultValue != zero {
		description = description + fmt.Sprintf(" (default %v)", defaultValue)
	}
	if strings.Contains(description, Req) {
		commandRequiredOptionExample = commandRequiredOptionExample + fmt.Sprintf("-%s %T ", flagName, defaultValue)
	}
	return flagFunc(flagName, defaultValue, description)
}

// Custom usage message
func customUsage(output io.Writer, description, fieldWidth string) func() {
	return func() {
		fmt.Fprintf(output, "Usage: %s %s[OPTIONS]\n\n", func() string { e, _ := os.Executable(); return filepath.Base(e) }(), commandRequiredOptionExample)
		fmt.Fprintf(output, "Description:\n  %s\n\n", description)
		fmt.Fprintf(output, "Options:\n")
		printOptionsUsage(fieldWidth, false)
	}
}

// Print options usage message
func printOptionsUsage(fieldWidth string, currentValue bool) {
	flag.VisitAll(func(f *flag.Flag) {
		value := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%T", f.Value), "*flag.", ""), "Value", ""), "bool", "")
		if currentValue {
			value = f.Value.String()
		}
		format := "  -%-" + fieldWidth + "s %s\n"
		fmt.Printf(format, f.Name+" "+value, f.Usage)
	})
}
