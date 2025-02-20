package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
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
	commandDescription      = "Output zip file tool."
	commandOptionFieldWidth = 12
	optionOutputPath        = flag.String("o" /*  */, "" /*    */, UsageRequiredPrefix+"Output path of dummy zip file")
	optionByteSize          = flag.Int("b" /*     */, 1024 /*  */, "Byte size for zip file")
	optionHelp              = flag.Bool("h" /*    */, false /* */, "Help")
)

func init() {
	formatUsage(commandDescription, commandOptionFieldWidth)
}

// # Build: APP="/tmp/tool"; MAIN="main.go"; GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *optionHelp || *optionOutputPath == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-30s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
	fmt.Printf("\n\n")

	zipFileDummy := CreateZipFileOnMemory(*optionByteSize)
	fmt.Printf("CreateZipFileOnMemory:\n")
	fmt.Printf("%v\n\n", zipFileDummy.Bytes())

	CreateZipFileConcrete(*optionOutputPath, *optionByteSize)
	fmt.Printf("CreateZipFileConcrete:\n")
	fmt.Printf("Path: %s\n", *optionOutputPath)
}

type ConstantDataUnbufferedReader struct {
	chunkSize          int
	repetitionsCurrent int
	repetitionsMax     int
	remainingByteSize  int
}

func NewConstantDataUnbufferedReader(byteSize int) *ConstantDataUnbufferedReader {
	chunkByteSize := 1024
	return &ConstantDataUnbufferedReader{
		chunkSize:          chunkByteSize,
		repetitionsCurrent: 0,
		repetitionsMax:     byteSize/chunkByteSize + 1,
		remainingByteSize:  byteSize % chunkByteSize,
	}
}

func (r *ConstantDataUnbufferedReader) Read(p []byte) (n int, err error) {
	if r.repetitionsCurrent >= r.repetitionsMax {
		return 0, io.EOF
	}
	chunkSize := r.chunkSize
	if r.repetitionsCurrent == r.repetitionsMax-1 {
		chunkSize = r.remainingByteSize
	}
	if chunkSize != 0 {
		copy(p, bytes.Repeat([]byte("0"), chunkSize))
	}
	r.repetitionsCurrent++
	return chunkSize, nil
}
func CreateZipFileOnMemory(optionByteSize int) *bytes.Buffer {
	var dummyFileBytes []byte
	dummyZipFile := bytes.NewBuffer(dummyFileBytes)
	createZipFile(dummyZipFile, optionByteSize)
	return dummyZipFile

}

func CreateZipFileConcrete(optionOutputPath string, optionByteSize int) {
	dummyZipFile, err := os.Create(optionOutputPath)
	handleError(err, `os.Create(*optionOutputPath)`)
	defer dummyZipFile.Close()
	createZipFile(dummyZipFile, optionByteSize)
}

func createZipFile(dummyZipFile io.Writer, optionByteSize int) {
	zipWriter := zip.NewWriter(dummyZipFile)
	w1, err := zipWriter.Create("dummy.tsv")
	handleError(err, `zipWriter.Create("dir1/dummy.tsv")`)
	constantDataUnbufferedReader := NewConstantDataUnbufferedReader(optionByteSize)
	_, err = io.Copy(w1, constantDataUnbufferedReader)
	handleError(err, `io.Copy(w1, constantDataUnbufferedReader)`)
	err = zipWriter.Close()
	handleError(err, `zipWriter.Close()`)
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
	usage := strings.Replace(strings.Replace(usageLines[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	re := regexp.MustCompile(` +(-\S+)(?: (\S+))?\n*(\s+)(.*)\n`)
	usage += re.ReplaceAllStringFunc(strings.Join(usageLines[1:], "\n"), func(m string) string {
		parts := re.FindStringSubmatch(m)
		return fmt.Sprintf("  %-"+strconv.Itoa(optionFieldWidth)+"s %s\n", parts[1]+" "+strings.TrimSpace(parts[2]), parts[4])
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}
