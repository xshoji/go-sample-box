package sameKeyDifferentType

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Sports Sports `json:"sports,omitempty"`
}

type Sports interface {
	// 何でも良い
	getType()
}

type SportsBaseball struct {
	Name string `json:"name"`
	Experience string `json:"experience"`
	Position string `json:"position"`
	InningsPitched int `json:"inningsPitched"`
	Strikeouts int `json:"strikeouts"`
}
func (*SportsBaseball) getType() {}


type SportsSwimming struct {
	Name string `json:"name"`
	Experience string `json:"experience"`
	Style string `json:"style"`
	Length string `json:"length"`
	Time float64 `json:"time"`
}
func (*SportsSwimming) getType() {}


// > interface要素を持つstructへのJSON Unmarshal - すぎゃーんメモ
// > https://memo.sugyan.com/entry/2018/06/23/232559
func (u *User) UnmarshalJSON(data []byte) error {
	type alias User
	a := struct {
		Sports json.RawMessage `json:"sports"`
		*alias
	}{
		alias: (*alias)(u),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var template struct { Name string `json:"name"` }
	if err := json.Unmarshal(a.Sports, &template); err != nil || len(template.Name) == 0 {
		panic("Unkown sports.")
	}

	if template.Name == "Baseball" {
		var sportsBaseball SportsBaseball
		if err := json.Unmarshal(a.Sports, &sportsBaseball); err == nil {
			u.Sports = &sportsBaseball
			return nil
		}
	}
	if template.Name == "Swimming" {
		var sportsSwimming SportsSwimming
		if err := json.Unmarshal(a.Sports, &sportsSwimming); err == nil {
			u.Sports = &sportsSwimming
			return nil
		}
	}
	return nil
}

func Run() {
	fmt.Println("--[ sameKeyDifferentType ]-----------------")
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
	bytes, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(bytes))

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
	bytes, _ = json.MarshalIndent(user2, "", "    ")
	fmt.Println(string(bytes))
	fmt.Println("")
}
