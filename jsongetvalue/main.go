package main

import (
	"encoding/json"
	"fmt"
	"github.com/xshoji/go-sample-box/jsongetvalue/jsonutil"
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
	var jsonObject interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonObject)
	if err != nil {
		println(err.Error())
	}

	fmt.Println(jsonObject)
	fmt.Print("name: ")
	fmt.Println(jsonutil.Get(jsonObject, "name"))
	fmt.Printf("{\"age\":%v}\n", jsonutil.ToJsonString(jsonutil.Get(jsonObject, "age")))
	fmt.Printf("{\"list\":%v}\n", jsonutil.ToJsonString(jsonutil.Get(jsonObject, "list")))
	list := jsonutil.Get(jsonObject, "list").([]interface{})
	for i, value := range list {
		fmt.Printf("%v:%v\n", i, value)
	}
	fmt.Println()
}
