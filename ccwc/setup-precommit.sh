#!/bin/bash
# Pre-commit setup script for Go project

set -e

echo "🚀 Setting up pre-commit hooks for Go project..."

# Check if pre-commit is installed
if ! command -v pre-commit &> /dev/null; then
    echo "❌ pre-commit is not installed."
    echo "Please install it using one of the following methods:"
    echo ""
    echo "Using pip:"
    echo "  pip install pre-commit"
    echo ""
    echo "Using brew (macOS):"
    echo "  brew install pre-commit"
    echo ""
    echo "Using conda:"
    echo "  conda install -c conda-forge pre-commit"
    echo ""
    exit 1
fi

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    echo "❌ Not in a git repository. Please run 'git init' first."
    exit 1
fi

# Install the git hook scripts
echo "📦 Installing pre-commit hooks..."
pre-commit install

# Run against all files to test
echo "🧪 Running pre-commit against all files..."
pre-commit run --all-files || {
    echo "⚠️  Some hooks failed. This is normal on first run."
    echo "   Files have been automatically fixed where possible."
    echo "   Please review the changes and commit them."
}

echo ""
echo "✅ Pre-commit setup complete!"
echo ""
echo "Usage:"
echo "  - Hooks will run automatically on 'git commit'"
echo "  - Run manually: 'pre-commit run --all-files'"
echo "  - Update hooks: 'pre-commit autoupdate'"
echo "  - Skip hooks (not recommended): 'git commit --no-verify'"
echo ""
echo "Available make targets for manual testing:"
echo "  - make check    # Run format, vet, and test"
echo "  - make ci       # Full CI pipeline"
echo "  - make fmt      # Format code"
echo "  - make test     # Run tests"
