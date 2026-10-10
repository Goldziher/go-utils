---
type: Decision
title: Generic Type Constraints
x-ai-rulez:
  kind: rule
  id: generic-type-constraints
  metadata:
    priority: critical
---

# Generic Type Constraints

All utility functions must use appropriate generic type constraints (comparable, any, or custom interfaces). Follow existing patterns in sliceutils, maputils, and structutils packages. Always ensure type safety at compile time rather than runtime reflection where possible.
