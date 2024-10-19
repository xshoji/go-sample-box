package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// Set value by -ldflags "-X xxx"
	md5info string
	//go:embed main.go
	srcBytes []byte
	// Optional parameters
	paramsPrintMd5 = flag.Bool("print-md5", false, "[optional] Print md5 main.go")
	paramsPrintSrc = flag.Bool("print-src", false, "[optional] Print main.go")
	paramsPanic    = flag.Bool("p", false /* */, "\n[optional] Spawn panic")
	paramsHelp     = flag.Bool("h", false /* */, "\nhelp")
)

// Build
// ------
// $ APP="/tmp/app"; MAIN="main.go"; GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X 'main.md5info=$(md5 ${MAIN})'" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}
	if *paramsPrintMd5 {
		fmt.Println(md5info)
	}
	if *paramsPrintSrc {
		fmt.Printf("%s", srcBytes)
	}
	if *paramsPanic {
		log.Panic("Panic!")
	}
}
