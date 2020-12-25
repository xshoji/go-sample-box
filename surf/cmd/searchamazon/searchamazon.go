package main

import (
	"flag"
	"fmt"
	"gopkg.in/headzoo/surf.v1"
	"os"
)

var (
	//
	// Define short parameters and description ( set default value here if you need ).
	//
	keywordStringFlag = flag.String("k", "" /*  */, "[req] Keyword")
	helpBoolFlag      = flag.Bool("h", false /*   */, "\n[opt] help")
)

const (
	AmazonUrl                        = `https://www.amazon.co.jp/`
	AmazonSearchFormId               = `#nav-search-bar-form`
	AmazonSearchFormKeywordFieldName = `field-keywords`
)

//
// Usage:
// go run cmd/searchamazon/searchamazon.go -k="macbook"
//
func main() {

	flag.Parse()
	if *helpBoolFlag || *keywordStringFlag == "" {
		flag.Usage()
		os.Exit(0)
	}

	bow := surf.NewBrowser()
	err := bow.Open(AmazonUrl)
	if err != nil {
		panic(err)
	}

	form, err := bow.Form(AmazonSearchFormId)
	if err != nil {
		panic(err)
	}

	err = form.Input(AmazonSearchFormKeywordFieldName, *keywordStringFlag)
	if err != nil {
		panic(err)
	}
	err = form.Submit()
	if err != nil {
		panic(err)
	}

	fmt.Println(">> Title:")
	fmt.Println(bow.Title())
	fmt.Println()
	fmt.Println(">> Images:")
	for i, image := range bow.Images() {
		fmt.Println(i, image.ID, image.URL)
	}
	fmt.Println()
	fmt.Println(">> Links:")
	for i, link := range bow.Links() {
		fmt.Println(i, link.URL)
	}

}
