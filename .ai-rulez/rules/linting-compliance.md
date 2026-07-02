---
priority: medium
---

# Linting Compliance

All code must pass golangci-lint checks defined in .golangci.yml before committing. Run 'golangci-lint run' locally. Address all issues or add justified nolint directives with explanations. poly enforces this automatically via the shared reusable validate workflow in CI.
