package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log.Printf("PID: %d\n", os.Getpid())
	f, _ := os.CreateTemp("", "*_defer_safety_sample.txt")
	log.Printf("temp file: %s\n", f.Name())
	_ = os.WriteFile(f.Name(), []byte("aaa"), 0655)

	// Successful handling
	defer os.Remove(f.Name())
	// Signal handling
	go func() {
		// 1. Create channel for os.Signal and Notify interrupt signal to channel.
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM) // syscall.SIGINT, syscall.SIGTERM ( [!] Cannot handle os.Kill )

		// 2. Wait signal
		s := <-signalChannel
		log.Printf("Signal is received: %v\n", s)
		os.Remove(f.Name())
		signalValue := int(s.(syscall.Signal))
		os.Exit(signalValue)
	}()

	// Wait
	duration := 1_000 * time.Millisecond
	for i := 0; i >= 0; i++ {
		log.Printf("Sleep: i=%d, totalSecond=%v\n", i, time.Duration(i)*duration)
		time.Sleep(duration)
	}
}
