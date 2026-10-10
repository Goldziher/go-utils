---
type: Decision
title: Testing Coverage
x-ai-rulez:
  kind: rule
  id: testing-coverage
  metadata:
    priority: critical
---

# Testing Coverage

Every exported function must have corresponding test cases in *_test.go files. Follow the table-driven test pattern used throughout the codebase. Run 'go test ./...' to verify all tests pass before committing. Maintain high coverage as monitored by SonarCloud.
