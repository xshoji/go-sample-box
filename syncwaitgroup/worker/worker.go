package worker

import (
	"math/rand"
	"time"

	"fmt"
	"sync"
)

// Worker Worker
type Worker struct {
	Name      string
	WaitGroup *sync.WaitGroup
}

// NewWorker NewWorker
func NewWorker(name string, wg *sync.WaitGroup) *Worker {
	w := new(Worker)
	w.WaitGroup = wg
	w.Name = name
	return w
}

// Start Start
func (w Worker) Start() {
	// > go - golang sync.WaitGroup never completes - Stack Overflow
	// > https://stackoverflow.com/questions/27893304/golang-sync-waitgroup-never-completes
	defer w.WaitGroup.Done()
	// - [Goでrandを使うときは忘れずにSeedを設定しないといけない - Qiita](https://qiita.com/makiuchi-d/items/9c4af327bc8502cdcdce)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		// - [Golang Cookbook： generate random number in a given range](http://golangcookbook.blogspot.com/2012/11/generate-random-number-in-given-range.html)
		millsec := rand.Intn(1000-1) + 1
		time.Sleep(time.Duration(millsec) * time.Millisecond)
		fmt.Printf("name: %s, i: %d, millsec: %d\n", w.Name, i, millsec)
	}
}
