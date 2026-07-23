package retry

import (
	"context"
	"errors"
	"testing"
)

type transport struct {
	calls int
	fail  error
}

func (t *transport) Execute(context.Context) error {
	t.calls++
	if t.calls < 3 {
		return t.fail
	}
	return nil
}

type classifier struct{}

func (classifier) Retryable(error) bool { return true }

type backoff struct{ attempts []int }

func (b *backoff) Wait(_ context.Context, a int) error {
	b.attempts = append(b.attempts, a)
	return nil
}
func TestPublicRetryContract(t *testing.T) {
	tr := &transport{fail: errors.New("temporary")}
	bo := &backoff{}
	if err := Do(context.Background(), 3, tr, classifier{}, bo); err != nil || tr.calls != 3 || len(bo.attempts) != 2 || bo.attempts[0] != 1 {
		t.Fatalf("err=%v calls=%d waits=%v", err, tr.calls, bo.attempts)
	}
}
