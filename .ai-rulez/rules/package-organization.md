---
type: Decision
title: Package Organization
x-ai-rulez:
  kind: rule
  id: package-organization
  metadata:
    priority: high
---

# Package Organization

Organize code into focused utility packages (dateutils, sliceutils, maputils, stringutils, structutils, urlutils). Each package should contain a single Go file with implementation and a corresponding _test.go file. Keep packages flat - no nested subpackages.
