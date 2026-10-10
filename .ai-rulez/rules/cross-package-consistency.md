---
type: Decision
title: Cross Package Consistency
x-ai-rulez:
  kind: rule
  id: cross-package-consistency
  metadata:
    priority: medium
---

# Cross-Package Consistency

Maintain naming and behavior consistency across utility packages. Similar operations should use similar names (e.g., Filter, Map, ForEach exist in both sliceutils and maputils). Function signatures should follow established patterns for predicates, mappers, and reducers.
