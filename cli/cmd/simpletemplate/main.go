package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Req        = "\u001B[33m(REQ)\u001B[0m "
	TimeFormat = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	commandDescription      = "A sample command demonstrating simple template usage in Go CLI applications."
	commandOptionFieldWidth = "12" // Recommended width = general: 12, bool only: 5
	// Command options (the -h and --help flags are provided by default in the flag package)	commandDescription     = "Here is the command description."
	optionFilePath        = flag.String("f" /*   */, "" /*                         */, Req+"File path")
	optionUrl             = flag.String("u" /*   */, "https://httpbin.org/get" /*  */, "URL")
	optionLineIndex       = flag.Int("l" /*      */, 10 /*                         */, "Index of line")
	optionDebug           = flag.Bool("d" /*     */, false /*                      */, "Debug mode")
	optionDurationWaitSec = flag.Duration("w" /* */, 1*time.Second /*              */, "Duration of wait seconds")

	// Set environment variable
	environmentValueLoopCount, _ = strconv.Atoi(cmp.Or(os.Getenv("LOOP_COUNT"), "10"))
)

func init() {
	flag.Usage = customUsage(new(bytes.Buffer), commandDescription, commandOptionFieldWidth)
}

// Build:
// $ GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /tmp/tool main.go
// $ go build -ldflags="-s -w" -trimpath -o /tmp/$(basename "${PWD}") main.go
func main() {

	flag.Parse()
	if *optionFilePath == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("[ Environment variable ]\n")
	fmt.Printf("  LOOP_COUNT: %d\n\n", environmentValueLoopCount)
	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-30s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
	fmt.Printf("\n\n")

	if *optionDebug {
		fmt.Printf("[ Debug ] debug mode: true\n\n")
	}

	time.Sleep(*optionDurationWaitSec)
	contents := ReadAllFileContents(optionFilePath)
	fmt.Println(strings.Split(contents, "\n")[*optionLineIndex])

	res := HttpGetJson(*optionUrl + "?" + fmt.Sprintf("optionLineIndex=%d", *optionLineIndex))
	fmt.Println(res)
	fmt.Println(res.(map[string]any)["args"].(map[string]any)["optionLineIndex"])

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

func customUsage(b *bytes.Buffer, description string, optionFieldWidth string) func() {
	func() { flag.CommandLine.SetOutput(b); flag.PrintDefaults(); flag.CommandLine.SetOutput(nil) }()
	return func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS] [-h, --help]\n\n", func() string { e, _ := os.Executable(); return filepath.Base(e) }())
		fmt.Fprintf(flag.CommandLine.Output(), "Description:\n  %s\n\n", description)
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		re := regexp.MustCompile(`(?m)^ +(-\S+)(?: (\S+))?\n*(\s+)(.*)\n`)
		fmt.Fprint(flag.CommandLine.Output(), re.ReplaceAllStringFunc(b.String(), func(m string) string {
			return fmt.Sprintf("  %-"+optionFieldWidth+"s %s\n", re.FindStringSubmatch(m)[1]+" "+strings.TrimSpace(re.FindStringSubmatch(m)[2]), re.FindStringSubmatch(m)[4])
		}))
	}
}
