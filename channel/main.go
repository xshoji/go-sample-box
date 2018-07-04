package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/channel/consumer"
	"github.com/xshoji/go-sample-box/channel/producer"
)

func main() {
	channel := make(chan string)
	producer := producer.NewProducer(channel)
	consumer := consumer.NewConsumer(channel)

	fmt.Println("Consumer start")
	go consumer.Consume()

	fmt.Println("Produce start")
	for _, message := range []string{"hello", "my", "name", "is", "taro"} {
		producer.Produce(message)
	}
}
