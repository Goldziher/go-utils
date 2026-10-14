---
type: Decision
title: "Migrated Linting Compliance"
x-ai-rulez:
  kind: rule
  id: legacy-linting-compliance
  metadata:
    priority: medium
---

All code must pass golangci-lint checks defined in .golangci.yml before committing. Run 'golangci-lint run' locally. Address all issues or add justified nolint directives with explanations. Pre-commit hooks enforce this automatically.
