package main

import (
	"encoding/json"
	"fmt"
	"github.com/xshoji/go-sample-box/getjsonvalue/jsonutil"
)

func main() {
	jsonString := `
{
  "name": "taro",
  "age": 16
}
`
	var jsonObject interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonObject)
	if err != nil {
		println(err.Error())
	}

	fmt.Println(jsonObject)
	fmt.Print("name: ")
	fmt.Println(jsonutil.Get(jsonObject, "name"))
	fmt.Printf(`{"age":%v}`, jsonutil.ToJsonString(jsonutil.Get(jsonObject, "age")))
	fmt.Println()
}
