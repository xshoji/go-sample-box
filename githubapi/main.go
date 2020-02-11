package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var helpFlag = flag.Bool("help", false, "help")
var commandFlag = flag.String("command", "", "[required] command ( issues | issueComments )")
var paramsFlag = flag.String("params", "", "[optional] params ( json format )")
var user string
var repository string
var accessToken string
var githubApiUrlBase string

func init() {
	flag.BoolVar(helpFlag, "h", false, "= -help")
	flag.StringVar(commandFlag, "c", "", "= -command")
	flag.StringVar(paramsFlag, "p", "", "= -params")
	user = os.Getenv("GITHUB_USER")
	repository = os.Getenv("GITHUB_REPOSITORY")
	accessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	githubApiUrlBase = "https://api.github.com"
}

// cp setup.sh.dist setup.sh
// source setup.sh
// go run main.go -c=issues -p='{"issueNumber":null}'
// go run main.go -c=issueComments -p='{"issueNumber":"23"}'
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

	fullUrl := fmt.Sprintf("%s/repos/%s/%s/issues%s?access_token=%s", githubApiUrlBase, user, repository, issueNumber, accessToken)
	fmt.Println(getJsonResponse(fullUrl))
}

func issueComments() {

	var paramMap map[string]interface{}
	json.Unmarshal([]byte(*paramsFlag), &paramMap)
	value, ok := paramMap["issueNumber"]
	if !ok {
		log.Fatal("-p {\"issueNumber\":\"1234\"} is required.")
	}
	var issueNumber string
	if value != nil {
		issueNumber = "/" + value.(string)
	}

	fullUrl := fmt.Sprintf("%s/repos/%s/%s/issues%s/comments?access_token=%s", githubApiUrlBase, user, repository, issueNumber, accessToken)
	fmt.Println(getJsonResponse(fullUrl))
}

func getJsonResponse(url string) string {
	resp, err := http.Get(url)

	// Response handling
	body, err := ioutil.ReadAll(resp.Body)
	var result interface{}
	json.Unmarshal(body, &result)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Panic("resp.Body.Close() failed.")
		}
	}()
	var buf bytes.Buffer
	err = json.Indent(&buf, body, "", "  ")
	if err != nil {
		panic(err)
	}
	return buf.String()
}
