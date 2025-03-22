package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const (
	UsageRequiredPrefix = "\u001B[33m(REQ)\u001B[0m "
)

var (
	// Command options ( the -h, --help option is defined by default in the flag package )
	commandDescription      = "Here is the command description."
	commandOptionFieldWidth = "12" // recommended width = general: 12, bool only: 5
	optionName              = flag.String("n" /*  */, "" /*      */, UsageRequiredPrefix+"Name")
	optionWeight            = flag.Float64("f" /* */, 60.0 /*    */, "Weight")
)

// Build command [ OS = linux, darwin, windows ]
// $ go run main.go build linux /tmp/tool
func main() {

	// Build mode
	handleBuildMode()

	// Format usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	usage := strings.Replace(b.String(), ":", " [OPTIONS] [-h, --help]\n\nDescription:\n  "+commandDescription+"\n\nOptions:\n", 1)
	re := regexp.MustCompile(`[^,] +(-\S+)(?: (\S+))?\n*(\s+)(.*)\n`)
	flag.Usage = func() {
		_, _ = fmt.Fprint(flag.CommandLine.Output(), re.ReplaceAllStringFunc(usage, func(m string) string {
			return fmt.Sprintf("  %-"+commandOptionFieldWidth+"s %s\n", re.FindStringSubmatch(m)[1]+" "+strings.TrimSpace(re.FindStringSubmatch(m)[2]), re.FindStringSubmatch(m)[4])
		}))
	}

	flag.Parse()
	// Required parameter
	if *optionName == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-9s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
}

// =======================================
// Dev tools
// =======================================

func handleBuildMode() {
	if len(os.Args) < 2 || os.Args[1] != "build" {
		return
	}
	targetOs := os.Args[2]
	outputPath := os.Args[3]
	_, sourcePath, _, _ := runtime.Caller(1)
	fmt.Println(sourcePath)
	cmd := exec.Command("go", "build", `-ldflags=-s -w`, "-trimpath", "-o", outputPath, sourcePath)
	env := os.Environ()
	env = append(env, "GOOS="+targetOs, "GOARCH=amd64")
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Build completed successfully!")
	os.Exit(0)
}
