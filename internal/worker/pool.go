// Package worker owns goroutine lifecycle and bounded submission.
package worker

import (
	"context"
	"errors"
	"sync"
)

var ErrStopped = errors.New("worker pool stopped")
var ErrOverloaded = errors.New("worker queue full")

type Job struct {
	ID  string
	Run func(context.Context) error
}
type Pool struct {
	ctx     context.Context
	cancel  context.CancelFunc
	jobs    chan Job
	mu      sync.RWMutex
	stopped bool
	wg      sync.WaitGroup
}

func New(parent context.Context, workers, capacity int) *Pool {
	ctx, cancel := context.WithCancel(parent)
	pool := &Pool{ctx: ctx, cancel: cancel, jobs: make(chan Job, capacity)}
	pool.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go pool.work()
	}
	return pool
}
func (p *Pool) Submit(job Job) error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.stopped {
		return ErrStopped
	}
	select {
	case p.jobs <- job:
		return nil
	default:
		return ErrOverloaded
	}
}
func (p *Pool) Stop() {
	p.mu.Lock()
	if p.stopped {
		p.mu.Unlock()
		return
	}
	p.stopped = true
	p.cancel()
	p.mu.Unlock()
	p.wg.Wait()
}
func (p *Pool) work() {
	defer p.wg.Done()
	for {
		select {
		case <-p.ctx.Done():
			return
		case job := <-p.jobs:
			if job.Run != nil {
				_ = job.Run(p.ctx)
			}
		}
	}
}
