// Package store provides deterministic in-memory delivery records.
package store

import "sync"

type Memory struct {
	mu        sync.Mutex
	completed []string
	dead      []string
}

func (m *Memory) Complete(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.completed = append(m.completed, id)
}
func (m *Memory) DeadLetter(id string) { m.mu.Lock(); defer m.mu.Unlock(); m.dead = append(m.dead, id) }
func (m *Memory) Snapshot() ([]string, []string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]string(nil), m.completed...), append([]string(nil), m.dead...)
}
