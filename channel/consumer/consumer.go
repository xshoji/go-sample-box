package consumer

import "fmt"

// Consumer Consumer
type Consumer struct {
	Channel <-chan string
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
