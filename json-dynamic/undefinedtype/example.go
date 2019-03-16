package undefinedtype

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string      `json:"name"`
	Gender string      `json:"gender"`
	Age    int         `json:"age"`
	Sports interface{} `json:"sports,omitempty"`
}

func Run() {
	fmt.Println("--[ undefinedType ]-----------------")
	json1 := `
	{
	  "name":"taro",
	  "gender":"male",
	  "age":16,
	  "sports": {
	    "name": "Baseball",
	    "experience":"3 years",
	    "position":"Pitcher",
	    "inningsPitched":215,
	    "strikeouts":222
	  }
	}
	`
	var user User
	json.Unmarshal([]byte(json1), &user)
	bytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(bytes))
	if sportsName, ok := user.Sports.(map[string]interface{})["name"]; ok {
		fmt.Println(sportsName.(string))
	}
	if inningsPitched, ok := user.Sports.(map[string]interface{})["inningsPitched"]; ok {
		fmt.Println(int(inningsPitched.(float64)))
	}

	json2 := `
	{
	  "name":"hanako",
	  "gender":"female",
	  "age":20,
	  "sports": {
	    "name": "Swimming",
	    "experience": "5 years",
	    "style": "Freestyle",
	    "length": "100m",
	    "time": 46.91
	  }
	}
	`
	var user2 User
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "  ")
	fmt.Println(string(bytes))
	if sportsName, ok := user2.Sports.(map[string]interface{})["name"]; ok {
		fmt.Println(sportsName.(string))
	}
	if time, ok := user2.Sports.(map[string]interface{})["time"]; ok {
		fmt.Println(time.(float64))
	}
	fmt.Println("")
}
