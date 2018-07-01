package main

import (
	"fmt"

	"./utils"
)

func main() {
	valueInt := 100
	valueDouble := 100.12
	valueRune := '\n'
	valueString := "test"
	fmt.Printf("%v, %v, %v, %v\n", valueInt, valueDouble, valueRune, valueString)

	valueIntArray := [...]int{1, 2, 3}
	fmt.Printf("%v\n", valueIntArray)

	// メソッドを呼ぶ場合はパッケージ名を先頭につけて呼び出す
	fmt.Printf("%v\n", utils.JoinString("aaa", "bbb"))

	// 複数の戻り値はそのまま引数として使うことも可能
	fmt.Printf("%v\n", utils.JoinString(utils.GetMultiReturns()))
}
