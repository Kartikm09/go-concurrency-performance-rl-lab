set -euo pipefail
go run ./cmd/benchmark > benchmark_result.json
cat benchmark_result.json
