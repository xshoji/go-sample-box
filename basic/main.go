package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xshoji/go-sample-box/basic/structs"
	"github.com/xshoji/go-sample-box/basic/utils"
)

func main() {
	// primitive types
	valueInt := 100
	valueInt64 := time.Now().UnixNano()
	valueFloat64 := 100.12
	valueRune := '\n'
	valueString := "test"
	valueBytes := []byte("test")
	fmt.Println("<< primitive >>")
	fmt.Printf("valueInt, valueInt64, valueFloat, valueRune, valueString\n%v, %v, %v, %v, %v\n", valueInt, valueInt64, valueFloat64, valueRune, valueString)
	fmt.Println("")

	// 文字シーケンス
	valueString2 := `{
  "aaa", "bbb",
  "ccc": 111
}`
	fmt.Println("<< string sequence >>")
	fmt.Println(valueString2)
	fmt.Println("")

	// cast
	fmt.Println("<< cast >>")
	fmt.Printf("valueInt: %d, valueString: %s\n", valueInt, strconv.Itoa(valueInt))
	fmt.Printf("valueInt64: %d, valueString: %s\n", valueInt64, strconv.FormatInt(valueInt64, 10))
	fmt.Printf("valueFloat: %v, valueString: %s\n", valueFloat64, strconv.FormatFloat(valueFloat64, 'f', 4, 32))
	fmt.Printf("valueFloat: %v, valueInt: %v\n", valueFloat64, int(valueFloat64))
	fmt.Printf("valueInt: %v, valueFloat: %v\n", valueInt, float64(valueInt))
	fmt.Printf("valueString: %v, value[]byte: %v\n", valueString, []byte(valueString))
	fmt.Printf("value[]byte: %v, valueString: %s\n", valueBytes, string(valueBytes))
	fmt.Println("")

	// pointer
	_valueInt := &valueInt
	fmt.Println("<< pointer >>")
	fmt.Printf("valueInt pointer: %v\n", _valueInt)
	fmt.Printf("valueInt pointers value: %v\n", *_valueInt)
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
	fmt.Println("<< for >>")
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
	fmt.Println("<< function >>")
	fmt.Printf("%v\n", utils.JoinString("aaa", "bbb"))
	fmt.Printf("%v\n", utils.JoinString(utils.GetMultiReturns()))
	fmt.Println("")

	// 構造体の初期化
	fmt.Println("<< structs, method >>")
	taro := structs.NewUser("taro", 16)
	jiro := structs.NewUser("jiro", 99)
	// メソッドの呼び出し
	taro.Talk()
	jiro.Talk()
	fmt.Println("")

	// 無名構造体
	anonymousStruct := struct {
		Name     string
		Age      int
		Language string
	}{
		"taro",
		100,
		"Japanese",
	}
	fmt.Println("<< anonymous struct >>")
	fmt.Print(anonymousStruct)
	fmt.Println("")


	// 無名関数
	anonymousFunction := func() {
		fmt.Println("This is anonymousFunction.")
	}
	fmt.Println("<< anonymous function >>")
	anonymousFunction()
	fmt.Println("")

	// Create random integer
	fmt.Println("<< random integer >>")
	rand.Seed(time.Now().UnixNano())
	millsec := rand.Intn(1000-1) + 1
	fmt.Println(millsec)
	fmt.Println("")

	// Create random string
	fmt.Println("<< random string >>")
	var seed = strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(seed)
	var shaBytes = sha256.Sum256([]byte(seed))
	var randomString = hex.EncodeToString(shaBytes[:])
	fmt.Println(randomString)
	fmt.Println("")

}
