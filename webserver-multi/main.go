package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Port1 int `short:"p" long:"port1" description:"Listen port 1" default:"9090"`
	Port2 int `short:"" long:"port2" description:"Listen port 2" default:"9091"`
}

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

	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
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
		http.ListenAndServe(":"+strconv.Itoa(opts.Port1), server1) //監視するポートを設定します。
	}()

	go func() {
		http.ListenAndServe(":"+strconv.Itoa(opts.Port2), server2)
	}()

	// - [Go の並行処理 - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20130414/1365960707)
	<-finish
}
