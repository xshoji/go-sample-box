package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int8
	gender int8
}

func (p *Person) GetGender() int8 {
	return p.gender
}

func main() {
	p := Person{
		Name:   "aaa",
		Age:    16,
		gender: 1,
	}
	t := reflect.TypeOf(p)
	fmt.Println("t: ", t)
	fmt.Println("t.Name(): ", t.Name())
	fmt.Println("t.NumField(): ", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i), t.Field(i).IsExported())
	}
	pt := reflect.PtrTo(t)
	fmt.Println("pt.NumMethod(): ", pt.NumMethod())
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println(pt.Method(i), pt.Method(i).IsExported())
	}

}
