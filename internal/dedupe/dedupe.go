// Package dedupe tracks bounded synthetic webhook identifiers.
package dedupe

import "sync"

type Set struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

func New() *Set { return &Set{seen: make(map[string]struct{})} }
func (s *Set) First(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.seen[id]; ok {
		return false
	}
	s.seen[id] = struct{}{}
	return true
}
