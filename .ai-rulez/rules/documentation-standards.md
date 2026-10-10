---
type: Decision
title: Documentation Standards
x-ai-rulez:
  kind: rule
  id: documentation-standards
  metadata:
    priority: medium
---

# Documentation Standards

Document all exported functions with godoc comments following Go conventions. Maintain corresponding markdown documentation in docs/ directory organized by package. Each function should have a dedicated .md file with usage examples for the MkDocs-generated documentation site.
