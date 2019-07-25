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
	if *helpFlag || *urlFlag == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("URL : " + *urlFlag)

	// create document object.
	res, _ := http.Get(*urlFlag)
	node, _ := html.Parse(res.Body)
	doc := goquery.NewDocumentFromNode(node)
	renderNodeList(doc.Nodes)

	// print title tag node and inner text.
	fmt.Println("<< basic >>")
	titleSelection := doc.Find("title").First()
	fmt.Println("Print html.Node")
	renderNode(titleSelection.Nodes[0])
	fmt.Println()
	// get by goquery.Selection
	fmt.Println("[goquery.Selection]      : " + titleSelection.Text())
	// get by html.Node
	fmt.Println("[html.Node (tag)]        : " + titleSelection.Nodes[0].Data)
	fmt.Println("[html.Node (inner text)] : " + titleSelection.Nodes[0].FirstChild.Data)
	fmt.Println()
	fmt.Println()

	// find image tag nodes and iterate printing an attribute value each.
	fmt.Println("<< iteration >>")
	fmt.Println("Print image tag src list")
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		// get by goquery
		imageUrl, _ := s.Attr("src")
		fmt.Println("[goquery.Selection.Attr] : " + imageUrl)

		// get by html.node
		for _, attr := range s.Nodes[0].Attr {
			fmt.Println("  [html.Node.Attr] " + attr.Key + " : " + attr.Val)
		}
	})
	fmt.Println()
	fmt.Println()

	// find image tag nodes and iterate printing an attribute value each.
	fmt.Println("<< find tags >>")
	fmt.Println("Print input list")
	doc.Find(`tr > td > form > input`).Each(func(_ int, s *goquery.Selection) {
		renderNodeList(s.Nodes)
	})
	fmt.Println()
	fmt.Println()

}

func renderNodeList(nodes []*html.Node) {
	for _, node := range nodes {
		renderNode(node)
	}
}

func renderNode(n *html.Node) {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	fmt.Println(buf.String())
}
