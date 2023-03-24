package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// Set value by -ldflags "-X xxx"
	md5info   string
	base64src string
	// Optional parameters
	paramsMd5       = flag.Bool("m", false, "\n[optional] Print md5 main.go")
	paramsBase64Src = flag.Bool("s", false, "\n[optional] Print main.go")
	paramsPanic     = flag.Bool("p", false, "\n[optional] Spawn panic")
	paramsHelp      = flag.Bool("h", false, "\nhelp")
)

// Build
// ------
// $ APP="/tmp/app"; MAIN="main.go"; GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X main.md5info=$(md5 ${MAIN} |sed 's/ //g') -X main.base64src=$(cat ${MAIN} |base64)" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}
	if *paramsMd5 {
		fmt.Println(md5info)
	}
	if *paramsBase64Src {
		fmt.Println(base64src)
	}
	if *paramsPanic {
		log.Panic("Panic!")
	}
}
