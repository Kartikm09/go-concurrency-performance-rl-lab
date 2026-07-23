# Architecture

`httpapi` owns request validation and overload responses; `queue` owns bounded FIFO
state; `worker` owns goroutine lifecycle; `retry` separates policy and timing; `dedupe`,
`store`, and `metrics` are concurrency-safe ports. Commands only compose these packages.

The evaluator creates a candidate-visible workspace and a separate internal held-out
workspace. See `DECISIONS.md` for the evidence boundary and reward model.
