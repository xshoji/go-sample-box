package main

import "fmt"

func main() {
	valueInt := 100
	valueDouble := 100.12
	valueRune := '\n'
	valueString := "test"
	fmt.Printf("%v, %v, %v, %v\n", valueInt, valueDouble, valueRune, valueString)
	
	valueIntArray := [...]int{1,2,3}
	fmt.Printf("%v\n", valueIntArray)
}