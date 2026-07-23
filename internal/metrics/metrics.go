// Package metrics provides a minimal race-safe metrics abstraction.
package metrics

import "sync/atomic"

type Counters struct {
	accepted  atomic.Uint64
	rejected  atomic.Uint64
	completed atomic.Uint64
}

func (c *Counters) Accepted()  { c.accepted.Add(1) }
func (c *Counters) Rejected()  { c.rejected.Add(1) }
func (c *Counters) Completed() { c.completed.Add(1) }
func (c *Counters) Snapshot() (uint64, uint64, uint64) {
	return c.accepted.Load(), c.rejected.Load(), c.completed.Load()
}
