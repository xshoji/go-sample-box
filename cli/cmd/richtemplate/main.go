package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	UsageRequiredPrefix = "\u001B[33m[RQD]\u001B[0m "
	UsageDummy          = "########"
	TimeFormat          = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	// Command options
	commandDescription     = "Here is the command description."
	commandOptionMaxLength = 0
	optionFilePath         = defineFlagValue("f", "file-path" /*  */, UsageRequiredPrefix+"File path" /* */, "").(*string)
	optionUrl              = defineFlagValue("u", "url" /*        */, "URL" /*                           */, "https://httpbin.org/get").(*string)
	optionLineIndex        = defineFlagValue("l", "line-index" /* */, "Index of line" /*                 */, 10).(*int)
	optionPrintSrc         = defineFlagValue("p", "print-src" /*  */, "Print source code" /*             */, false).(*bool)
	optionHelp             = defineFlagValue("h", "help" /*       */, "Show help" /*                     */, false).(*bool)

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
	formatUsage(commandDescription, &commandOptionMaxLength, new(bytes.Buffer))
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
		if a.Usage == UsageDummy {
			return
		}
		fmt.Printf("-%-"+strconv.Itoa(commandOptionMaxLength)+"s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
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

// Helper function for flag
func defineFlagValue(short, long, description string, defaultValue any) (f any) {
	switch defaultValue.(type) {
	case string:
		f = flag.String(short, "", UsageDummy)
		flag.StringVar(f.(*string), long, defaultValue.(string), description)
	case int:
		f = flag.Int(short, 0, UsageDummy)
		flag.IntVar(f.(*int), long, defaultValue.(int), description)
	case bool:
		f = flag.Bool(short, false, UsageDummy)
		flag.BoolVar(f.(*bool), long, defaultValue.(bool), description)
	default:
		panic("unsupported flag type")
	}
	return
}

func formatUsage(description string, maxLength *int, buffer *bytes.Buffer) {
	// Get default flags usage
	func() { flag.CommandLine.SetOutput(buffer); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + ".*\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageFirst := strings.Replace(strings.Replace(strings.Split(buffer.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	usageOptions := re.FindAllString(buffer.String(), -1)
	for _, v := range usageOptions {
		*maxLength = max(*maxLength, len(re.ReplaceAllString(v, "$1, -$3$4")))
	}
	usageOptionsRep := make([]string, 0)
	for _, v := range usageOptions {
		usageOptionsRep = append(usageOptionsRep, fmt.Sprintf("%-6s%-"+strconv.Itoa(*maxLength)+"s", re.ReplaceAllString(v, "  $1,"), re.ReplaceAllString(v, "-$3$4"))+re.ReplaceAllString(v, "$5\n"))
	}
	sort.SliceStable(usageOptionsRep, func(i, j int) bool {
		return strings.Count(usageOptionsRep[i], UsageRequiredPrefix) > strings.Count(usageOptionsRep[j], UsageRequiredPrefix)
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usageFirst+strings.Join(usageOptionsRep, "")) }
}
