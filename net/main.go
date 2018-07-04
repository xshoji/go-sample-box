package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func main() {
	resp, err := http.Get("https://blockchain.info//latestblock")
	if err != nil {
		log.Panic("Error occured. | ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Panic("Status is not OK. | ", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
