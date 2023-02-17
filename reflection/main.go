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

func (p *Person) getGender() int8 {
	return p.gender
}

func main() {

	s := "s"
	ts := reflect.TypeOf(s)
	fmt.Println("ts: ", ts)
	fmt.Println("ts.Name(): ", ts.Name())
	fmt.Println("ts.Kind(): ", ts.Kind())

	p := Person{
		Name:   "aaa",
		Age:    16,
		gender: 1,
	}
	tp := reflect.TypeOf(p)
	fmt.Println("tp: ", tp)
	fmt.Println("tp.Name(): ", tp.Name())
	if reflect.Struct == tp.Kind() {
		fmt.Println("tp.Kind(): ", tp.Kind())
	}
	fmt.Println("tp.NumField(): ", tp.NumField())
	for i := 0; i < tp.NumField(); i++ {
		fmt.Println(tp.Field(i), tp.Field(i).Type, tp.Field(i).IsExported())
	}
	pt := reflect.PtrTo(tp)
	fmt.Println("pt.NumMethod(): ", pt.NumMethod())
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println(pt.Method(i), pt.Method(i).IsExported())
	}
}
