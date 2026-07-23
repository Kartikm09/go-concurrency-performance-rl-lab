# Benchmark Evidence

Recorded on 2026-07-23 with Go 1.26.5 on an Apple M1 host. These local measurements
are reproducibility evidence, not cross-machine guarantees.

## Application Benchmarks

Command: `go test -run=^$ -bench=. -benchmem ./...`

```text
BenchmarkJSONBatch-8          35822   35895 ns/op   17630 B/op   502 allocs/op
BenchmarkQueueRoundTrip-8  39271994      35.67 ns/op       0 B/op     0 allocs/op
```

## Task 004 Before And After

| Workspace | `lock_acquisitions` | Result |
| --- | ---: | --- |
| Published baseline | 100 | Reproduced with `scripts/run_benchmark.sh` |
| Golden patch | 1 | Accepted; configured threshold is at most 1 |

The deterministic lock counter avoids treating noisy wall-clock timing as the reward.
The accepted report in `reports/evaluations/task-004/solution/result.json` also passed
race-enabled public, held-out, and regression tests.
