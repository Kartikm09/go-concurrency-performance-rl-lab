// Package dedupe tracks process-local synthetic webhook admissions.
package dedupe

import "sync"

type Set struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

func New() *Set { return &Set{seen: make(map[string]struct{})} }
func (s *Set) First(id string) bool {
	first, _ := s.Admit(id, func() error { return nil })
	return first
}

// Admit remembers id only after accept succeeds, so rejected work remains retryable.
// accept runs under the set lock and must be bounded and must not reenter this Set.
func (s *Set) Admit(id string, accept func() error) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.seen[id]; ok {
		return false, nil
	}
	if err := accept(); err != nil {
		return false, err
	}
	s.seen[id] = struct{}{}
	return true, nil
}
