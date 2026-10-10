---
type: Decision
title: Immutability And Side Effects
x-ai-rulez:
  kind: rule
  id: immutability-and-side-effects
  metadata:
    priority: medium
---

# Immutability and Side Effects

Follow functional programming patterns used in sliceutils and maputils. Functions should not mutate input parameters unless explicitly named (e.g., 'Copy' creates new instances, 'ForEach' is readonly). When mutation is necessary, document it clearly in function comments.
