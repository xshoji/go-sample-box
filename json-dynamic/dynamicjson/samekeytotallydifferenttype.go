package dynamicjson

import (
	"encoding/json"
	"fmt"
)

type UserD struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Sports SportsD `json:"sports,omitempty"`
}

type SportsD interface {
	// 何でも良い
	getType()
}

type SportsBaseballD struct {
	Position       string `json:"position"`
	InningsPitched int    `json:"inningsPitched"`
	Strikeouts     int    `json:"strikeouts"`
}

func (*SportsBaseballD) getType() {}

type SportsSwimmingD struct {
	Style  string  `json:"style"`
	Length string  `json:"length"`
	Time   float64 `json:"time"`
}

func (*SportsSwimmingD) getType() {}

// > interface要素を持つstructへのJSON Unmarshal - すぎゃーんメモ
// > https://memo.sugyan.com/entry/2018/06/23/232559
func (u *UserD) UnmarshalJSON(data []byte) error {
	type Alias UserD
	a := struct {
		Sports json.RawMessage `json:"sports"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var sportsBaseball SportsBaseballD
	if err := json.Unmarshal(a.Sports, &sportsBaseball); err == nil && len(sportsBaseball.Position) > 0 {
		u.Sports = &sportsBaseball
		return nil
	}

	var sportsSwimming SportsSwimmingD
	if err := json.Unmarshal(a.Sports, &sportsSwimming); err == nil && len(sportsSwimming.Style) > 0 {
		u.Sports = &sportsSwimming
		return nil
	}

	return nil
}

func RunSameKeyTotallyDifferentType() {
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
	var user UserD
	json.Unmarshal([]byte(json1), &user)
	bytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(bytes))
	// > go - Golang Cast interface to struct - Stack Overflow
	// > https://stackoverflow.com/questions/50939497/golang-cast-interface-to-struct
	if _, ok := user.Sports.(*SportsSwimmingD); ok {
		fmt.Println("user has SportsSwimming")
	}
	if _, ok := user.Sports.(*SportsBaseballD); ok {
		fmt.Println("user has SportsBaseball")
		fmt.Println(user.Sports.(*SportsBaseballD).Position)
		fmt.Println(user.Sports.(*SportsBaseballD).InningsPitched)
		fmt.Println(user.Sports.(*SportsBaseballD).Strikeouts)
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
	var user2 UserD
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "  ")
	fmt.Println(string(bytes))
	fmt.Println("")
}
