package sameKeyDifferentJsonStructure

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Hobby Hobby `json:"hobby,omitempty"`
}

type Hobby interface {
	// 何でも良い
	getType()
}

type HobbyGame struct {
	PlatForm string `json:"platform"`
	PlayTimeAverage string `json:"playTimeAverage"`
}
func (*HobbyGame) getType() {}


type HobbySimple string
func (*HobbySimple) getType() {}


// > interface要素を持つstructへのJSON Unmarshal - すぎゃーんメモ
// > https://memo.sugyan.com/entry/2018/06/23/232559
func (u *User) UnmarshalJSON(data []byte) error {
	type alias User
	a := struct {
		Hobby json.RawMessage `json:"hobby"`
		*alias
	}{
		alias: (*alias)(u),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var hobbyGame HobbyGame
	if err := json.Unmarshal(a.Hobby, &hobbyGame); err == nil && len(hobbyGame.PlatForm) > 0 {
		u.Hobby = &hobbyGame
		return nil
	}

	var hobbySimple HobbySimple
	if err := json.Unmarshal(a.Hobby, &hobbySimple); err == nil && len(hobbySimple) > 0 {
		u.Hobby = &hobbySimple
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

	json2 := `
	{
	  "name":"hanako",
	  "gender":"female",
	  "age":20,
	  "hobby": "Movie"
	}
	`
	var user2 User
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "    ")
	fmt.Println(string(bytes))
	if _, ok := user2.Hobby.(*HobbyGame); ok {
		fmt.Println("user has HobbyGame")
	}
	if _, ok := user2.Hobby.(*HobbySimple); ok {
		fmt.Println("user has HobbySimple")
		// "hobby"'s pointer
		p := user2.Hobby.(*HobbySimple)
		// value of "hobby"'s pointer
 		fmt.Println(*p)
	}
	fmt.Println("")
}
