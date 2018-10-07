package undefinedType

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Hobby interface{} `json:"hobby,omitempty"`
}

func Run() {
	fmt.Println("--[ undefinedType ]-----------------")
	json1 := `
	{
	  "name":"taro",
	  "gender":"male",
	  "age":16,
	  "hobby": {
	    "platform": "PS4",
	    "playTimeAverage": "2 hours"
	  }
	}
	`
	var user User
	json.Unmarshal([]byte(json1), &user)
	bytes, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(bytes))
	if platform, ok := user.Hobby.(map[string]interface{})["platform"]; ok {
		fmt.Println(platform.(string))
	}
	if playTimeAverage, ok := user.Hobby.(map[string]interface{})["playTimeAverage"]; ok {
		fmt.Println(playTimeAverage.(string))
	}

	json2 := `
	{
	  "name":"hanako",
	  "gender":"female",
	  "age":20,
	  "hobby": {
	    "distributionCompany": "20th Century Fox",
	    "watchingTimeAverage": "2 hours"
	  }
	}
	`
	var user2 User
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "    ")
	fmt.Println(string(bytes))
	if distributionCompany, ok := user2.Hobby.(map[string]interface{})["distributionCompany"]; ok {
		fmt.Println(distributionCompany.(string))
	}
	if watchingTimeAverage, ok := user2.Hobby.(map[string]interface{})["watchingTimeAverage"]; ok {
		fmt.Println(watchingTimeAverage.(string))
	}
	fmt.Println("")
}
