package main

import "context"
import "golang.org/x/sync/errgroup"

func main() {
	rootCtx := context.Background()
	_, _ = errgroup.WithContext(rootCtx)
}
