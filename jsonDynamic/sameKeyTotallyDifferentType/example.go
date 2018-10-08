package sameKeyTotallyDifferentType

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


type SportsSwimming struct {
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

	var sportsBaseball SportsBaseball
	if err := json.Unmarshal(a.Sports, &sportsBaseball); err == nil && len(sportsBaseball.Position) > 0 {
		u.Sports = &sportsBaseball
		return nil
	}

	var sportsSwimming SportsSwimming
	if err := json.Unmarshal(a.Sports, &sportsSwimming); err == nil && len(sportsSwimming.Style) > 0{
		u.Sports = &sportsSwimming
		return nil
	}

	return nil
}

func Run() {
	fmt.Println("--[ sameKeyTotallyDifferentType ]-----------------")
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
	// > go - Golang Cast interface to struct - Stack Overflow
	// > https://stackoverflow.com/questions/50939497/golang-cast-interface-to-struct
	if _, ok := user.Sports.(*SportsSwimming); ok {
		fmt.Println("user has SportsSwimming")
	}
	if _, ok := user.Sports.(*SportsBaseball); ok {
		fmt.Println("user has SportsBaseball")
		fmt.Println(user.Sports.(*SportsBaseball).Position)
		fmt.Println(user.Sports.(*SportsBaseball).InningsPitched)
		fmt.Println(user.Sports.(*SportsBaseball).Strikeouts)
	}

	json2 := `
	{
	  "name":"hanako",
	  "gender":"female",
	  "age":20,
	  "sports": {
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
