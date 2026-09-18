# Native reliability repair

Original commit: `4fed237c457eedc83a222ead72b375cebb0ad14a`.

The original Go 1.26.5 source passed `go test ./...`, `go vet ./...` and `go test -race ./...`
in credential-free macOS Seatbelt execution. An earlier full baseline build was interrupted
because simultaneous toolchain compilation overloaded the host; that interruption is an
infrastructure result, not a source failure. Tests use synthetic payloads only.

## Reproduced defects

**Rejected admission consumes its key (high).** Fill a one-slot queue, send webhook ID `retry`
and observe 429; dequeue the occupant and resend. Expected 202 with one new job; observed 200
with no job because the dedupe set was changed before enqueue. The HTTP regression failed
at `retry status=200`. `Set.Admit` now holds one lock across the short queue admission and ID
commit. It stores an ID only after success, so another request cannot acknowledge an
admission that failed. Queue admission performs no external I/O. The original `First` API
remains available. A 16-caller regression protects one admission under concurrent duplicates.

**Nonpositive retry budgets report success (medium).** `retry.Do` with budget 0 or -1 previously
returned nil without executing its operation. The regression observed `attempts=0 err=<nil>
calls=0`. It now returns `ErrInvalidAttempts` before any operation or wait. The existing
positive-budget retry test still protects attempt and backoff counts.

**Pinned Go standard-library vulnerabilities (high).** The pre-existing Security workflow
failed at the original commit on 2026-09-14
([Security run](https://github.com/Kartikm09/go-concurrency-performance-rl-lab/actions/runs/34831028891)): GO-2026-6090, GO-2026-6089 and GO-2026-5972 were
reachable in Go 1.26.5. The supported module minimum/toolchain, CI and Docker build now use
Go 1.26.8 on the same release branch. No task ID or historical baseline source was changed.

On 2026-09-18, the repaired tree passed `make setup && make verify-all && make fuzz &&
make benchmark` with Go 1.26.8 on macOS arm64 in credential-free Seatbelt isolation. This
includes formatting, vet, golangci-lint (zero issues), eight native test functions plus the
fuzz seed, race detection, six evaluator unit tests, and all sixteen existing evaluator
controls. The fuzz run completed 18,428 executions. A separate
`go run golang.org/x/vuln/cmd/govulncheck@v1.6.0 ./...` reported no vulnerabilities.

The benchmark run measured `BenchmarkJSONBatch` at 52,377 ns/op, 17,626 B/op and 502 allocs/op;
`BenchmarkQueueRoundTrip` measured 48.34 ns/op with zero allocations. These single-machine
observations are not before/after performance claims. An actual compiled API process also
returned health 200, first webhook 202 and duplicate webhook 200 over loopback HTTP; no
external service was contacted. The behavioral queue-drain and concurrent-admission
assertions run as native regression tests.

The ID set is in-memory and unbounded; restart persistence, ID expiry and distributed
admission remain outside this lab.
Historical benchmark and Docker reports are historical evidence, not fresh deployment claims.

An independent internal review checked dedupe-to-queue lock ordering, unsuccessful-admission
rollback, regression assertions and aligned runtime pins. It found no remaining blocker;
this is a bounded code review, not external contributor acceptance or production certification.

The unchanged native evaluator accepts four references and rejects its twelve existing
negative/scope controls. Three plausible controls fail behavioral/race assertions; task-002's
plausible control fails vet because `ErrOverloaded` is absent. Its current `incomplete_solution`
classification does not establish behavioral rejection. Stronger two-negative calibration was
implemented in the primary TypeScript track; this native audit does not relabel that gap.
