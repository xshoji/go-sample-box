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
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	UsageRequiredPrefix = "\u001B[33m[RQD]\u001B[0m "
	TimeFormat          = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	//go:embed main.go
	srcBytes []byte

	// Command options
	commandDescription      = "Here is the command description."
	commandOptionFieldWidth = 12
	optionFilePath          = flag.String("f" /*  */, "" /*                         */, UsageRequiredPrefix+"File path")
	optionUrl               = flag.String("u" /*  */, "https://httpbin.org/get" /*  */, "URL")
	optionLineIndex         = flag.Int("l" /*     */, 10 /*                         */, "Index of line")
	optionPrintSrc          = flag.Bool("p" /*    */, false /*                      */, "Print main.go")
	optionHelp              = flag.Bool("h" /*    */, false /*                      */, "Help")

	// Set environment variable
	environmentValueLoopCount, _ = strconv.Atoi(GetEnvOrDefault("LOOP_COUNT", "10"))
)

func init() {
	formatUsage(commandDescription, commandOptionFieldWidth)
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

	fmt.Printf("[ Environment variable ]\n")
	fmt.Printf("  LOOP_COUNT: %d\n\n", environmentValueLoopCount)
	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-30s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
	fmt.Printf("\n\n")

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

// formatUsage optionFieldWidth [ recommended width = general: 12, bool only: 5 ]
func formatUsage(description string, optionFieldWidth int) {
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	usageLines := strings.Split(b.String(), "\n")
	usageFirst := strings.Replace(strings.Replace(usageLines[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	re := regexp.MustCompile(` +(-\S+)(?: (\S+))?\n*(\s+)(.*)\n`)
	usageOptions := strings.Split(re.ReplaceAllStringFunc(strings.Join(usageLines[1:], "\n"), func(m string) string {
		return fmt.Sprintf("  %-"+strconv.Itoa(optionFieldWidth)+"s %s\n", re.FindStringSubmatch(m)[1]+" "+strings.TrimSpace(re.FindStringSubmatch(m)[2]), re.FindStringSubmatch(m)[4])
	}), "\n")
	sort.SliceStable(usageOptions, func(i, j int) bool {
		return strings.Count(usageOptions[i], UsageRequiredPrefix) > strings.Count(usageOptions[j], UsageRequiredPrefix)
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usageFirst+strings.Join(usageOptions, "\n")) }
}
