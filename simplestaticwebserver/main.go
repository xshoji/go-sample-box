package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

//> go - Flag command line parsing in golang - Stack Overflow
//> https://stackoverflow.com/questions/19761963/flag-command-line-parsing-in-golang

var documentRootDirectoryFlag = flag.String("documentRootDirectory", "", "[require] documentRootDirectory")
var portFlag = flag.String("port", "8080", "[optional] port")
var helpFlag = flag.Bool("help", false, "help")

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(documentRootDirectoryFlag, "d", "", "= -documentRootDirectory")
	flag.StringVar(portFlag, "p", "8080", "= -port")
}

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *helpFlag == true || *documentRootDirectoryFlag == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("documentRootDirectory: ", *documentRootDirectoryFlag)

	http.Handle("/", http.FileServer(http.Dir(*documentRootDirectoryFlag)))

	fmt.Println("publish: http://localhost:" + *portFlag)

	if err := http.ListenAndServe(":"+*portFlag, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
