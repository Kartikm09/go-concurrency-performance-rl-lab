# Acceptance Criteria

- Stop returns within 500ms for idle workers.
- Concurrent submit and stop has no race or panic.
- Accepted jobs complete once.
- Repeated stop is safe.
