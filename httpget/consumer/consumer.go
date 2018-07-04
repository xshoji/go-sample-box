package consumer

import (
	"fmt"

	"github.com/xshoji/go-sample-box/httpget/client"
)

// Consumer Consumer
type Consumer struct {
	Channel  <-chan string
	ClientTx *client.Client
}

// NewConsumer NewConsumer
func NewConsumer(ch <-chan string, clientTx *client.Client) *Consumer {
	c := new(Consumer)
	c.Channel = ch
	c.ClientTx = clientTx
	return c
}

// Consume Consume
func (c Consumer) Consume() {

	if c.Channel == nil {
		panic("Channel is nil.")
	}

	// - [go - How to check for an empty struct? - Stack Overflow](https://stackoverflow.com/questions/28447297/how-to-check-for-an-empty-struct/28447372#28447372)
	// - [null - nil detection in Go - Stack Overflow](https://stackoverflow.com/questions/20240179/nil-detection-in-go)
	if c.ClientTx == nil {
		panic("Client is empty.")
	}

	var result interface{}
	for v := range c.Channel {
		fmt.Println("Consumer is recieved: ", v)
		result = c.ClientTx.GetWithPathAsObject(v)
		fmt.Printf(
			"hash: %v, time: %.f, tx_index: %.f\n",
			result.(map[string]interface{})["hash"].(string),
			result.(map[string]interface{})["time"].(float64),
			result.(map[string]interface{})["tx_index"].(float64),
		)
	}
}
