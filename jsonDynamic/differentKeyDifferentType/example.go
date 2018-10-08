package differentKeyDifferentType

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	SportsBaseball *SportsBaseball `json:"sportsBaseball,omitempty"`
	SportsSwimming *SportsSwimming `json:"sportsSwimming,omitempty"`
}

type SportsBaseball struct {
	Name string `json:"name"`
	Experience string `json:"experience"`
	Position string `json:"position"`
	InningsPitched int `json:"inningsPitched"`
	Strikeouts int `json:"strikeouts"`
}

type SportsSwimming struct {
	Name string `json:"name"`
	Experience string `json:"experience"`
	Style string `json:"style"`
	Length string `json:"length"`
	Time float64 `json:"time"`
}


func Run() {
	fmt.Println("--[ differentKeyDifferentType ]-----------------")
	json1 := `
	{
	  "name":"taro",
	  "gender":"male",
	  "age":16,
	  "sportsBaseball": {
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
	bytes, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(bytes))

	json2 := `
	{
	  "name":"hanako",
	  "gender":"female",
	  "age":20,
	  "sportsSwimming": {
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
	bytes, _ = json.MarshalIndent(user2, "", "    ")
	fmt.Println(string(bytes))
	fmt.Println("")
}
