package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"strconv"
)

type options struct {
	Port int `short:"p" long:"port" description:"Listen port" default:"9090"`
}

//-------------------
// << Installation >>
// $ go get -u github.com/gobuffalo/packr/packr
//
// << Build >>
// $ packr; go build -o /tmp/webapp .
// $ packr; GOOS=linux GOARCH=amd64 go build -o /tmp/webapp .
// $ /tmp/./webapp
//-------------------
// - [Using HTML Templates from a Folder (Complied) · GolangCode](https://golangcode.com/using-html-templates-from-a-folder-complied/)
//var templates = template.Must(template.ParseGlob("resources/*"))
var box = packr.NewBox("./resources")

// - [Choosing A Library to Embed Static Assets in Go](https://tech.townsourced.com/post/embedding-static-files-in-go/)
// - [gobuffalo/packr： The simple and easy way to embed static files into Go binaries.](https://github.com/gobuffalo/packr)
// - [Listening multiple ports on golang http servers](https://gist.github.com/filewalkwithme/0199060b2cb5bbc478c5)
func main() {

	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}

	// - [web applications - How do I serve CSS and JS in Go Lang - Stack Overflow](https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang)
	http.Handle("/", http.FileServer(box))
	port := ":" + strconv.Itoa(opts.Port)
	fmt.Printf("server %s\n", port)
	err := http.ListenAndServe(port, nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
