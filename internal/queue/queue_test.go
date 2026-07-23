package queue

import (
	"errors"
	"testing"
)

func TestQueueBackpressureAndFIFO(t *testing.T) {
	q := New(1)
	if err := q.Enqueue(Job{ID: "1"}); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue(Job{ID: "2"}); !errors.Is(err, ErrOverloaded) {
		t.Fatalf("want overload, got %v", err)
	}
	job, ok := q.Dequeue()
	if !ok || job.ID != "1" {
		t.Fatalf("unexpected job: %#v", job)
	}
}
func BenchmarkQueueRoundTrip(b *testing.B) {
	q := New(1)
	for i := 0; i < b.N; i++ {
		_ = q.Enqueue(Job{ID: "x"})
		_, _ = q.Dequeue()
	}
}
