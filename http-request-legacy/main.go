package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/http-request-legacy/client"
)

func main() {
	clientHttpbin := client.NewClient("http://httpbin.org")
	fmt.Println("[Get]")
	fmt.Println(clientHttpbin.Get("/get").GetBody())
	resultGet := clientHttpbin.Get("/get?param1=aaa&param2=bbb")
	fmt.Println(resultGet.GetBody())
	fmt.Println("origin : " + resultGet.GetBodyAsObject().(map[string]interface{})["origin"].(string))
	fmt.Println("url    : " + resultGet.GetBodyAsObject().(map[string]interface{})["url"].(string))
	fmt.Println("")

	fmt.Println("[Post]")
	resultPost := clientHttpbin.Post("/post", map[string][]string{
		"param1": {"aaa"},
	})
	fmt.Println(resultPost.GetBody())
	fmt.Println("data : " + resultPost.GetBodyAsObject().(map[string]interface{})["data"].(string))
	fmt.Println("")
}
