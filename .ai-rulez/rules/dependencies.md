---
type: Decision
title: Dependencies
x-ai-rulez:
  kind: rule
  id: dependencies
  metadata:
    priority: medium
---

# Dependencies

Minimize external dependencies to maintain utility library simplicity. Use Go standard library whenever possible. Dependencies are managed via go.mod and monitored by Dependabot. Justify any new dependency additions.
