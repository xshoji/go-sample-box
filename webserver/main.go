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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get all query strings as map
		params := r.URL.Query()
		// Get query string as single value
		query := r.URL.Query().Get("name")
		fmt.Fprintf(w, "Hello world\n")
		fmt.Fprintf(w, "name: %s\n", query)
		fmt.Fprintf(w, "Query strings: %v\n", params)
	})
	err := http.ListenAndServe(":"+strconv.Itoa(opts.Port), nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
