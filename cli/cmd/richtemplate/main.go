package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	CommandDescription  = "Here is the command description."
	UsageRequiredPrefix = "\u001B[33m[required]\u001B[0m "
	UsageDummy          = "########"
	TimeFormat          = "2006-01-02 15:04:05.9999 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	// Required flag
	optionFilePath = defineStringFlag("f", "file-path", UsageRequiredPrefix+"file path", "")

	// Optional flag
	optionUrl       = defineStringFlag("u", "url", "url", "https://httpbin.org/get")
	optionLineIndex = defineIntFlag("l", "line-index", "index of line", 10)
	optionPrintSrc  = defineBoolFlag("p", "print-src", "print source code")
	optionHelp      = defineBoolFlag("h", "help", "help")

	// Set environment variable
	environmentValueLoopCount, _ = strconv.Atoi(GetEnvOrDefault("LOOP_COUNT", "10"))

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
	formatUsage()
}

// # Build: APP="/tmp/tool"; MAIN="main.go"; GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *optionPrintSrc {
		fmt.Printf("%s", srcBytes)
		os.Exit(0)
	}
	if *optionHelp || *optionFilePath == "" {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Printf("[ Environment variable ]\nLOOP_COUNT: %d\n\n", environmentValueLoopCount)
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("-%-35s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
	fmt.Printf("\n\n")

	contents := ReadAllFileContents(optionFilePath)
	fmt.Println(strings.Split(contents, "\n")[*optionLineIndex])

	res := HttpGetJson(*optionUrl + "?" + fmt.Sprintf("optionLineIndex=%d", *optionLineIndex))
	fmt.Println(res)
	fmt.Println(Get(res, "args.optionLineIndex"))

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

func GetEnvOrDefault(key string, defaultValue string) string {
	value := defaultValue
	v := os.Getenv(key)
	if v != "" {
		value = v
	}
	return value
}

func handleError(err error, prefixErrMessage string) {
	if err != nil {
		fmt.Printf("%s [ERROR %s]: %v\n", time.Now().Format(TimeFormat), prefixErrMessage, err)
	}
}

func defineStringFlag(short, long, description, defaultValue string) (v *string) {
	// Define short parameters ( this default value and usage will be not used ).
	v = flag.String(short, "", UsageDummy)
	// Define long parameters and description ( set default value here if you need ).
	flag.StringVar(v, long, defaultValue, description)
	return
}

func defineIntFlag(short, long, description string, defaultValue int) (v *int) {
	v = flag.Int(short, 0, UsageDummy)
	flag.IntVar(v, long, defaultValue, description)
	return
}

func defineBoolFlag(short, long, description string) (v *bool) {
	v = flag.Bool(short, false, UsageDummy)
	flag.BoolVar(v, long, false, description)
	return
}

func formatUsage() {
	// Get default flags usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	// Get default flags usage
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + ".*\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageOptions := re.FindAllString(b.String(), -1)
	maxLength := 0.0
	sort.Slice(usageOptions, func(i, j int) bool {
		maxLength = math.Max(maxLength, math.Max(float64(len(re.ReplaceAllString(usageOptions[i], "$1, -$3$4"))), float64(len(re.ReplaceAllString(usageOptions[j], "$1, -$3$4")))))
		if len(strings.Split(usageOptions[i]+usageOptions[j], UsageRequiredPrefix))%2 == 1 {
			return strings.Compare(usageOptions[i], usageOptions[j]) == -1
		} else {
			return strings.Index(usageOptions[i], UsageRequiredPrefix) >= 0
		}
	})
	usage := strings.Replace(strings.Replace(strings.Split(b.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + CommandDescription + "\n\nOptions:\n"
	for _, v := range usageOptions {
		usage += fmt.Sprintf("%-6s%-"+strconv.Itoa(int(maxLength))+"s", re.ReplaceAllString(v, "  $1,"), re.ReplaceAllString(v, "-$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
