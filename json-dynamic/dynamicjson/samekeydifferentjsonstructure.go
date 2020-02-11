package dynamicjson

import (
	"encoding/json"
	"fmt"
)

type UserB struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Sports SportsB `json:"sports,omitempty"`
}

type SportsB interface {
	// 何でも良い
	getType()
}

type SportsBaseballB struct {
	Position       string `json:"position"`
	InningsPitched int    `json:"inningsPitched"`
	Strikeouts     int    `json:"strikeouts"`
}

func (*SportsBaseballB) getType() {}

type SportsSimpleB string

func (*SportsSimpleB) getType() {}

// > interface要素を持つstructへのJSON Unmarshal - すぎゃーんメモ
// > https://memo.sugyan.com/entry/2018/06/23/232559
func (u *UserB) UnmarshalJSON(data []byte) error {
	type Alias UserB
	a := struct {
		Sports json.RawMessage `json:"sports"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var sportsBaseball SportsBaseballB
	if err := json.Unmarshal(a.Sports, &sportsBaseball); err == nil && len(sportsBaseball.Position) > 0 {
		u.Sports = &sportsBaseball
		return nil
	}

	var sportsSimple SportsSimpleB
	if err := json.Unmarshal(a.Sports, &sportsSimple); err == nil && len(sportsSimple) > 0 {
		u.Sports = &sportsSimple
		return nil
	}

	return nil
}

func RunSameKeyDifferentJsonStructure() {
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
	var user UserB
	json.Unmarshal([]byte(json1), &user)
	bytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(bytes))

	json2 := `
	{
	  "name":"hanako",
	  "gender":"female",
	  "age":20,
	  "sports": "Karate"
	}
	`
	var user2 UserB
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "  ")
	fmt.Println(string(bytes))
	if _, ok := user2.Sports.(*SportsBaseballB); ok {
		fmt.Println("user has SportsBaseball")
	}
	if _, ok := user2.Sports.(*SportsSimpleB); ok {
		fmt.Println("user has SportsSimple")
		// "sports"'s pointer
		p := user2.Sports.(*SportsSimpleB)
		// value of "sports"'s pointer
		fmt.Println(*p)
	}
	fmt.Println("")
}
