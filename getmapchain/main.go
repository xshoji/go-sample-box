package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	jsonString := `
{
  "name": "taro",
  "age": 16,
  "items": [
    "book",
    "pen",
    "water"
  ],
  "friends": [
    {
      "name": "jiro",
      "age": 16
    }
  ],
  "school": {
    "name": "Univ. of California",
    "established": "100 years"
  }
}
`
	var jsonObject interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonObject)
	if err != nil {
		println(err.Error())
	}

	fmt.Println(jsonObject)
	fmt.Print("name: ")
	fmt.Println(GetValueInInterface(jsonObject, "name"))
	fmt.Print("items: ")
	fmt.Println(GetValueInInterface(jsonObject, "items"))
	fmt.Print("items.0: ")
	fmt.Println(GetValueInInterface(jsonObject, "items.0"))
	fmt.Print("items.10: ")
	fmt.Println(GetValueInInterface(jsonObject, "items.10"))
	fmt.Print("friends: ")
	fmt.Println(GetValueInInterface(jsonObject, "friends"))
	fmt.Print("friends.0: ")
	fmt.Println(GetValueInInterface(jsonObject, "friends.0"))
	fmt.Print("friends.0.name: ")
	fmt.Println(GetValueInInterface(jsonObject, "friends.0.name"))
	fmt.Print("friends.1.name: ")
	fmt.Println(GetValueInInterface(jsonObject, "friends.1.name"))
	fmt.Print("school: ")
	fmt.Println(GetValueInInterface(jsonObject, "school"))
	fmt.Print("school.established: ")
	fmt.Println(GetValueInInterface(jsonObject, "school.established"))
	fmt.Print("school.hoge: ")
	fmt.Println(GetValueInInterface(jsonObject, "school.hoge"))
}

func GetValueInInterface(object interface{}, keyChain string) interface{} {
	keys := strings.Split(keyChain, ".")
	var v interface{}
	var values []interface{}
	var i int
	var ok bool
	var result interface{}
	var exists bool
	for _, key := range keys {
		exists = false
		if _, ok = object.(map[string]interface{}); ok {
			exists = true
			object = object.(map[string]interface{})[key]
			result = object
			continue
		}
		if values, ok = object.([]interface{}); ok {
			for i, v = range values {
				if strconv.FormatInt(int64(i), 10) == key {
					exists = true
					object = v
					result = object
					continue
				}
			}
		}
	}
	if exists {
		return result
	}
	return nil
}
