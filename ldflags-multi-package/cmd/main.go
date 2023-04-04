package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xshoji/go-sample-box/ldflags-multi-package/sub"
	"os"
)

var (
	// Set value by -ldflags "-X xxx"
	base64src string
	// Optional parameters
	paramsPrintSrc = flag.Bool("print-src", false, "[optional] Print main.go")
	paramsPanic    = flag.Bool("p", false /* */, "\n[optional] Spawn panic")
	paramsHelp     = flag.Bool("h", false /* */, "\nhelp")
)

// Build
// ------
// $ APP="/tmp/app"; MAIN="cmd/main.go"; GOOS=darwin GOARCH=amd64; toKV() { echo "\"${1}\":\"$(cat ${1} |base64)\""; }; export -f toKV; toBase64() { find . -type f -not -path '*/.*' |xargs -I{} bash -c "toKV {}" |awk 'NR==1{printf "{"}{printf "%s%s",sep,$0; sep=","}END{print "}"}' |base64; }; export -f toBase64; go build -ldflags="-s -w -X 'main.base64src=$(toBase64)'" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}
	if *paramsPrintSrc {
		bytesJson, _ := base64.StdEncoding.DecodeString(base64src)
		fmt.Printf("%s\n", bytesJson)
		var fileContentMap map[string]string
		_ = json.Unmarshal(bytesJson, &fileContentMap)
		for filePath, contentBase64 := range fileContentMap {
			fmt.Println("--------")
			fmt.Println("filePath:" + filePath)
			bytesFileContent, _ := base64.StdEncoding.DecodeString(contentBase64)
			fmt.Println("contents:\n" + string(bytesFileContent))
		}
		os.Exit(0)
	}

	fmt.Println(sub.GetValue())
}
