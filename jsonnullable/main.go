package main

import (
	"encoding/json"
	"github.com/xshoji/go-sample-box/jsongetvalue/jsonnullable"
	"log"
)

type User struct {
	Name    jsonnullable.NullString  `json:"name"`
	Length  jsonnullable.NullFloat64 `json:"length"`
	IsChild jsonnullable.NullBool    `json:"isChild"`
}

func main() {
	jsonString := `
[
  {
    "name": "taro",
    "length": 180.5,
    "isChild": true
  },
  {
    "name": null,
    "length": null,
    "isChild": null
  }
]
`
	// json string to object
	var users []User
	err := json.Unmarshal([]byte(jsonString), &users)
	if err != nil {
		println(err.Error())
	}
	log.Printf("%v", users)

	// object to json string
	bytes, err := json.Marshal(users)
	if err != nil {
		println(err.Error())
	}
	log.Printf("%v", string(bytes))

}
