package pool

import "testing"

func TestRegressionSubmitAfterStop(t *testing.T) {
	p := New(1)
	p.Stop()
	if err := p.Submit(nil); err != ErrStopped {
		t.Fatalf("err=%v", err)
	}
}
