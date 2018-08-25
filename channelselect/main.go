package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/xshoji/go-sample-box/channelselect/consumer"
	"github.com/xshoji/go-sample-box/channelselect/producer"
)

type options struct {
	WorkerCount int `short:"w" long:"workers" description:"Consuming worker count" default:"3"`
}

// - [Go の channel 処理パターン集 · Hori Blog](https://hori-ryota.com/blog/golang-channel-pattern/)
// - [GoのChannelを使いこなせるようになるための手引 - Qiita](https://qiita.com/awakia/items/f8afa070c96d1c9a04c9)
func main() {

	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}

	// channel capacity = 100
	channel := make(chan string, 100)
	producer := producer.NewProducer(channel)

	fmt.Printf("Consuming workers : %v\n", opts.WorkerCount)
	for i := 0; i < opts.WorkerCount; i++ {
		consumer := consumer.NewConsumer(channel)
		go consumer.Consume()
	}

	// Add request
	// curl http://localhost:9090?name=xshoji
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get query string as single value
		name := r.URL.Query().Get("name")
		producer.Produce(name)
	})

	// Close request
	// curl http://localhost:9090/close
	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		close(channel)
		time.Sleep(time.Duration(100) * time.Millisecond)
		fmt.Println("Exit server.")
		os.Exit(0)
	})

	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
