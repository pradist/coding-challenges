# Makefile Usage Guide

This Makefile provides common Go testing and development commands.

## Quick Start

```bash
# Show all available commands
make help

# Run basic tests
make test

# Run tests with coverage
make coverage-report

# Run full check (format + vet + test)
make check
```

## Common Commands

### Testing

- `make test` - Run all tests
- `make test-verbose` - Run tests with verbose output
- `make test-race` - Run tests with race detector
- `make coverage` - Run tests with coverage percentage
- `make coverage-report` - Generate detailed coverage report
- `make coverage-html` - Generate HTML coverage report
- `make coverage-total` - Show only total coverage percentage

### Code Quality

- `make fmt` - Format Go code
- `make vet` - Run go vet
- `make check` - Run format, vet, and test together

### Build & Dependencies

- `make build` - Build the project
- `make mod-tidy` - Tidy module dependencies
- `make clean` - Clean build artifacts

### CI/CD

- `make ci` - Full CI pipeline (tidy, format, vet, test, coverage)
- `make test-all` - Run all types of tests with coverage

## Examples

```bash
# Development workflow
make check              # Quick check before committing
make coverage-html      # Generate coverage report to review
make clean              # Clean up artifacts

# CI/CD pipeline
make ci                 # Full pipeline for continuous integration

# Coverage analysis
make coverage-total     # Quick coverage percentage
make coverage-html      # Detailed HTML report
```

## Requirements

- Go 1.16+
- Optional: `golangci-lint` for the `make lint` command
- Optional: `entr` for the `make watch` command
