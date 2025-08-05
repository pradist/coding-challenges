# Go Makefile for common testing commands

# Variables
GO_MODULE := $(shell go list -m)
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Test targets
.PHONY: test
test: ## Run all tests
	go test ./...

.PHONY: test-verbose
test-verbose: ## Run all tests with verbose output
	go test -v ./...

.PHONY: test-short
test-short: ## Run tests with short flag (skip long-running tests)
	go test -short ./...

.PHONY: test-race
test-race: ## Run tests with race detector
	go test -race ./...

.PHONY: test-count
test-count: ## Run tests multiple times (default: 3)
	go test -count=3 ./...

# Coverage targets
.PHONY: coverage
coverage: ## Run tests with coverage
	go test -cover ./...

.PHONY: coverage-report
coverage-report: ## Generate coverage report
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -func=$(COVERAGE_FILE)

.PHONY: coverage-html
coverage-html: ## Generate HTML coverage report
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"

.PHONY: coverage-total
coverage-total: ## Show total coverage percentage
	@go test -coverprofile=$(COVERAGE_FILE) ./... > /dev/null 2>&1
	@go tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print "Total Coverage: " $$3}'

# Benchmark targets
.PHONY: bench
bench: ## Run benchmarks
	go test -bench=. ./...

.PHONY: bench-mem
bench-mem: ## Run benchmarks with memory allocation stats
	go test -bench=. -benchmem ./...

.PHONY: bench-cpu
bench-cpu: ## Run benchmarks with CPU profiling
	go test -bench=. -cpuprofile=cpu.prof ./...

# Lint and format targets
.PHONY: fmt
fmt: ## Format Go code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: lint
lint: ## Run golangci-lint (requires golangci-lint to be installed)
	golangci-lint run

# Build targets
.PHONY: build
build: ## Build the project
	go build ./...

.PHONY: clean
clean: ## Clean build artifacts and coverage files
	go clean ./...
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML) cpu.prof mem.prof

# Module targets
.PHONY: mod-tidy
mod-tidy: ## Run go mod tidy
	go mod tidy

.PHONY: mod-verify
mod-verify: ## Verify module dependencies
	go mod verify

.PHONY: mod-download
mod-download: ## Download module dependencies
	go mod download

# Composite targets
.PHONY: check
check: fmt vet test ## Run format, vet, and test

.PHONY: ci
ci: mod-tidy fmt vet test coverage-report ## Full CI pipeline

.PHONY: test-all
test-all: test test-race coverage-html ## Run all types of tests with coverage

# Watch target (requires entr or similar tool)
.PHONY: watch
watch: ## Watch for changes and run tests (requires 'entr' command)
	find . -name '*.go' | entr -c make test
