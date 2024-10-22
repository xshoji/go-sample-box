package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	CommandDescription  = "Command description is here."
	UsageRequiredPrefix = "\u001B[33m[required]\u001B[0m "
	UsageDummy          = "########"
	TimeFormat          = "2006-01-02 15:04:05.9999 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	// Define short parameters ( don't set default value ).
	paramsFilePath  = flag.String("f", "", UsageDummy)
	paramsUrl       = flag.String("u", "", UsageDummy)
	paramsLineIndex = flag.Int("l", 0, UsageDummy)
	paramsPrintSrc  = flag.Bool("p", false, UsageDummy)
	paramsHelp      = flag.Bool("h", false, UsageDummy)

	// ColorPrinter colorize string
	ColorPrinter = struct {
		Red      string
		Green    string
		Yellow   string
		Colorize func(string, string) string
	}{
		Red:    "\033[31m",
		Green:  "\033[32m",
		Yellow: "\033[33m",
		Colorize: func(color string, text string) string {
			if runtime.GOOS == "windows" {
				return text
			}
			colorReset := "\033[0m"
			return color + text + colorReset
		},
	}
)

func init() {
	// Define long parameters and description ( set default value here if you need ).
	//
	// Required parameters
	flag.StringVar(paramsFilePath /* */, "file-path" /*  */, "" /*    */, UsageRequiredPrefix+"file path")
	// Optional parameters
	flag.StringVar(paramsUrl /*      */, "url" /*        */, "https://httpbin.org/get" /*    */, "URL")
	flag.IntVar(paramsLineIndex /*   */, "line-index" /* */, 10 /*                           */, "index of line")
	flag.BoolVar(paramsPrintSrc /*   */, "print-src" /*  */, false /*                        */, "print source code")
	flag.BoolVar(paramsHelp /*       */, "help" /*       */, false /*                        */, "help")

	adjustUsage()
}

// # Build: APP="/tmp/tool"; MAIN="main.go"; GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *paramsPrintSrc {
		fmt.Printf("%s", srcBytes)
		os.Exit(0)
	}
	if *paramsHelp || *paramsFilePath == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println("file-path:", *paramsFilePath)
	fmt.Println("line-index:", *paramsLineIndex)

	contents := readAllFileContents(paramsFilePath)
	fmt.Println(strings.Split(contents, "\n")[*paramsLineIndex])

	res := HttpGetJson(*paramsUrl + "?" + fmt.Sprintf("paramsLineIndex=%d", *paramsLineIndex))
	fmt.Println(res)
	fmt.Println(Get(res, "args.paramsLineIndex"))

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

func readAllFileContents(filePath *string) string {
	file, err := os.Open(*filePath)
	handleError(err, "os.Open(*filePath)")
	defer createFileCloseDeferFunc(file)()

	contents, err := io.ReadAll(file)
	handleError(err, "io.ReadAll(file)")
	return string(contents)
}

func createFileCloseDeferFunc(file *os.File) func() {
	return func() {
		if err := file.Close(); err != nil {
			log.Panic(err)
		}
	}
}

// =======================================
// Common Utils
// =======================================

// Handle error
func handleError(err error, prefixErrMessage string) {
	if err != nil {
		fmt.Printf("%s [ERROR %s]: %v\n", time.Now().Format(TimeFormat), prefixErrMessage, err)
	}
}

func adjustUsage() {
	// Get default flags usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	// Get default flags usage
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + "\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageParams := re.FindAllString(b.String(), -1)
	maxLengthParam := 0.0
	sort.Slice(usageParams, func(i, j int) bool {
		maxLengthParam = math.Max(maxLengthParam, math.Max(float64(len(re.ReplaceAllString(usageParams[i], "$1, -$3$4"))), float64(len(re.ReplaceAllString(usageParams[j], "$1, -$3$4")))))
		if len(strings.Split(usageParams[i]+usageParams[j], UsageRequiredPrefix))%2 == 1 {
			return strings.Compare(usageParams[i], usageParams[j]) == -1
		} else {
			return strings.Index(usageParams[i], UsageRequiredPrefix) >= 0
		}
	})
	usage := strings.Replace(strings.Replace(strings.Split(b.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + CommandDescription + "\n\nOptions:\n"
	for _, v := range usageParams {
		usage += fmt.Sprintf("%-"+strconv.Itoa(int(maxLengthParam+4.0))+"s", re.ReplaceAllString(v, "  $1, -$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
