package sameKeyDifferentJsonStructure

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
	Position string `json:"position"`
	InningsPitched int `json:"inningsPitched"`
	Strikeouts int `json:"strikeouts"`
}
func (*SportsBaseball) getType() {}


type SportsSimple string
func (*SportsSimple) getType() {}


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

	var sportsBaseball SportsBaseball
	if err := json.Unmarshal(a.Sports, &sportsBaseball); err == nil && len(sportsBaseball.Position) > 0 {
		u.Sports = &sportsBaseball
		return nil
	}

	var sportsSimple SportsSimple
	if err := json.Unmarshal(a.Sports, &sportsSimple); err == nil && len(sportsSimple) > 0 {
		u.Sports = &sportsSimple
		return nil
	}

	return nil
}

func Run() {
	fmt.Println("--[ sameKeyDifferentJsonStructure ]-----------------")
	json1 := `
	{
	  "name":"taro",
	  "gender":"male",
	  "age":16,
	  "sports": {
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
	  "sports": "Karate"
	}
	`
	var user2 User
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "    ")
	fmt.Println(string(bytes))
	if _, ok := user2.Sports.(*SportsBaseball); ok {
		fmt.Println("user has SportsBaseball")
	}
	if _, ok := user2.Sports.(*SportsSimple); ok {
		fmt.Println("user has SportsSimple")
		// "sports"'s pointer
		p := user2.Sports.(*SportsSimple)
		// value of "sports"'s pointer
 		fmt.Println(*p)
	}
	fmt.Println("")
}
