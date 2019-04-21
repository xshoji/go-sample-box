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
	var valueInt8 int8
	valueInt8 = 1
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
	// > Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記
	// > http://matope.hatenablog.com/entry/2014/04/22/101127
	fmt.Println("<< cast >>")
	fmt.Printf("valueInt8: %d, valueString: %s\n", valueInt, strconv.Itoa(int(valueInt8)))
	fmt.Printf("valueInt: %d, valueString: %s\n", valueInt, strconv.Itoa(valueInt))
	fmt.Printf("valueInt64: %d, valueString: %s\n", valueInt64, strconv.FormatInt(valueInt64, 10))
	fmt.Printf("valueFloat: %v, valueString: %s\n", valueFloat64, strconv.FormatFloat(valueFloat64, 'f', 4, 32))
	fmt.Printf("valueFloat: %v, valueInt: %v\n", valueFloat64, int(valueFloat64))
	fmt.Printf("valueInt: %v, valueFloat: %v\n", valueInt, float64(valueInt))
	fmt.Printf("valueString: %v, value[]byte: %v\n", valueString, []byte(valueString))
	valueInt, _ = strconv.Atoi(valueString)
	fmt.Printf("valueString: %v, valueInt: %v\n", valueString, valueInt)
	valueInt64, _ = strconv.ParseInt(valueString, 10, 64)
	fmt.Printf("valueString: %v, valueInt64: %v\n", valueString, valueInt64)
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
	valueStringSlice = make([]string, 5)
	fmt.Println(valueStringSlice)
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
	createRandomNumber := func() int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(1000000-1) + 1
	}
	fmt.Println(createRandomNumber())
	fmt.Println("")

	// Create random string
	fmt.Println("<< random string >>")
	createRandomString := func() string {
		seed := strconv.FormatInt(time.Now().UnixNano(), 10)
		shaBytes := sha256.Sum256([]byte(seed))
		return hex.EncodeToString(shaBytes[:])
	}
	fmt.Println(createRandomString())
	fmt.Println("")

	fmt.Println("<< size >>")
	sizeString := "aaa"
	sizeArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sizeSlice := make([]string, 15)
	fmt.Printf("Size of string : %v\n", len(sizeString))
	fmt.Printf("Size of array : %v\n", len(sizeArray))
	fmt.Printf("Size of slice : %v\n", len(sizeSlice))

	// Nullable variables
	// > go - How to make a nullable field in Golang struct? - Stack Overflow
	// > https://stackoverflow.com/questions/51998165/how-to-make-a-nullable-field-in-golang-struct
	fmt.Println("<< Nullable values handling >>")
	fmt.Println("")
	fmt.Println("[ pointer ]")
	user := struct {
		Name *string
		Age  *int
	}{}
	user.Name = &valueString
	user.Age = &valueInt
	fmt.Println("Refer values")
	fmt.Println(*user.Name)
	fmt.Println(*user.Age)
	fmt.Println(user.Name == nil)
	fmt.Println(user.Age == nil)

	fmt.Println("")
	fmt.Println("Refer nil values")
	user.Name = nil
	user.Age = nil
	fmt.Println(user.Name == nil)
	fmt.Println(user.Age == nil)
	fmt.Println("")

	// DateTime format
	// > Golangでの日付のフォーマット指定の方法について - Qiita
	// > https://qiita.com/unbabel/items/c8782420391c108e3cac
	fmt.Println("<< now >>")
	now := time.Now()
	fmt.Println("plane:")
	fmt.Println(now)
	fmt.Println("formatted:")
	fmt.Println(now.Format("2006 / 01 [January(Jan)] / 02 [Monday(Mon)] 15:04:05 [MST]"))
	fmt.Println("")
}
