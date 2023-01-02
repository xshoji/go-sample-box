package util

import "fmt"

type person struct{}

var Person = person{}

func (p *person) EchoName() {
	fmt.Println("Person!")
}
