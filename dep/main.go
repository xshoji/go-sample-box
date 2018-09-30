package main

import (
	"github.com/franela/goreq"
	"fmt"
	"io/ioutil"
)

//-------------------
// [setup]
// dep init
// dep ensure -add "github.com/franela/goreq"
//
// [build]
// $ dep ensure
// $ go run main.go
//-------------------
func main() {

	param := &struct {
		Name string
		Age int
	}{}
	param.Name = "taro"
	param.Age = 10

	res, err := goreq.Request{
		Uri: "http://httpbin.org/get",
		QueryString: param,
	}.Do()

	if err != nil {
		panic("ERROR")
	}

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(b))

}