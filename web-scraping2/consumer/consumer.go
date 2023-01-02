package consumer

// Consumer Consumer
type Consumer struct {
	Channel <-chan string
	scrapingFunc func(string, string, string)
	isClosed bool
}

// NewConsumer NewConsumer
func NewConsumer(ch <-chan string, scrapingFunc func(string, string, string) ) *Consumer {
	c := new(Consumer)
	c.Channel = ch
	c.scrapingFunc = scrapingFunc
	c.isClosed = false
	return c
}

// Consume Consume
func (c Consumer) Consume() {

	if c.Channel == nil {
		panic("Channel is nil.")
	}

	for v := range c.Channel {
		c.scrapingFunc(v)
	}
}
