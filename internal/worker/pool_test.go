package worker

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestIdlePoolStopsWithoutLeak(t *testing.T) {
	pool := New(context.Background(), 4, 4)
	done := make(chan struct{})
	go func() { pool.Stop(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("stop timed out")
	}
}
func TestAcceptedJobRunsAtMostOnce(t *testing.T) {
	pool := New(context.Background(), 1, 1)
	var count atomic.Int64
	if err := pool.Submit(Job{ID: "x", Run: func(context.Context) error { count.Add(1); return nil }}); err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Millisecond)
	pool.Stop()
	if count.Load() != 1 {
		t.Fatalf("count=%d", count.Load())
	}
}
