---
type: Decision
title: Package Documentation
x-ai-rulez:
  kind: rule
  id: package-documentation
  metadata:
    priority: high
---

# Package Documentation

All exported functions, types, and constants must have godoc comments. Documentation should include usage examples and edge case behavior. Each package utility function should have a corresponding markdown file in the docs/ directory matching the existing structure (e.g., docs/sliceutils/functionname.md).
