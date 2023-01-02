package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	callableOnceOnly("first")
	callableOnceOnly("second")
	callableOnceOnly("third")
}

func callableOnceOnly(value string) {
	once.Do(func() {
		fmt.Println(value)
	})
}
