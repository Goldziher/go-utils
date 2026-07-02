---
priority: high
description: "development-workflow"
---

# Development Workflow

## Development Workflow

### Prerequisites
- Go 1.18+ (required for generics support)
- poly (for linting and formatting)
- golangci-lint (for linting)

### Initial Setup

```bash
# Clone and navigate to the repository
cd /Users/naamanhirschfeld/workspace/go-utils

# Install dependencies
go mod download

# poly runs in CI via the shared reusable validate workflow.
# Verify locally before committing:
poly fmt --check . && poly lint .
```

### Running Tests

```bash
# Run all tests with coverage
go test ./... -v -cover

# Run tests for specific package
go test ./sliceutils -v
go test ./maputils -v
go test ./dateutils -v
go test ./stringutils -v
go test ./structutils -v
go test ./urlutils -v
```

**Important**: 100% test coverage is mandatory for all contributions. Each utility package has a corresponding `_test.go` file that must be updated when adding new functions.

### Linting

```bash
# Run golangci-lint (configuration in .golangci.yml)
golangci-lint run

# Auto-fix issues where possible
golangci-lint run --fix
```

### Commit Convention

This project enforces [Conventional Commits](https://www.conventionalcommits.org/) via poly:

```bash
# Valid commit message examples
git commit -m "feat: add new Map function to sliceutils"
git commit -m "fix: correct edge case in dateutils.Overlap"
git commit -m "docs: update contributing guide"
git commit -m "test: add coverage for stringutils.Capitalize"
```

### CI Pipeline

The project uses GitHub Actions (`.github/workflows/ci.yaml` and `.github/workflows/docs.yaml`) to:
- Run tests across multiple Go versions
- Check code coverage
- Run linters
- Build and deploy documentation to GitHub Pages
- Analyze code quality with SonarCloud
