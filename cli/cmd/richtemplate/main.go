package main

import (
	"bytes"
	"cmp"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
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

	// Command options ( the -h, --help option is defined by default in the flag package )
	commandDescription     = "Here is the command description."
	commandOptionMaxLength = 0
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
	formatUsage(commandDescription, &commandOptionMaxLength, new(bytes.Buffer))
}

// Build:
// $ GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /tmp/tool main.go
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
	switch v := defaultValue.(type) {
	case string:
		f = flag.String(short, "", UsageDummy)
		flag.StringVar(f.(*string), long, v, flagUsage)
	case int:
		f = flag.Int(short, 0, UsageDummy)
		flag.IntVar(f.(*int), long, v, flagUsage)
	case bool:
		f = flag.Bool(short, false, UsageDummy)
		flag.BoolVar(f.(*bool), long, v, flagUsage)
	case float64:
		f = flag.Float64(short, 0.0, UsageDummy)
		flag.Float64Var(f.(*float64), long, v, flagUsage)
	default:
		panic("unsupported flag type")
	}
	return
}

func formatUsage(description string, maxLength *int, buffer *bytes.Buffer) {
	func() { flag.CommandLine.SetOutput(buffer); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	usageOption := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+"+UsageDummy+"\n\\s*").ReplaceAllString(buffer.String(), "")
	re := regexp.MustCompile("\\s(-\\S+)( *\\S*)( *\\S*)+\n\\s+(.+)")
	usageFirst := strings.Replace(strings.Replace(strings.Split(usageOption, "\n")[0], ":", " [OPTIONS] [-h, --help]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	usageOptions := re.FindAllString(usageOption, -1)
	for _, v := range usageOptions {
		*maxLength = max(*maxLength, len(re.ReplaceAllString(v, " -$1")+re.ReplaceAllString(v, "$2"))+2)
	}
	usageOptionsRep := make([]string, 0)
	for _, v := range usageOptions {
		usageOptionsRep = append(usageOptionsRep, fmt.Sprintf("  -%-1s,%-"+strconv.Itoa(*maxLength)+"s%s", strings.Split(re.ReplaceAllString(v, "$4"), UsageDummy)[0], re.ReplaceAllString(v, " -$1")+re.ReplaceAllString(v, "$2"), strings.Split(re.ReplaceAllString(v, "$4"), UsageDummy)[1]+"\n"))
	}
	sort.SliceStable(usageOptionsRep, func(i, j int) bool {
		return strings.Count(usageOptionsRep[i], UsageRequiredPrefix) > strings.Count(usageOptionsRep[j], UsageRequiredPrefix)
	})
	flag.Usage = func() { _, _ = fmt.Fprint(flag.CommandLine.Output(), usageFirst+strings.Join(usageOptionsRep, "")) }
}
