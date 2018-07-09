package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/xshoji/go-sample-box/sync/worker"
)

func main() {
	// - [syncパッケージを利用した同期処理 - はじめてのGo言語](http://cuto.unirita.co.jp/gostudy/post/go_sync/)
	workerGroup := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wk := worker.NewWorker(strings.Join([]string{"worker_", strconv.Itoa(i)}, ""), workerGroup)
		workerGroup.Add(1)
		go wk.Start()
	}
	workerGroup.Wait()
}
