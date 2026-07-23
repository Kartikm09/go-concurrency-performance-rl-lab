set -euo pipefail
files=$(gofmt -l .)
if [ -n "$files" ]; then echo "$files"; exit 1; fi
