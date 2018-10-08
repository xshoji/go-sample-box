package differentKeyDifferentType

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	HobbyGame *HobbyGame `json:"hobbyGame,omitempty"`
	HobbyMovie *HobbyMovie `json:"hobbyMovie,omitempty"`
}

type HobbyGame struct {
	Type string `json:"type"`
	PlatForm string `json:"platform"`
	PlayTimeAverage string `json:"playTimeAverage"`
}

type HobbyMovie struct {
	Type string `json:"type"`
	DistributionCompany string `json:"distributionCompany"`
	WatchingTimeAverage string `json:"watchingTimeAverage"`
}

func Run() {
	fmt.Println("--[ differentKeyDifferentType ]-----------------")
	json1 := `
	{
	  "name":"taro",
	  "gender":"male",
	  "age":16,
	  "hobbyGame": {
	    "type": "Game",
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
	  "hobbyMovie": {
	    "type": "Movie",
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
