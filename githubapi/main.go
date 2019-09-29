package main

import (
	"fmt"
	"github.com/xshoji/go-sample-box/githubapi/client"
	"os"
)

func main() {

	user := os.Getenv("GITHUB_USER")
	repository := os.Getenv("GITHUB_REPOSITORY")
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	apiClient := client.NewClient("https://api.github.com")
	res := apiClient.Get("/repos/" + user + "/" + repository + "/issues?access_token=" + accessToken)
	fmt.Println(res.GetBody())
}
