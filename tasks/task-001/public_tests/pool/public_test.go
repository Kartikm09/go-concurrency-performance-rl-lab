package pool

import (
	"testing"
	"time"
)

func TestPublicIdleStopIsBounded(t *testing.T) {
	p := New(4)
	done := make(chan struct{})
	go func() { p.Stop(); close(done) }()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("stop leaked idle workers")
	}
	p.Stop()
}
