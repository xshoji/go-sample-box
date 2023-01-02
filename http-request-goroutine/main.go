package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/httprequest-goroutine/client"
	"github.com/xshoji/go-sample-box/httprequest-goroutine/consumer"
)

const (
	// ChannelCapacity ChannelCapacity
	ChannelCapacity = 1
	// ConcurrentCount ConcurrentCount
	ConcurrentCount = 1
)

// - [http - The Go Programming Language](https://golang.org/pkg/net/http/)
// - [networking - Access HTTP response as string in Go - Stack Overflow](https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go)
func main() {
	clientLatestblock := client.NewClient("https://blockchain.info/latestblock")
	clientRawblock := client.NewClient("https://blockchain.info/rawblock")
	clientRawTx := client.NewClient("https://blockchain.info/rawtx")

	fmt.Println("GetAsString")
	fmt.Println(clientLatestblock.Get("").GetBody())
	fmt.Println("")

	fmt.Println("GetAsObject")
	resultLatestBlock := clientLatestblock.Get("").GetBodyAsObject()
	fmt.Printf("%#v\n", resultLatestBlock)
	fmt.Println("")

	fmt.Println("GetAsObject parsed")
	fmt.Printf(
		"hash: %v\ntime: %.f\nblock_index: %.f\nheight: %.f\ntxIndexes_count: %d\n",
		resultLatestBlock.(map[string]interface{})["hash"].(string),
		resultLatestBlock.(map[string]interface{})["time"].(float64),
		resultLatestBlock.(map[string]interface{})["block_index"].(float64),
		resultLatestBlock.(map[string]interface{})["height"].(float64),
		len(resultLatestBlock.(map[string]interface{})["txIndexes"].([]interface{})),
	)

	fmt.Println("Get block hash")
	blockHash := resultLatestBlock.(map[string]interface{})["hash"].(string)
	fmt.Println(blockHash)
	fmt.Println("")

	// Fixed Block hash which has many transactions
	// blockHash = "0000000000000000002b7601d833e402abe8d6dd2a8337d00b1ad905c6d10247"

	fmt.Println("Get transaction hashs")
	resultSingleBlock := clientRawblock.Get("/" + blockHash).GetBodyAsObject()
	transactions := resultSingleBlock.(map[string]interface{})["tx"].([]interface{})
	var transactionHashs []string
	for _, transaction := range transactions {
		transactionHashs = append(transactionHashs, transaction.(map[string]interface{})["hash"].(string))
	}
	fmt.Printf("%v", transactionHashs)
	fmt.Println("")

	// - [Go の channel 処理パターン集 · Hori Blog](https://hori-ryota.com/blog/golang-channel-pattern/#%E5%AE%9A%E7%BE%A9%E3%81%AE%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
	txhashChannel := make(chan string, ChannelCapacity)
	// ConcurrentCount数で並列処理する
	var consumers []*consumer.Consumer
	var c *consumer.Consumer
	for i := 0; i < ConcurrentCount; i++ {
		c = consumer.NewConsumer(txhashChannel, clientRawTx)
		go c.Consume()
		consumers = append(consumers, c)
	}

	for _, txHash := range transactionHashs {
		fmt.Println("txHash is send: ", txHash)
		txhashChannel <- txHash
	}
}
