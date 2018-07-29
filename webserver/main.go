package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	help := flag.Bool("h", false, "help")
	port := flag.Int("port", 9090, "[optional] port")
	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *help == true {
		flag.Usage()
		os.Exit(0)
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
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
