// Package queue provides a bounded, concurrency-safe FIFO.
package queue

import (
	"errors"
	"sync"
)

var ErrOverloaded = errors.New("queue capacity reached")

type Job struct {
	ID      string
	Payload []byte
}
type Stats struct {
	Accepted uint64
	Rejected uint64
	Depth    int
}
type Queue struct {
	mu       sync.Mutex
	capacity int
	jobs     []Job
	accepted uint64
	rejected uint64
}

func New(capacity int) *Queue {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	return &Queue{capacity: capacity, jobs: make([]Job, 0, capacity)}
}
func (q *Queue) Enqueue(job Job) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == q.capacity {
		q.rejected++
		return ErrOverloaded
	}
	job.Payload = append([]byte(nil), job.Payload...)
	q.jobs = append(q.jobs, job)
	q.accepted++
	return nil
}
func (q *Queue) Dequeue() (Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == 0 {
		return Job{}, false
	}
	job := q.jobs[0]
	copy(q.jobs, q.jobs[1:])
	q.jobs = q.jobs[:len(q.jobs)-1]
	return job, true
}
func (q *Queue) Stats() Stats {
	q.mu.Lock()
	defer q.mu.Unlock()
	return Stats{q.accepted, q.rejected, len(q.jobs)}
}
