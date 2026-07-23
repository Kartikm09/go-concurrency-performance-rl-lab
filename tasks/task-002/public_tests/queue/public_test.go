package queue

import "testing"

func TestPublicBackpressureAndMetrics(t *testing.T) {
	q := New(1)
	if err := q.Enqueue("a"); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue("b"); err != ErrOverloaded {
		t.Fatalf("err=%v", err)
	}
	stats := q.Stats()
	if stats.Accepted != 1 || stats.Rejected != 1 || stats.Depth != 1 {
		t.Fatalf("stats=%+v", stats)
	}
}
