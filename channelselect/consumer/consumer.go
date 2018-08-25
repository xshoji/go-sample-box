package consumer

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Consumer Consumer
type Consumer struct {
	Channel1    <-chan string
	Channel2    <-chan string
	ChannelDone <-chan bool
	Identifir   string
}

// NewConsumer NewConsumer
func NewConsumer(ch <-chan string, ch2 <-chan string, chDone <-chan bool) *Consumer {
	c := new(Consumer)
	c.Channel1 = ch
	c.Channel2 = ch2
	c.ChannelDone = chDone
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
	if c.Channel1 == nil {
		panic("Channel1 is nil.")
	}
	if c.Channel2 == nil {
		panic("Channel2 is nil.")
	}
	if c.ChannelDone == nil {
		panic("ChannelDone is nil.")
	}

	// for v := range c.Channel {
	// 	duration := time.Duration(rand.Intn(1000-1)+1) * time.Millisecond
	// 	fmt.Printf("[%v] Consumer sleep: %v\n", c.Identifir, duration)
	// 	time.Sleep(duration)
	// 	fmt.Printf("[%v] Consumer recieved: %v\n", c.Identifir, v)
	// }
	printMessage := func(message string, isSuccessRecieved bool, number string) {
		if isSuccessRecieved == false {
			fmt.Printf("[%v] Consumer [Channel%v] exit.\n", c.Identifir, number)
		}
		duration := time.Duration(rand.Intn(1000-1)+1) * time.Millisecond
		//		fmt.Printf("[%v] Consumer sleep: %v\n", c.Identifir, duration)
		time.Sleep(duration)
		fmt.Printf("[%v] Consumer [Channel%v] [sleep:%v] recieved: %v\n", c.Identifir, number, duration, message)
	}
	func() {
		for {
			select {
			case v, ok := <-c.Channel1:
				printMessage(v, ok, "1")
			case v, ok := <-c.Channel2:
				printMessage(v, ok, "2")
			case v := <-c.ChannelDone:
				if v == true {
					fmt.Printf("[%v] Consumer exit.\n", c.Identifir)
					return
				}
			}
		}
	}()
}
