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
//-----
// $ for ((i=0; i < 100; i++)); do openssl rand -base64 12 | fold -w 10 | head -1; done |xargs -I{} curl http://localhost:9090/channel1?name={}
// $ for ((i=0; i < 100; i++)); do openssl rand -base64 12 | fold -w 10 | head -1; done |xargs -I{} curl http://localhost:9090/channel2?name={}
// $ curl http://localhost:9090/done
// $ curl http://localhost:9090/done
// $ curl http://localhost:9090/done
// $ curl http://localhost:9090/close
func main() {

	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}

	// channel capacity = 100
	channel1 := make(chan string, 100)
	channel2 := make(chan string, 100)
	channelDone := make(chan bool, 100)
	producer1 := producer.NewProducer(channel1)
	producer2 := producer.NewProducer(channel2)

	fmt.Printf("Consuming workers : %v\n", opts.WorkerCount)
	for i := 0; i < opts.WorkerCount; i++ {
		consumer := consumer.NewConsumer(channel1, channel2, channelDone)
		go consumer.Consume()
	}

	// Add request
	// curl http://localhost:9090?name=xshoji
	http.HandleFunc("/channel1", func(w http.ResponseWriter, r *http.Request) {
		// Get query string as single value
		name := r.URL.Query().Get("name")
		producer1.Produce(name)
	})
	// Add request
	// curl http://localhost:9090?name=xshoji
	http.HandleFunc("/channel2", func(w http.ResponseWriter, r *http.Request) {
		// Get query string as single value
		name := r.URL.Query().Get("name")
		producer2.Produce(name)
	})
	// Add request
	// curl http://localhost:9090?name=xshoji
	http.HandleFunc("/done", func(w http.ResponseWriter, r *http.Request) {
		channelDone <- true
	})

	// Close request
	// curl http://localhost:9090/close
	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		close(channel1)
		close(channel2)
		close(channelDone)
		time.Sleep(time.Duration(100) * time.Millisecond)
		fmt.Println("Exit server.")
		os.Exit(0)
	})

	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
