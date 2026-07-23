# Acceptance Criteria

- Transport, classifier, and backoff are independent interfaces.
- Attempt count and terminal errors match the baseline.
- Backoff receives deterministic attempt numbers.
- No wall-clock sleeps in tests.
