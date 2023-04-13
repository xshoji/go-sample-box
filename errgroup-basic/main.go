package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"os"
	"time"
)

const UsageParamNameWidth = "10"

var (
	// Define short parameters
	paramsGroups        = flag.Int("g", 5, "errgroups count")
	paramsMaxWaitSecond = flag.Int("m", 10, "max wait second")
	paramsHelp          = flag.Bool("h", false, "\nhelp")
)

func init() {
	flag.Parse()
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}

}

// Go の goroutine / channel は全然簡単じゃないので errgroup を使おう - 音速きなこおはぎ
// https://eihigh.hatenablog.com/entry/2023/04/08/220538
func main() {

	fmt.Println("errgroups count:", *paramsGroups)
	fmt.Println("max wait second:", *paramsMaxWaitSecond)

	rootCtx := context.Background()
	errGroup, _ := errgroup.WithContext(rootCtx)

	createRandomNumber := func() int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(*paramsMaxWaitSecond-1) + 1
	}

	for i := 0; i < *paramsGroups; i++ {
		// cmd/compile: add GOEXPERIMENT=loopvar · Issue #57969 · golang/go
		// https://github.com/golang/go/issues/57969
		v := i
		errGroup.Go(func() error {
			duration := time.Duration(createRandomNumber()) * time.Second
			fmt.Printf("%d: Duration=%v\n", v, duration)
			time.Sleep(duration)
			fmt.Printf("%d: Completed sleep! (duration:%v)\n", v, duration)
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		log.Fatal(err)
	}
}
