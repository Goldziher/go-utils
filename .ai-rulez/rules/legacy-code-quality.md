---
type: Decision
title: "Migrated Code Quality"
x-ai-rulez:
  kind: rule
  id: legacy-code-quality
  metadata:
    priority: high
---

All code must pass golangci-lint checks using the configuration in .golangci.yml. Pre-commit hooks (.pre-commit-config.yaml) enforce linting standards. Address all linter warnings before committing.
