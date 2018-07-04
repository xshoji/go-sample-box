package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/net/client"
)

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func main() {
	client := client.NewClient("https://blockchain.info//latestblock")
	fmt.Println(client.GetAsString())
}
