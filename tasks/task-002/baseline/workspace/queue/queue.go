package queue

type Queue struct{ items []string }

func New() *Queue                    { return &Queue{} }
func (q *Queue) Enqueue(item string) { q.items = append(q.items, item) }
func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}
