package sameKeyDifferentStructure

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
	Type string `json:"type"`
	Length string `json:"length"`
	PlatForm string `json:"platform"`
	PlayTimeAverage string `json:"playTimeAverage"`
}
func (*HobbyGame) getType() {}


type HobbyMovie struct {
	Type string `json:"type"`
	Length string `json:"length"`
	DistributionCompany string `json:"distributionCompany"`
	WatchingTimeAverage string `json:"watchingTimeAverage"`
}
func (*HobbyMovie) getType() {}


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


	var template struct { Type string `json:"type"` }
	if err := json.Unmarshal(a.Hobby, &template); err != nil || len(template.Type) == 0 {
		panic("Unkown hobby type.")
	}

	if template.Type == "Game" {
		var hobby HobbyGame
		if err := json.Unmarshal(a.Hobby, &hobby); err == nil {
			u.Hobby = &hobby
			return nil
		}
	}
	if template.Type == "Movie" {
		var hobby HobbyMovie
		if err := json.Unmarshal(a.Hobby, &hobby); err == nil {
			u.Hobby = &hobby
			return nil
		}
	}
	return nil
}

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func Run() {
	fmt.Println("[ sameKeyDifferentStructure ]")
	json1 := `
	{
	  "name":"taro",
	  "gender":"male",
	  "age":16,
	  "hobby": {
	    "type": "Game",
	    "platform": "PS4",
	    "length": "5 years",
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
	  "hobby": {
	    "type": "Movie",
	    "distributionCompany": "20th Century Fox",
	    "length": "10 years",
	    "watchingTimeAverage": "2 hours"
	  }
	}
	`
	var user2 User
	json.Unmarshal([]byte(json2), &user2)
	bytes, _ = json.MarshalIndent(user2, "", "    ")
	fmt.Println(string(bytes))
	fmt.Println("")
}
