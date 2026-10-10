---
type: Decision
title: Error Handling And Panics
x-ai-rulez:
  kind: rule
  id: error-handling-and-panics
  metadata:
    priority: high
---

# Error Handling and Panics

Utility functions should not panic except for truly unrecoverable errors. Prefer returning error values or boolean success indicators. Document panic conditions explicitly in godoc comments if unavoidable (e.g., reflection-based operations in structutils).
