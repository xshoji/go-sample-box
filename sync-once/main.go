package main

import (
	"fmt"
	"sync"
)

var once1 sync.Once
var once2 sync.Once

func main() {
	callableOnceOnly1("first")
	callableOnceOnly1("second")
	callableOnceOnly1("third")

	callableOnceOnly2("111")
	callableOnceOnly2("222")
	callableOnceOnly2("333")
}

func callableOnceOnly1(value string) {
	once1.Do(func() {
		fmt.Println(value)
	})
}
func callableOnceOnly2(value string) {
	once2.Do(func() {
		fmt.Println(value)
	})
}
