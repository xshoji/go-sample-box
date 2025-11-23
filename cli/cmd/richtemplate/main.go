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
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	UsageRequiredPrefix = "\033[33m" + "(REQ)" + "\033[0m "
	UsageDummy          = "########"
	TimeFormat          = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	// Command options (the -h and --help flags are provided by default in the flag package)	commandDescription     = "Here is the command description."
	commandDescription     = "A sample command demonstrating rich template usage in Go CLI applications."
	commandOptionMaxLength = "18"
	optionFilePath         = defineFlagValue("f", "file-path" /*  */, UsageRequiredPrefix+"File path" /* */, "").(*string)
	optionUrl              = defineFlagValue("u", "url" /*        */, "URL" /*                           */, "https://httpbin.org/get").(*string)
	optionLineIndex        = defineFlagValue("l", "line-index" /* */, "Index of line" /*                 */, 10).(*int)
	optionPrintSrc         = defineFlagValue("p", "print-src" /*  */, "Print source code" /*             */, false).(*bool)

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
	flag.Usage = customUsage(os.Stdout, os.Args[0], commandDescription, commandOptionMaxLength)
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
	flag.VisitAll(func(a *flag.Flag) {
		if a.Usage == UsageDummy {
			return
		}
		fmt.Printf("  -%-30s %s\n",
			fmt.Sprintf("%-1s, -%s %v", strings.Split(a.Usage, UsageDummy)[0], a.Name, a.Value),
			strings.Trim(strings.Split(a.Usage, UsageDummy)[1], "\n"))
	})
	fmt.Printf("\n\n")

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

// Helper function for flag
func defineFlagValue(short, long, description string, defaultValue any) (f any) {
	flagUsage := short + UsageDummy + description
	defaultValueDescription := ""
	switch v := defaultValue.(type) {
	case bool:
		f = flag.Bool(short, false, UsageDummy)
		flag.BoolVar(f.(*bool), long, v, flagUsage)
	case string:
		var d string
		if d != defaultValue.(string) {
			defaultValueDescription = fmt.Sprintf(" (default %s)", defaultValue.(string))
		}
		f = flag.String(short, "", UsageDummy)
		flag.StringVar(f.(*string), long, v, flagUsage+defaultValueDescription)
	case int:
		var d int
		if d != defaultValue.(int) {
			defaultValueDescription = fmt.Sprintf(" (default %d)", defaultValue.(int))
		}
		f = flag.Int(short, 0, UsageDummy)
		flag.IntVar(f.(*int), long, v, flagUsage+defaultValueDescription)
	case float64:
		var d float64
		if d != defaultValue.(float64) {
			defaultValueDescription = fmt.Sprintf(" (default %f)", defaultValue.(float64))
		}
		f = flag.Float64(short, 0.0, UsageDummy)
		flag.Float64Var(f.(*float64), long, v, flagUsage+defaultValueDescription)
	default:
		panic("unsupported flag type")
	}
	return
}

func customUsage(output io.Writer, cmdName, description, fieldWidth string) func() {
	return func() {
		fmt.Fprintf(output, "Usage: %s [OPTIONS] [-h, --help]\n\n", cmdName)
		fmt.Fprintf(output, "Description:\n  %s\n\n", description)
		fmt.Fprintln(output, "Options:")

		optionUsages := make([]string, 0)
		flag.VisitAll(func(f *flag.Flag) {
			if f.Usage == UsageDummy {
				return
			}
			valueType := strings.Replace(strings.Replace(fmt.Sprintf("%T", f.Value), "*flag.", "", -1), "Value", "", -1)
			format := "  -%s, --%-" + fieldWidth + "s %s\n"
			short := strings.Split(f.Usage, UsageDummy)[0]
			mainUsage := strings.Split(f.Usage, UsageDummy)[1]
			optionUsages = append(optionUsages, fmt.Sprintf(format, short, f.Name+" "+valueType, mainUsage))
		})
		sort.SliceStable(optionUsages, func(i, j int) bool {
			return strings.Count(optionUsages[i], UsageRequiredPrefix) > strings.Count(optionUsages[j], UsageRequiredPrefix)
		})
		fmt.Fprint(output, strings.Join(optionUsages, ""))
	}
}
