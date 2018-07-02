package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/go-first-sample/structs"
	"github.com/xshoji/go-sample-box/go-first-sample/utils"
)

func main() {
	// primitive types
	valueInt := 100
	valueDouble := 100.12
	valueRune := '\n'
	valueString := "test"
	fmt.Printf("%v, %v, %v, %v\n", valueInt, valueDouble, valueRune, valueString)

	// array
	valueIntArray := [...]int{1, 2, 3}
	fmt.Printf("%v\n", valueIntArray)

	// slice
	valueStringSlice := []string{"aaa", "bbb", "ccc"}
	fmt.Printf("%v\n", valueStringSlice)
	valueStringSlice = append(valueStringSlice, "ddd")
	fmt.Printf("%v\n", valueStringSlice)

	for _, v := range valueStringSlice {
		fmt.Println(v)
	}

	// public method
	// Package prefix is required when calling method
	fmt.Printf("%v\n", utils.JoinString("aaa", "bbb"))

	fmt.Printf("%v\n", utils.JoinString(utils.GetMultiReturns()))

	// 構造体の初期化
	u := structs.User{Name: "taro", Age: 16}
	// メソッドの呼び出し
	u.Talk()
}
