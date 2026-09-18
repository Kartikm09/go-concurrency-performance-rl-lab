package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync"
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

func TestRejectedWebhookCanRetryAfterQueueDrains(t *testing.T) {
	q := queue.New(1)
	if err := q.Enqueue(queue.Job{ID: "occupant"}); err != nil {
		t.Fatal(err)
	}
	counts := &metrics.Counters{}
	handler := New(q, dedupe.New(), counts).Routes()
	send := func() int {
		request := httptest.NewRequest(http.MethodPost, "/webhooks", bytes.NewBufferString(`{"id":"retry","payload":"kept"}`))
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		return response.Code
	}
	if status := send(); status != http.StatusTooManyRequests {
		t.Fatalf("first status=%d", status)
	}
	q.Dequeue()
	if status := send(); status != http.StatusAccepted {
		t.Fatalf("retry status=%d; failed admissions must not be deduplicated", status)
	}
	if status := send(); status != http.StatusOK {
		t.Fatalf("duplicate status=%d", status)
	}
	job, ok := q.Dequeue()
	if !ok || job.ID != "retry" || string(job.Payload) != "kept" {
		t.Fatalf("lost retry: job=%+v ok=%v", job, ok)
	}
	accepted, rejected, _ := counts.Snapshot()
	if accepted != 1 || rejected != 1 {
		t.Fatalf("accepted=%d rejected=%d", accepted, rejected)
	}
}

func TestConcurrentWebhookDuplicatesAreAdmittedOnce(t *testing.T) {
	q := queue.New(1)
	handler := New(q, dedupe.New(), &metrics.Counters{}).Routes()
	const requests = 16
	statuses := make(chan int, requests)
	start := make(chan struct{})
	var workers sync.WaitGroup
	for range requests {
		workers.Add(1)
		go func() {
			defer workers.Done()
			<-start
			request := httptest.NewRequest(http.MethodPost, "/webhooks", bytes.NewBufferString(`{"id":"same","payload":"demo"}`))
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, request)
			statuses <- response.Code
		}()
	}
	close(start)
	workers.Wait()
	close(statuses)
	accepted := 0
	for status := range statuses {
		if status == http.StatusAccepted {
			accepted++
		} else if status != http.StatusOK {
			t.Fatalf("unexpected status=%d", status)
		}
	}
	if accepted != 1 || q.Stats().Depth != 1 {
		t.Fatalf("accepted=%d stats=%+v", accepted, q.Stats())
	}
}
