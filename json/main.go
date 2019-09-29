package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/xshoji/go-sample-box/json/client"
)

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func main() {
	clientLatestblock := client.NewClient("https://blockchain.info/latestblock")

	// Print by object
	block := clientLatestblock.GetAsObject()
	fmt.Printf("hash: %v\n", block.Hash)
	fmt.Printf("time: %.f\n", block.Time)
	fmt.Printf("height: %.f\n", block.Height)
	fmt.Printf("block_index: %.f\n", block.BlockIndex)

	// Print as json
	jsonBytes, err := json.Marshal(block)
	if err != nil {
		log.Panic("error")
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	// > GoでJSON文字列を整形し直す - 年中アイス
	// > https://reiki4040.hatenablog.com/entry/2017/08/30/070000
	// Print as pretty json
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(jsonString), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
