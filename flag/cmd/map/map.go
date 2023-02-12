package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type paramsMap map[string]string

func NewParamsMap() paramsMap {
	var m paramsMap
	m = make(map[string]string)
	return m
}

func (p *paramsMap) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *paramsMap) Set(v string) error {
	keyValue := strings.Split(v, ":")
	m := map[string]string(*p)
	m[keyValue[0]] = keyValue[1]
	return nil
}

var (
	paramsValueMap = NewParamsMap()
	paramsHelp     = flag.Bool("h", false /*   */, "\ndebug")
)

func main() {

	flag.Var(&paramsValueMap, "m", "[required] key values ( can be Specified multiple times. format = key:value )")
	flag.Parse()
	if *paramsHelp || len(paramsValueMap) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(paramsValueMap)
}
