//go:generate go run github.com/gobuffalo/packr/v2/packr2 -v

package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"strconv"
)

type options struct {
	Port int `short:"p" long:"port" description:"Listen port" default:"9090"`
}

// > go generate? · Issue #48 · gobuffalo/packr
// > https://github.com/gobuffalo/packr/issues/48#issuecomment-439528246
var box = packr.New("resources", "./resources")

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
	fmt.Printf("http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
