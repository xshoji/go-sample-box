package main

import (
	"flag"
	"os"
)

var (
	paramsHelp = flag.Bool("h", false /*   */, "\ndebug")
)

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}

}
