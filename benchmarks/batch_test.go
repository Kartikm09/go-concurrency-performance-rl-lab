package benchmarks

import (
	"encoding/json"
	"testing"
)

func BenchmarkJSONBatch(b *testing.B) {
	values := make([]map[string]any, 100)
	for i := range values {
		values[i] = map[string]any{"id": i, "payload": "synthetic"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(values)
	}
}
