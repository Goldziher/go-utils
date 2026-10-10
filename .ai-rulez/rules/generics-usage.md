---
type: Decision
title: Generics Usage
x-ai-rulez:
  kind: rule
  id: generics-usage
  metadata:
    priority: critical
---

# Generics Usage

Leverage Go 1.18+ generics extensively for type-safe utility functions. Use constraint types (comparable, any) appropriately. All slice and map utilities must use generic type parameters to maintain API simplicity across different types.
