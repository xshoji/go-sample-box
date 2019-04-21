package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// > 【Go言語】Ctrl+cなどによるSIGINTの捕捉とdeferの実行 - DRYな備忘録
// > http://otiai10.hatenablog.com/entry/2018/02/19/165228
func main() {

	//-------------------------------
	// Define defer process
	deferfunc := func() {
		log.Println("Call defer function")
	}
	defer deferfunc()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)

	go func() {
		<-signals
		fmt.Println("")
		log.Println("Catch signals")
		deferfunc()
		log.Println("Execute os.Exit()")
		os.Exit(0)
	}()

	//-------------------------------
	// Define main logic
	for i := 0; i < 5; i++ {
		log.Printf("Loop count is %v\n", i)
		duration := time.Duration(1000 * time.Millisecond)
		time.Sleep(duration)
	}

	log.Println("Finish loop")
}
