// Package httpapi exposes synthetic webhook ingestion.
package httpapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/dedupe"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/metrics"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/queue"
)

type Handler struct {
	queue   *queue.Queue
	dedupe  *dedupe.Set
	metrics *metrics.Counters
}

func New(q *queue.Queue, d *dedupe.Set, m *metrics.Counters) *Handler { return &Handler{q, d, m} }
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("POST /webhooks", h.ingest)
	return mux
}
func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = io.WriteString(w, "{\"status\":\"ok\"}\n")
}
func (h *Handler) ingest(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 64<<10)
	var request struct {
		ID      string `json:"id"`
		Payload string `json:"payload"`
	}
	if json.NewDecoder(r.Body).Decode(&request) != nil || request.ID == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if !h.dedupe.First(request.ID) {
		w.WriteHeader(http.StatusOK)
		return
	}
	if err := h.queue.Enqueue(queue.Job{ID: request.ID, Payload: []byte(request.Payload)}); err != nil {
		h.metrics.Rejected()
		http.Error(w, "overloaded", http.StatusTooManyRequests)
		return
	}
	h.metrics.Accepted()
	w.WriteHeader(http.StatusAccepted)
}
