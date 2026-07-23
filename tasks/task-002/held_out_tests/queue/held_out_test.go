package queue

import "testing"

func TestHeldOutFIFOAfterOverload(t *testing.T) {
	q := New(2)
	_ = q.Enqueue("a")
	_ = q.Enqueue("b")
	_ = q.Enqueue("c")
	first, _ := q.Dequeue()
	second, _ := q.Dequeue()
	if first != "a" || second != "b" {
		t.Fatalf("%s %s", first, second)
	}
}
