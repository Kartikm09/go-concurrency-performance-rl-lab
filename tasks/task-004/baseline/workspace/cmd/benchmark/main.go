package main

import (
	"candidate/task4/batch"
	"fmt"
)

func main() {
	events := make([]batch.Event, 100)
	for i := range events {
		events[i] = batch.Event{ID: fmt.Sprint(i), Payload: "synthetic"}
	}
	encoder := &batch.Encoder{}
	_, _ = encoder.Encode(events)
	fmt.Printf("{\"metric\":\"lock_acquisitions\",\"value\":%d}\n", encoder.LockCount())
}
