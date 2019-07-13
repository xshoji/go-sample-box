package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

var helpFlag = flag.Bool("help", false, "help")
var urlFlag = flag.String("url", "", "[require] url")

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(urlFlag, "u", "", "= -url")
}

func main() {

	flag.Parse()
	if *helpFlag == true || *urlFlag == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("URL : " + *urlFlag)

	res, _ := http.Get(*urlFlag)
	node, _ := html.Parse(res.Body)
	doc := goquery.NewDocumentFromNode(node)
	for _, node := range doc.Nodes {
		renderNode(node)
	}
}

func renderNode(n *html.Node) {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	fmt.Println(buf.String())
}
