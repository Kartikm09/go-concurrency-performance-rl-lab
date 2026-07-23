package batch

import "testing"

func TestHeldOutEmptyAndLockMetric(t *testing.T) {
	encoder := &Encoder{}
	output, err := encoder.Encode(nil)
	if err != nil || len(output) != 0 {
		t.Fatalf("output=%q err=%v", output, err)
	}
	if encoder.LockCount() > 1 {
		t.Fatalf("locks=%d", encoder.LockCount())
	}
}
