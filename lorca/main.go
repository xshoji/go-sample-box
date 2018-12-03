package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func create() {
	var seed = strconv.FormatInt(time.Now().UnixNano(), 10)
	var shaBytes = sha256.Sum256([]byte(seed))
	var randomString = hex.EncodeToString(shaBytes[:])
	var filePath = "/tmp/" + randomString
	fmt.Println("Create!: " + filePath)
	f, err := os.Create(filePath)
	check(err)
	defer f.Close()
}

// > zserge/lorca： Build cross-platform modern desktop apps in Go + HTML5
// > https://github.com/zserge/lorca
// go run gen.go
// go build -o /tmp/app main.go assets.go
func main() {
	ui, _ := lorca.New("", "", 480, 320)
	defer ui.Close()

	// - [web applications - How do I serve CSS and JS in Go Lang - Stack Overflow](https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang)
	ln, err := net.Listen("tcp", "127.0.0.1:0") //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	go 	http.Serve(ln, http.FileServer(FS))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	ui.Bind("create", create)
	<-ui.Done()
}