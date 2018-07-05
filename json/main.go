package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/json/client"
)

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func main() {
	clientLatestblock := client.NewClient("https://blockchain.info/latestblock")
	block := clientLatestblock.GetAsObject()
	fmt.Printf("hash: %v\n", block.Hash)
	fmt.Printf("time: %.f\n", block.Time)
	fmt.Printf("height: %.f\n", block.Height)
	fmt.Printf("block_index: %.f\n", block.BlockIndex)
	fmt.Printf("txIndexesCount: %d\n", len(block.TxIndexes))
}
