---
type: Decision
title: Testing And Coverage
x-ai-rulez:
  kind: rule
  id: testing-and-coverage
  metadata:
    priority: critical
---

# Testing and Coverage

Write comprehensive unit tests for all utility functions in parallel test files (*_test.go). Maintain high coverage standards as monitored by SonarCloud. Each utility package (sliceutils, maputils, dateutils, stringutils, structutils, urlutils) must have corresponding test coverage with table-driven tests for multiple input scenarios.
