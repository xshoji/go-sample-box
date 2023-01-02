package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	// Define boot arguments.
	argsPort1 = flag.Int("p1", 9090 /*   */, "[optional] Listen port 1")
	argsPort2 = flag.Int("p2", 9091 /*   */, "[optional] Listen port 2")
	argsHelp  = flag.Bool("h", false /* */, "\nhelp")
	// Logger 時刻と時刻のマイクロ秒、ディレクトリパスを含めたファイル名を出力
	logger = log.New(os.Stdout, "[Logger] ", log.Llongfile|log.LstdFlags)
)

// Response Response
type Response struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	// - [データをJSONに変換するときに任意のフォーマットを設定する - Qiita](https://qiita.com/taizo/items/2c3a338f1aeea86ce9e2)
	Time time.Time `json:"time"`
}

// NewResponse NewResponse
func NewResponse(name string, message string, time time.Time) Response {
	response := new(Response)
	response.Name = name
	response.Message = message
	response.Time = time
	return *response
}

// - [Listening multiple ports on golang http servers](https://gist.github.com/filewalkwithme/0199060b2cb5bbc478c5)
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

	finish := make(chan bool)

	server1 := http.NewServeMux()
	// - [web applications - How do I serve CSS and JS in Go Lang - Stack Overflow](https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang)
	server1.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		time.Now()
		responses := []Response{
			NewResponse("111", "This is 111 message", time.Now()),
			NewResponse("222", "This is 222 message", time.Now()),
		}
		json.NewEncoder(w).Encode(responses)
	})

	server2 := http.NewServeMux()
	server2.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		time.Now()
		responses := []Response{
			NewResponse("aaa", "This is aaa message", time.Now()),
			NewResponse("bbb", "This is bbb message", time.Now()),
		}
		json.NewEncoder(w).Encode(responses)
	})

	go func() {
		port := ":" + strconv.Itoa(*argsPort1)
		fmt.Printf("server1 %s\n", port)
		http.ListenAndServe(port, server1) //監視するポートを設定します。
	}()

	go func() {
		port := ":" + strconv.Itoa(*argsPort2)
		fmt.Printf("server2 %s\n", port)
		http.ListenAndServe(port, server2)
	}()

	// - [Go の並行処理 - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20130414/1365960707)
	<-finish
}
