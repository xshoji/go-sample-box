package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// > ［Go］ SIGINT受信時にも後片付けをしたい｜tabatak｜note
// > https://note.com/kltl/n/n3a8c00049f92
func main() {

	// Define defer process
	defer func() {
		log.Println("Call defer function")
	}()

	// Create channels
	// 1. Create channel for os.Signal and Notify interrupt signal to channel.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	// 2. Create channel for done main process.
	doneChannel := make(chan interface{})

	// Execute main process.
	go DoMain(doneChannel)

	// Select each channels.
	select {
	case <-doneChannel:
		log.Printf("Complete.\n")
	case s := <-signalChannel:
		fmt.Println()
		log.Printf("Signal is received: %v\n", s)
	}

}

func DoMain(doneChannel chan interface{}) {
	for i := 0; i < 5; i++ {
		log.Printf("Loop count is %v\n", i)
		duration := time.Duration(1000 * time.Millisecond)
		time.Sleep(duration)
	}
	doneChannel <- struct{}{}
}
