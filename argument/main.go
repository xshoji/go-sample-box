package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	title := flag.String("title", "", "title")
	count := flag.Int("count", 1, "count")
	isDebug := flag.Bool("debug", false, "debug")
	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *title == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("count: ", *count)
	fmt.Println("title: ", *title)
	fmt.Println("debug: ", *isDebug)

}
