package retry

import (
	"context"
	"errors"
	"testing"
)

type classifier bool

func (c classifier) Retryable(error) bool { return bool(c) }

type recorder struct{ attempts []int }

func (r *recorder) Wait(_ context.Context, attempt int) error {
	r.attempts = append(r.attempts, attempt)
	return nil
}
func TestDeterministicAttempts(t *testing.T) {
	transient := errors.New("transient")
	calls := 0
	waits := &recorder{}
	err := Do(context.Background(), 3, classifier(true), waits, func(context.Context) error {
		calls++
		if calls < 3 {
			return transient
		}
		return nil
	})
	if err != nil || calls != 3 || len(waits.attempts) != 2 {
		t.Fatalf("err=%v calls=%d waits=%v", err, calls, waits.attempts)
	}
}

func TestNonpositiveAttemptBudgetIsRejected(t *testing.T) {
	for _, attempts := range []int{0, -1} {
		calls := 0
		waits := &recorder{}
		err := Do(context.Background(), attempts, classifier(true), waits, func(context.Context) error { calls++; return nil })
		if !errors.Is(err, ErrInvalidAttempts) || calls != 0 || len(waits.attempts) != 0 {
			t.Fatalf("attempts=%d err=%v calls=%d", attempts, err, calls)
		}
	}
}
