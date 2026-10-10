---
type: Decision
title: Testing Standards
x-ai-rulez:
  kind: rule
  id: testing-standards
  metadata:
    priority: critical
---

# Testing Standards

Write comprehensive table-driven tests in *_test.go files for every exported function. Maintain high code coverage as monitored by SonarCloud. Run tests with 'go test ./...' and ensure all tests pass before committing.
