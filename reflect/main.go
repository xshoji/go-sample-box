package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int8
	Tags   []string
	gender int8
}

func (p *Person) GetGender() int8 {
	return p.gender
}

// Cannot access private method, even via reflect.
func (p *Person) getGender() int8 {
	return p.gender
}

func main() {

	s := "s"
	ts := reflect.TypeOf(s)
	fmt.Println("ts:", ts)
	fmt.Println("  ts.Name():", ts.Name())
	fmt.Println("  ts.Kind():", ts.Kind())

	p := Person{
		Name:   "aaa",
		Age:    16,
		gender: 1,
	}
	tp := reflect.TypeOf(p)
	fmt.Println("tp:", tp)
	fmt.Println("  tp.Name():", tp.Name())
	if reflect.Struct == tp.Kind() {
		fmt.Println("  tp.Kind():", tp.Kind())
	}
	fmt.Println("  tp.NumField():", tp.NumField())
	for i := 0; i < tp.NumField(); i++ {
		fmt.Println("    tp.Field(i), tp.Field(i).Type, tp.Field(i).IsExported():", tp.Field(i), tp.Field(i).Type, tp.Field(i).IsExported())
	}
	vp := reflect.ValueOf(p)
	fmt.Println("vp: ", vp)
	fmt.Println("  vp.NumField(): ", vp.NumField())
	for i := 0; i < vp.NumField(); i++ {
		fmt.Println("    vp.Field(i), vp.Field(i).Type():", vp.Field(i), vp.Field(i).Type())
		fmt.Println("      vp.Field(i).IsValid():     ", vp.Field(i).IsValid())
		fmt.Println("      vp.Field(i).IsZero():      ", vp.Field(i).IsZero())
		fmt.Println("      vp.Field(i).CanAddr():     ", vp.Field(i).CanAddr())
		fmt.Println("      vp.Field(i).CanSet():      ", vp.Field(i).CanSet())
		fmt.Println("      vp.Field(i).CanComplex():  ", vp.Field(i).CanComplex())
		fmt.Println("      vp.Field(i).CanInt():      ", vp.Field(i).CanInt())
		fmt.Println("      vp.Field(i).CanInterface():", vp.Field(i).CanInterface())
		if vp.Field(i).CanInterface() {
			fmt.Println("      vp.Field(i).Interface():", vp.Field(i).Interface())
		}
	}
	pt := reflect.PtrTo(tp)
	fmt.Println("pt:", pt)
	fmt.Println("  pt.NumMethod():", pt.NumMethod())
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("    pt.Method(i), pt.Method(i).IsExported():", pt.Method(i), pt.Method(i).IsExported())
	}

}
