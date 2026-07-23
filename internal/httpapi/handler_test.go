package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/dedupe"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/metrics"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/queue"
)

func TestIngestionContract(t *testing.T) {
	handler := New(queue.New(1), dedupe.New(), &metrics.Counters{}).Routes()
	request := httptest.NewRequest(http.MethodPost, "/webhooks", bytes.NewBufferString(`{"id":"evt-1","payload":"demo"}`))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusAccepted {
		t.Fatalf("status=%d", response.Code)
	}
}
func FuzzWebhookHandler(f *testing.F) {
	f.Add([]byte(`{"id":"seed","payload":"x"}`))
	f.Fuzz(func(t *testing.T, body []byte) {
		if len(body) > 70000 {
			t.Skip()
		}
		handler := New(queue.New(2), dedupe.New(), &metrics.Counters{}).Routes()
		request := httptest.NewRequest(http.MethodPost, "/webhooks", bytes.NewReader(body))
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if response.Code < 100 || response.Code > 599 {
			t.Fatalf("invalid status %d", response.Code)
		}
	})
}
