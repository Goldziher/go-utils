---
type: Decision
title: Cicd Compliance
x-ai-rulez:
  kind: rule
  id: cicd-compliance
  metadata:
    priority: medium
---

# CI/CD Compliance

Ensure code passes all GitHub Actions workflows (ci.yaml, docs.yaml). CI runs tests, linting, and SonarCloud quality gates. Maintain quality gate status, zero vulnerabilities, and address bugs flagged by SonarCloud.
