package retry

import (
	"context"
	"testing"
)

type successTransport struct{}

func (successTransport) Execute(context.Context) error { return nil }
func TestRegressionImmediateSuccess(t *testing.T) {
	if err := Do(context.Background(), 3, successTransport{}, neverRetry{}, failBackoff{}); err != nil {
		t.Fatal(err)
	}
}
