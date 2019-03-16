package producer

import "fmt"

// Producer Producer
type Producer struct {
	Channel chan<- string
}

// NewProducer NewProducer
func NewProducer(ch chan<- string) *Producer {
	p := new(Producer)
	p.Channel = ch
	return p
}

// Produce Produce
func (p Producer) Produce(value string) {

	if p.Channel == nil {
		panic("Channel is nil.")
	}

	fmt.Println("Producer produce: " + value)
	p.Channel <- value
}
