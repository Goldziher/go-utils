---
type: Decision
title: Cicd And Dependencies
x-ai-rulez:
  kind: rule
  id: cicd-and-dependencies
  metadata:
    priority: medium
---

# CI/CD and Dependencies

All changes must pass GitHub Actions CI pipeline (.github/workflows/ci.yaml) including tests, linting, and security checks. Keep dependencies minimal and up-to-date using Dependabot. Verify compatibility with the latest stable Go version and document any version-specific requirements in go.mod.
