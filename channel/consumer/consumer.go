package consumer

import "fmt"

// Consumer Consumer
type Consumer struct {
	Channel <-chan string
}

// NewConsumer NewConsumer
func NewConsumer(ch <-chan string) *Consumer {
	c := new(Consumer)
	c.Channel = ch
	return c
}

// Consume Consume
func (c Consumer) Consume() {

	if c.Channel == nil {
		panic("Channel is nil.")
	}

	for v := range c.Channel {
		fmt.Println("Consumer recieved: " + v)
	}
}
