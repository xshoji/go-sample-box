package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Port int `short:"p" long:"port" description:"Listen port" default:"9090"`
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}
	// Get request
	// curl http://localhost:8080?query=aaa\&test=name\&name=xshoji
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

	err := http.ListenAndServe(":"+strconv.Itoa(opts.Port), nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
