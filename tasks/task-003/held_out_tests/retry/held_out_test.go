package retry

import (
	"context"
	"errors"
	"testing"
)

type onceTransport struct{ err error }

func (t onceTransport) Execute(context.Context) error { return t.err }

type neverRetry struct{}

func (neverRetry) Retryable(error) bool { return false }

type failBackoff struct{}

func (failBackoff) Wait(context.Context, int) error { return errors.New("wait") }
func TestHeldOutNonRetryableStopsOnce(t *testing.T) {
	terminal := errors.New("terminal")
	if err := Do(context.Background(), 5, onceTransport{terminal}, neverRetry{}, failBackoff{}); err != terminal {
		t.Fatalf("err=%v", err)
	}
}
