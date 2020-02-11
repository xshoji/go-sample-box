package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type options struct {
	Port    int  `short:"p" long:"port" description:"Listen port" default:"8080"`
	PortTls int  `short:"s" long:"porttls" description:"Listen port on tls" default:"443"`
	Tsl     bool `short:"t" long:"tls" description:"Using TLS"`
}

func main() {

	//-------------------------
	// 引数のパース
	opts := *new(options)
	parser := flags.NewParser(&opts, flags.Default|flags.IgnoreUnknown)
	// set name
	parser.Name = "webserver"
	parser.LongDescription = "webserver"
	if _, err := parser.Parse(); err != nil {
		flagsError, _ := err.(*flags.Error)
		if flagsError.Type != flags.ErrHelp {
			// error時は明示的にHelpを表示してあげる
			fmt.Println()
			parser.WriteHelp(os.Stdout)
			fmt.Println()
		}
		fmt.Println()
		return
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
	var port string
	if opts.Tsl {
		port = ":" + strconv.Itoa(opts.PortTls)
		fmt.Printf("server(https) %s\n", port)
		err = http.ListenAndServeTLS(port, certFile, keyFile, nil)
	} else {
		port = ":" + strconv.Itoa(opts.Port)
		fmt.Printf("server(http) %s\n", port)
		err = http.ListenAndServe(":"+strconv.Itoa(opts.Port), nil)
	}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
