package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xshoji/go-sample-box/githubapi/client"
	"log"
	"os"
)

var helpFlag = flag.Bool("help", false, "help")
var commandFlag = flag.String("command", "", "[required] command")
var user string
var repository string
var accessToken string
var apiClient *client.Client

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(commandFlag, "c", "", "= -command")
	user = os.Getenv("GITHUB_USER")
	repository = os.Getenv("GITHUB_REPOSITORY")
	accessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	apiClient = client.NewClient("https://api.github.com")
}

func main() {

	flag.Parse()
	// Required parameter
	// - [Can Go's `flag` package print usage? - Stack Overflow](https://stackoverflow.com/questions/23725924/can-gos-flag-package-print-usage)
	if *helpFlag || *commandFlag == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println("command: ", *commandFlag)

	if *commandFlag == "issues" {
		issues()
	} else {
		log.Fatal("Unknown command.")
	}
}

func issues() {
	res := apiClient.Get("/repos/" + user + "/" + repository + "/issues?access_token=" + accessToken)
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(res.GetBody()), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
