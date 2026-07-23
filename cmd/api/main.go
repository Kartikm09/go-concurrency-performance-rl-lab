package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/dedupe"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/httpapi"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/metrics"
	"github.com/Kartikm09/go-concurrency-performance-rl-lab/internal/queue"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8083"
	}
	handler := httpapi.New(queue.New(128), dedupe.New(), &metrics.Counters{})
	log.Printf("webhook API listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler.Routes()))
}
