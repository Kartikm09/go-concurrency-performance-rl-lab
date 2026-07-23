package queue

import "testing"

func TestRegressionInvalidCapacity(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	_ = New(0)
}
