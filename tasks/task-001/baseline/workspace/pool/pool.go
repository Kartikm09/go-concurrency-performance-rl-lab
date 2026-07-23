package pool

import (
	"context"
	"errors"
	"sync"
)

var ErrStopped = errors.New("stopped")

type Job func(context.Context)
type Pool struct {
	ctx     context.Context
	cancel  context.CancelFunc
	jobs    chan Job
	stopped bool
	wg      sync.WaitGroup
}

func New(workers int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool{ctx: ctx, cancel: cancel, jobs: make(chan Job)}
	p.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go p.work()
	}
	return p
}
func (p *Pool) Submit(job Job) error {
	if p.stopped {
		return ErrStopped
	}
	select {
	case p.jobs <- job:
		return nil
	default:
		return errors.New("busy")
	}
}
func (p *Pool) Stop() { p.stopped = true; p.cancel(); p.wg.Wait() }
func (p *Pool) work() {
	defer p.wg.Done()
	for job := range p.jobs {
		job(p.ctx)
	}
}
