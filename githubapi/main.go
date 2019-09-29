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
var paramsFlag = flag.String("params", "", "[optional] params ( json format )")
var user string
var repository string
var accessToken string
var apiClient *client.Client

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(commandFlag, "c", "", "= -command")
	flag.StringVar(paramsFlag, "p", "", "= -params")
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
	} else if *commandFlag == "issueComments" {
		issueComments()
	} else {
		log.Fatal("Unknown command.")
	}
}

func issues() {

	var paramMap map[string]interface{}
	json.Unmarshal([]byte(*paramsFlag), &paramMap)
	value, ok := paramMap["issueNumber"]
	if !ok {
		log.Fatal("-p {\"issueNumber\":\"1234\"} is required. (\"issueNumber\":null is allowed.)")
	}
	var issueNumber string
	if value != nil {
		issueNumber = "/" + value.(string)
	}

	res := apiClient.Get("/repos/" + user + "/" + repository + "/issues" + issueNumber + "?access_token=" + accessToken)
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(res.GetBody()), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func issueComments() {

	var paramMap map[string]interface{}
	json.Unmarshal([]byte(*paramsFlag), &paramMap)
	issueNumber, ok := paramMap["issueNumber"]
	if !ok {
		log.Fatal("-p {\"issueNumber\":\"1234\"} is required.")
	}

	res := apiClient.Get("/repos/" + user + "/" + repository + "/issues/" + issueNumber.(string) + "/comments?access_token=" + accessToken)
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(res.GetBody()), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
