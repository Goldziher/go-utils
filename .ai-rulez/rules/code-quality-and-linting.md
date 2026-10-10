---
type: Decision
title: Code Quality And Linting
x-ai-rulez:
  kind: rule
  id: code-quality-and-linting
  metadata:
    priority: high
---

# Code Quality and Linting

All code must pass golangci-lint checks as configured in .golangci.yml before commits. Use poly (poly.toml) to enforce quality gates; run `poly lint .` and `poly fmt --check .`. Maintain Go Report Card A+ rating and address all SonarCloud quality gate issues including maintainability, reliability, and security ratings.
