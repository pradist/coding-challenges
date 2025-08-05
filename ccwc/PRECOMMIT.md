# Pre-commit Setup Guide

This project uses [pre-commit](https://pre-commit.com/) to automatically run code quality checks before each commit.

## Installation

### 1. Install pre-commit

**Using pip:**

```bash
pip install pre-commit
```

**Using brew (macOS):**

```bash
brew install pre-commit
```

**Using conda:**

```bash
conda install -c conda-forge pre-commit
```

### 2. Set up the hooks

**Automated setup:**

```bash
make setup-precommit
# or
./setup-precommit.sh
```

**Manual setup:**

```bash
pre-commit install
```

## Usage

### Automatic (Recommended)

Once installed, hooks will run automatically on every `git commit`.

### Manual

```bash
# Run all hooks on all files
make pre-commit-run
# or
pre-commit run --all-files

# Run all hooks on staged files only
pre-commit run

# Update hook versions
make pre-commit-update
# or
pre-commit autoupdate
```

## What hooks are included?

### General hooks

- **trailing-whitespace**: Remove trailing whitespace
- **end-of-file-fixer**: Ensure files end with newline
- **check-yaml**: Validate YAML files
- **check-json**: Validate JSON files
- **check-added-large-files**: Prevent large files from being committed
- **check-merge-conflict**: Check for merge conflict markers

### Go-specific hooks

- **go-fmt**: Format Go code using `gofmt`
- **go-vet**: Run `go vet` for static analysis
- **go-imports**: Organize imports
- **go-mod-tidy**: Tidy module dependencies
- **go-unit-tests**: Run unit tests
- **golangci-lint**: Comprehensive Go linting
- **go-cyclo**: Check cyclomatic complexity
- **make-check**: Run our Makefile check target

## Skipping hooks (Not recommended)

If you need to skip hooks for a specific commit:

```bash
git commit --no-verify
```

## Troubleshooting

### Hook fails on first run

This is normal! Pre-commit may fix files automatically. Review the changes and commit them:

```bash
git add .
git commit -m "Fix pre-commit issues"
```

### Updating hooks

```bash
pre-commit autoupdate
```

### Re-installing hooks

```bash
pre-commit uninstall
pre-commit install
```

## Integration with Make

The pre-commit hooks integrate with your existing Makefile:

- `make check` runs format, vet, and test
- `make ci` runs the full CI pipeline
- Hooks use these same commands for consistency

## Configuration

The configuration is in `.pre-commit-config.yaml`. You can:

- Add new hooks
- Modify existing hook arguments
- Skip specific hooks for certain files
- Update hook versions

Example customization:

```yaml
- id: go-cyclo
  args: [-over=10]  # Lower complexity threshold
```
