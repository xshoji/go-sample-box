package main

import (
	"fmt"
	"github.com/xshoji/go-sample-box/enumstringer/enums"
)

func main() {
	// Print
	fmt.Println("<< Print enum values >>")
	fmt.Println(sports.Baseball)
	fmt.Println(sports.Swimming)
	fmt.Println(sports.Soccer)
	fmt.Println(sports.Karate)
	fmt.Println("")

	// Set to struct
	object := new(struct {
		Name   string
		Sports []sports.Sports
	})
	sportsList := make([]sports.Sports, 5)
	sportsList = append(sportsList, sports.Baseball)
	sportsList = append(sportsList, sports.Swimming)
	sportsList = append(sportsList, sports.Soccer)
	sportsList = append(sportsList, sports.Karate)
	object.Name = "SportList"
	object.Sports = sportsList
	fmt.Println("<< Print object values >>")
	fmt.Println(object)
}
