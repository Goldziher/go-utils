---
type: Decision
title: "Migrated Code Quality and Linting"
x-ai-rulez:
  kind: rule
  id: legacy-code-quality-and-linting
  metadata:
    priority: high
---

All code must pass golangci-lint checks as configured in .golangci.yml before commits. Use Poly hooks (.pre-commit-config.yaml) to enforce quality gates. Maintain Go Report Card A+ rating and address all SonarCloud quality gate issues including maintainability, reliability, and security ratings.
