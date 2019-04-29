package util

import "fmt"

type fruits struct{}

var Fruits = fruits{}

func (f *fruits) EchoName() {
	fmt.Println("Fruits!")
}
