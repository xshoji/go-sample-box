package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// If being piped to stdin, read stdin as sentence input.
	// > go - Check if there is something to read on STDIN in Golang - Stack Overflow
	// > https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang/26567513#26567513
	pipeInput := ""
	stat, _ := os.Stdin.Stat()
	// Check pipe input.
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// In case of the pipe input, read bytes from stdin.
		inputBytes, _ := ioutil.ReadAll(os.Stdin)
		pipeInput = string(inputBytes)
	}

	if pipeInput == "" {
		fmt.Printf("No input.\n")
		os.Exit(1)
	}

	fmt.Printf("%v\n", pipeInput)
}
