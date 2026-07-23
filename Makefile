PYTHON ?= python3
TASK ?= task-001
PATCH ?= tasks/$(TASK)/golden/solution.patch
REPORT ?= reports/evaluations/$(TASK)/manual
GOLANGCI_LINT_VERSION ?= v2.12.2
.PHONY: setup format lint test race fuzz benchmark evaluator-test evaluate verify-evaluations security-scan verify-all
setup:
	go version
format:
	test -z "$$(gofmt -l .)"
lint:
	go vet ./...
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run
test:
	go test ./...
race:
	go test -race ./...
fuzz:
	go test -run=^$$ -fuzz=Fuzz -fuzztime=2s ./internal/httpapi
benchmark:
	go test -run=^$$ -bench=. -benchmem ./...
evaluator-test:
	PYTHONPATH=evaluator/src $(PYTHON) -m unittest discover -s evaluator/tests -v
evaluate:
	PYTHONPATH=evaluator/src $(PYTHON) -m rl_evaluator.cli evaluate --repo-root . --task $(TASK) --patch $(PATCH) --output $(REPORT)
verify-evaluations:
	$(PYTHON) scripts/verify_evaluations.py
security-scan:
	$(PYTHON) scripts/secret_scan.py
verify-all: format lint test race evaluator-test verify-evaluations security-scan
	$(PYTHON) scripts/validate_repository.py
