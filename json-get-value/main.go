package main

import (
	"encoding/json"
	"fmt"

	"github.com/xshoji/go-sample-box/json-get-value/jsonutil"
)

func main() {
	jsonString := `
{
  "name": "taro",
  "age": 16,
  "list": [
    10,
    20,
    30
  ]
}
`
	var jsonObject any
	err := json.Unmarshal([]byte(jsonString), &jsonObject)
	if err != nil {
		println(err.Error())
	}

	fmt.Println(jsonObject)
	fmt.Print("name: ")
	fmt.Println(jsonutil.Get(jsonObject, "name"))
	fmt.Printf("{\"age\":%v}\n", jsonutil.ToJsonString(jsonutil.Get(jsonObject, "age")))
	fmt.Printf("{\"list\":%v}\n", jsonutil.ToJsonString(jsonutil.Get(jsonObject, "list")))
	list := jsonutil.Get(jsonObject, "list").([]any)
	for i, value := range list {
		fmt.Printf("%v:%v\n", i, value)
	}
	fmt.Println()
}
