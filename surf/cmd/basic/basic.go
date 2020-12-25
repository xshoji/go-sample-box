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
	urlStringFlag = flag.String("u", "" /*    */, "[req] URL")
	helpBoolFlag  = flag.Bool("h", false /*   */, "\n[opt] help")
)

//
// Usage:
// go run cmd/basic/basic.go -u="https://www.google.com"
//
func main() {

	flag.Parse()
	if *helpBoolFlag || *urlStringFlag == "" {
		flag.Usage()
		os.Exit(0)
	}

	bow := surf.NewBrowser()
	err := bow.Open(*urlStringFlag)
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
