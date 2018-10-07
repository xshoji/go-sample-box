package sameKeyTotallyDifferentStructure

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


type HobbyMovie struct {
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

	var hobbyGame HobbyGame
	if err := json.Unmarshal(a.Hobby, &hobbyGame); err == nil && len(hobbyGame.PlatForm) > 0 {
		u.Hobby = &hobbyGame
		return nil
	}

	var hobbyMovie HobbyMovie
	if err := json.Unmarshal(a.Hobby, &hobbyMovie); err == nil && len(hobbyMovie.DistributionCompany) > 0{
		u.Hobby = &hobbyMovie
		return nil
	}

	return nil
}

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func Run() {
	fmt.Println("[ sameKeyTotallyDifferentStructure ]")
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
	fmt.Println("")
}
