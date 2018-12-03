package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
)

var box = packr.NewBox("./resources")

// > zserge/lorca： Build cross-platform modern desktop apps in Go + HTML5
// > https://github.com/zserge/lorca
func main() {
	ui, _ := lorca.New("", "", 480, 320)
	defer ui.Close()

	// - [web applications - How do I serve CSS and JS in Go Lang - Stack Overflow](https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang)
	ln, err := net.Listen("tcp", "127.0.0.1:0") //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	go 	http.Serve(ln, http.FileServer(box))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	<-ui.Done()
}