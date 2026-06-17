---
description: Split tasks among subagents for parallel execution
name: parallelize
user_invocable: true
# Content-Hash: blake3:59bd0091e703c281080e66d1523d8c5252313e1d9bc0eca1ccc218bf27f549ee
# Source-Hash: blake3:42ec88d9fbc953ee4be083da31cedbaf02a3b623edff8de5828cb9d7c5768f3e
---

Split tasks among subagents. Analyze dependencies first — only parallelize truly independent work. Choose appropriate subagents, ensure each has full context, coordinate results.
