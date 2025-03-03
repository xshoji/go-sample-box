package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	params = struct {
		// Required parameters
		add      *int
		itemName *string
		// Optional parameters
		filesize *int
		binary   *string
		debug    *bool
	}{
		flag.Int("a", 0 /*           */, "[Required] add"),
		flag.String("i", "" /*       */, "[Required] item-name"),
		flag.Int("f", 10 /*          */, "[Optional] filesize"),
		flag.String("b", "000101" /* */, "[Optional] binary"),
		flag.Bool("d", false /*      */, "\ndebug"),
	}
)

// << Execution sample >>
//
// $ go run cmd/struct/main.go -a 100 -i "itemName"
// [ Command options ]
// -a 100          [Required] add
// -b 000101       [Optional] binary
// -d false        debug
// -f 10           [Optional] filesize
// -h false        help
// -i itemName     [Required] item-name
//
// $ go run cmd/struct/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build46637023/b001/exe/struct:
//	-a int
//	  	[Required] add
//	-b string
//	  	[Optional] binary (default "000101")
//	-d
//	  	debug
//	-f int
//	  	[Optional] filesize (default 10)
//	-h
//	  	help
//	-i string
//	  	[Required] item-name

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *params.add == 0 || *params.itemName == "" {
		fmt.Printf("\n[ERROR] Missing required option\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("-%s %-10v   %s\n", a.Name, a.Value, strings.Trim(a.Usage, "\n"))
	})
}
