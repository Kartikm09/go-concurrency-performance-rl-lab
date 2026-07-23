package batch

import (
	"bytes"
	"encoding/json"
	"sync"
)

type Event struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}
type Encoder struct {
	mu        sync.Mutex
	lockCount int
}

func (e *Encoder) Encode(events []Event) ([]byte, error) {
	var output bytes.Buffer
	for _, event := range events {
		e.mu.Lock()
		e.lockCount++
		encoded, err := json.Marshal(event)
		if err == nil {
			output.Write(encoded)
			output.WriteByte('\n')
		}
		e.mu.Unlock()
		if err != nil {
			return nil, err
		}
	}
	return output.Bytes(), nil
}
func (e *Encoder) LockCount() int { e.mu.Lock(); defer e.mu.Unlock(); return e.lockCount }
