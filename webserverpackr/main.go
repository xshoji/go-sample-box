package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/jessevdk/go-flags"
)

type options struct {
	Port int `short:"p" long:"port" description:"Listen port" default:"9090"`
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

//-------------------
// $ go get -u github.com/gobuffalo/packr
// $ go get -u github.com/gobuffalo/packr/packr
// $ packr
// $ GOOS=linux GOARCH=amd64 go build -o /tmp/webapp main.go a_main-packr.go
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
	err := http.ListenAndServe(":"+strconv.Itoa(opts.Port), nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
