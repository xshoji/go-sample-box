package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type paramsArray []string

func (p *paramsArray) String() string {
	return strings.Join(*p, ", ")
}

func (p *paramsArray) Set(v string) error {
	*p = append(*p, v)
	return nil
}

var (
	paramsValues paramsArray
	paramsHelp   = flag.Bool("h", false /*   */, "\ndebug")
)

func main() {

	flag.Var(&paramsValues, "v", "[required] values ( can be Specified multiple times )")
	flag.Parse()
	if *paramsHelp || len(paramsValues) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(paramsValues)
}
