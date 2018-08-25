package consumer

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
)

// Consumer Consumer
type Consumer struct {
	Channel   <-chan string
	Identifir string
}

// NewConsumer NewConsumer
func NewConsumer(ch <-chan string) *Consumer {
	c := new(Consumer)
	c.Channel = ch
	c.Identifir = getMD5Hash(strconv.FormatFloat(rand.Float64(), 'f', 4, 32))
	return c
}

// - [Golang - How to hash a string using MD5.](https://gist.github.com/sergiotapia/8263278)
func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Consume Consume
func (c Consumer) Consume() {

	fmt.Printf("[%v] Consumer started.\n", c.Identifir)
	if c.Channel == nil {
		panic("Channel is nil.")
	}

	for v := range c.Channel {
		fmt.Printf("[%v] Consumer recieved: %v\n", c.Identifir, v)
	}
	fmt.Printf("[%v] Consumer channle closed.\n", c.Identifir)
}
