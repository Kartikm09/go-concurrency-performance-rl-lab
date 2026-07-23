package main

import (
	"context"
	"fmt"

	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/worker"
)

func main() {
	pool := worker.New(context.Background(), 2, 8)
	_ = pool.Submit(worker.Job{ID: "demo", Run: func(context.Context) error { fmt.Println("completed demo"); return nil }})
	pool.Stop()
}
