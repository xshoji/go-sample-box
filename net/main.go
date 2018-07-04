package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/net/client"
)

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func main() {
	client := client.NewClient("https://blockchain.info//latestblock")
	fmt.Println("GetAsString")
	fmt.Println(client.GetAsString())
	fmt.Println("")

	fmt.Println("GetAsObject")
	result := client.GetAsObject()
	fmt.Printf("%#v\n", result)
	fmt.Println("")

	fmt.Println("GetAsObject parsed")
	fmt.Printf(
		"hash: %v\ntime: %.f\nblock_index: %.f\nheight: %.f\n",
		result.(map[string]interface{})["hash"].(string),
		result.(map[string]interface{})["time"].(float64),
		result.(map[string]interface{})["block_index"].(float64),
		result.(map[string]interface{})["height"].(float64),
	)
}
