package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Define channel
	channel := make(chan string, 100)

	// Add value to channel
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Add value to channel : %v\n", i)
			channel <- strconv.FormatInt(int64(i), 10)
		}
		fmt.Println("Channel is closed")
		close(channel)
	}()

	// Define wait group
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1) // Add worker
	go func() {
		// Do loop until channel is closed
		for v := range channel {
			duration := time.Duration(1000 * time.Millisecond)
			time.Sleep(duration)
			fmt.Printf("Goroutine recieve %v\n", v)
		}
		// WaitGroup is done when channel is closed
		waitGroup.Done()
	}()

	// Wait waitGroup
	waitGroup.Wait()
}
