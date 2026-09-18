# Changelog

All notable changes follow Keep a Changelog and Semantic Versioning.

## [Unreleased]

### Fixed

- Preserve webhook retries after rejected queue admission and serialize admission with deduplication.
- Reject nonpositive retry budgets rather than returning success without work.
- Use Go 1.26.8 for standard-library security fixes across the module, CI and container build.

### Planned

- Broaden cross-platform verification and contributor-authored task coverage.

## [0.1.0] - 2026-07-23

### Added

- Synthetic production baseline with unit, integration, regression, negative,
  and performance coverage.
- Four reproducible engineering tasks with isolated candidate workspaces,
  public and held-out tests, golden patches, incorrect patches, and reports.
- Deterministic Python evaluator, CI workflows, Docker assets, and open-source
  governance documentation.
