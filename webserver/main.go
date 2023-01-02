package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	// Define boot arguments.
	argsPort         = flag.Int("p", 8080 /*   */, "[optional] Listen port")
	argsHttpsEnabled = flag.Bool("s", false /* */, "\n[optional] Enable TLS (flag option)")
	argsHelp         = flag.Bool("h", false /* */, "\nhelp")
	// Logger 時刻と時刻のマイクロ秒、ディレクトリパスを含めたファイル名を出力
	logger = log.New(os.Stdout, "[Logger] ", log.Llongfile|log.LstdFlags)
)

func main() {

	//-------------------------
	// 引数のパース
	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *argsHelp {
		flag.Usage()
		os.Exit(0)
	}

	certFile, _ := filepath.Abs("server.crt")
	keyFile, _ := filepath.Abs("server.key")

	// Get request
	// curl http://localhost:8080/get?query=aaa\&test=name\&name=xshoji
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// - [How can I handle http requests of different methods to / in Go? - Stack Overflow](https://stackoverflow.com/questions/15240884/how-can-i-handle-http-requests-of-different-methods-to-in-go)
		if r.Method != "GET" {
			fmt.Fprintf(w, "Incorrect method. GET only.\n")
			return
		}
		// Get all query strings as map
		params := r.URL.Query()
		// Get query string as single value
		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, "Hello world\n")
		fmt.Fprintf(w, "name: %s\n", name)
		fmt.Fprintf(w, "Query strings: %v\n", params)
	})

	// Post request
	// curl http://localhost:8080/post -d "name=xshoji" -d "id=1001" -X POST
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			fmt.Fprintf(w, "Incorrect method. POST only.\n")
			return
		}
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.PostForm.Get("name")
		fmt.Fprintf(w, "name: %v\n", name)
		fmt.Fprintf(w, "Post data: %v\n", r.PostForm)
	})

	var err error
	port := ":" + strconv.Itoa(*argsPort)
	if *argsHttpsEnabled {
		logger.Printf("server(https) %s\n", port)
		err = http.ListenAndServeTLS(port, certFile, keyFile, nil)
	} else {
		logger.Printf("server(http) %s\n", port)
		err = http.ListenAndServe(port, nil)
	}
	if err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}
