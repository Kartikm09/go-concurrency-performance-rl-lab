# Retry policy separation

Separate retry classification, backoff timing, and transport execution while preserving behavior.

Submit a unified diff against the candidate workspace. The patch is evaluated
with public tests first and held-out correctness and regression tests only in an
internal copy. Include `CANDIDATE_NOTES.md` with the invariant, compatibility
impact, and commands used.
