package main

import (
	"fmt"

	. "github.com/xshoji/go-sample-box/basic/structs"
	. "github.com/xshoji/go-sample-box/basic/utils"
)

func main() {
	// primitive types
	valueInt := 100
	valueDouble := 100.12
	valueRune := '\n'
	valueString := "test"
	fmt.Printf("valueInt, valueDouble, valueRune, valueString\n%v, %v, %v, %v\n", valueInt, valueDouble, valueRune, valueString)
	fmt.Println("")

	// array
	valueIntArray := [...]int{1, 2, 3}
	fmt.Printf("valueIntArray\n%v\n", valueIntArray)
	fmt.Println("")

	// slice
	valueStringSlice := []string{"aaa", "bbb", "ccc"}
	fmt.Printf("valueStringSlice\n%v\n", valueStringSlice)
	valueStringSlice = append(valueStringSlice, "ddd")
	fmt.Printf("%v\n", valueStringSlice)
	valueStringSlice = valueStringSlice[0:2]
	fmt.Printf("%v\n", valueStringSlice)
	fmt.Println("")

	// for
	fmt.Println("for")
	for _, v := range valueStringSlice {
		fmt.Println(v)
	}
	fmt.Println("")

	// public method
	// Package prefix is required when calling method
	fmt.Println("function")
	fmt.Printf("%v\n", JoinString("aaa", "bbb"))
	fmt.Printf("%v\n", JoinString(GetMultiReturns()))
	fmt.Println("")

	// 構造体の初期化
	fmt.Println("structs, method")
	taro := NewUser("taro", 16)
	jiro := NewUser("jiro", 99)
	// メソッドの呼び出し
	taro.Talk()
	jiro.Talk()
	fmt.Println("")
}
