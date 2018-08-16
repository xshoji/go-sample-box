package main

import (
	"fmt"
	"strconv"

	"github.com/xshoji/go-sample-box/basic/structs"
	"github.com/xshoji/go-sample-box/basic/utils"
)

func main() {
	// primitive types
	valueInt := 100
	valueDouble := 100.12
	valueRune := '\n'
	valueString := "test"
	fmt.Printf("valueInt, valueDouble, valueRune, valueString\n%v, %v, %v, %v\n", valueInt, valueDouble, valueRune, valueString)
	fmt.Println("")

	// cast
	fmt.Printf("valueInt: %d, valueString: %s\n", valueInt, strconv.Itoa(valueInt))
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

	//map
	mapValues := map[string]string{
		"aaa": "aaa",
		"bbb": "aaa",
		"ccc": "aaa",
	}

	// for
	fmt.Println("for")
	fmt.Printf("valueStringSlice size:%d\n", len(valueStringSlice))
	for _, v := range valueStringSlice {
		fmt.Println(v)
	}
	for k, v := range mapValues {
		fmt.Println(k + ":" + v)
	}
	fmt.Println("")

	// public method
	// Package prefix is required when calling method
	fmt.Println("function")
	fmt.Printf("%v\n", utils.JoinString("aaa", "bbb"))
	fmt.Printf("%v\n", utils.JoinString(utils.GetMultiReturns()))
	fmt.Println("")

	// 構造体の初期化
	fmt.Println("structs, method")
	taro := structs.NewUser("taro", 16)
	jiro := structs.NewUser("jiro", 99)
	// メソッドの呼び出し
	taro.Talk()
	jiro.Talk()
	fmt.Println("")
}
