package pool

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
)

func TestHeldOutConcurrentSubmitAndStop(t *testing.T) {
	p := New(2)
	var ran atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); _ = p.Submit(func(context.Context) { ran.Add(1) }) }()
	}
	p.Stop()
	wg.Wait()
	if ran.Load() > 20 {
		t.Fatalf("ran=%d", ran.Load())
	}
}
