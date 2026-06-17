<!--
🤖 AI-RULEZ :: GENERATED FILE — DO NOT EDIT DIRECTLY
Project: go-utils
Generated: 2026-06-17 20:42:33
Source: .ai-rulez/config.toml
Target: .github/copilot-instructions.md
Content: rules=53, sections=0, agents=0

WHAT IS AI-RULEZ
AI-Rulez is a directory-based AI governance tool. All configuration lives in
the .ai-rulez/ directory. This file is auto-generated from source files.

.AI-RULEZ FOLDER ORGANIZATION
Root content (always included):
  .ai-rulez/config.toml    Main configuration (presets, profiles)
  .ai-rulez/rules/         Mandatory rules for AI assistants
  .ai-rulez/context/       Reference documentation
  .ai-rulez/skills/        Specialized AI prompts
  .ai-rulez/agents/        Agent definitions

Domain content (profile-specific):
  .ai-rulez/domains/{name}/rules/    Domain-specific rules
  .ai-rulez/domains/{name}/context/  Domain-specific documentation
  .ai-rulez/domains/{name}/skills/   Domain-specific AI prompts

Profiles in config.toml control which domains are included.

INSTRUCTIONS FOR AI AGENTS
1. NEVER edit this file (.github/copilot-instructions.md) - it is auto-generated

2. ALWAYS edit files in .ai-rulez/ instead:
   - Add/modify rules: .ai-rulez/rules/*.md
   - Add/modify context: .ai-rulez/context/*.md
   - Update config: .ai-rulez/config.toml
   - Domain-specific: .ai-rulez/domains/{name}/rules/*.md

3. PREFER using the MCP Server (if available):
   Command: npx -y ai-rulez@latest mcp
   Provides safe CRUD tools for reading and modifying .ai-rulez/ content

4. After making changes: ai-rulez generate

5. Complete workflow:
   a. Edit source files in .ai-rulez/
   b. Run: ai-rulez generate
   c. Commit both .ai-rulez/ and generated files

Documentation: https://github.com/Goldziher/ai-rulez
Content-Hash: blake3:b879eaa5ae0cff72d6bea2f87ba0a568ce5de2e200fde5af167b0a83778f4365
Source-Hash: blake3:42ec88d9fbc953ee4be083da31cedbaf02a3b623edff8de5828cb9d7c5768f3e
-->

# go-utils

A comprehensive Go utility library providing collection manipulation functions inspired by JavaScript and Python, leveraging Go 1.18+ generics for type-safe operations on slices, maps, structs, dates, strings, and URLs with utilities like filter, map, reduce, merge, flatten, and query string builders.

## Rules

### agent-workflow

**Priority:** high

Prefer subagents for non-trivial work — implementation, research, file exploration. Parallelize aggressively — launch independent subagents in a single message. Always critically review subagent output — check actual file changes, verify correctness, fix issues before reporting done. Never trust subagent summaries at face value; the summary describes intent, not necessarily what happened. Work in iterations: delegate → critically review → fix → verify. Run tests after every change — never assume code works without verification.

### anti-patterns

**Priority:** high

No magic numbers — use named constants. No global state — use dependency injection. No inheritance for code reuse — prefer composition. No bare exception handlers — catch specific types. No mocking internal services — use real objects for integration tests. No blocking I/O in async code paths — keep async paths fully async.

### atomic-commits

**Priority:** high

Each commit represents one logical change. Don't mix unrelated changes. Use conventional commits format (`feat:`, `fix:`, `chore:`, `refactor:`, `docs:`, `test:`). Keep commits small and focused for easier review and bisection.

### avoid-duplication

**Priority:** medium

Extract shared logic after the third repetition, not before. Three similar lines of code are better than a premature abstraction. When extracting, ensure the shared code has a single reason to change — if two callers would evolve the logic differently, keep them separate. Premature abstraction creates worse coupling than duplication.

### batch-operations

**Priority:** medium

Group related file reads and writes into single operations. Combine independent tool calls in parallel rather than sequentially. When making multiple edits to the same file, batch them into one edit operation. Prefer multi-file search tools over individual file reads when exploring.

### branch-hygiene

**Priority:** medium

Use descriptive branch names. Keep branches short-lived. Delete merged branches. Rebase or merge from main regularly to avoid drift.

### cicd-and-dependencies

**Priority:** medium

All changes must pass GitHub Actions CI pipeline (.github/workflows/ci.yaml) including tests, linting, and security checks. Keep dependencies minimal and up-to-date using Dependabot. Verify compatibility with the latest stable Go version and document any version-specific requirements in go.mod.

### cicd-compliance

**Priority:** medium

Ensure code passes all GitHub Actions workflows (ci.yaml, docs.yaml). CI runs tests, linting, and SonarCloud quality gates. Maintain quality gate status, zero vulnerabilities, and address bugs flagged by SonarCloud.

### code-quality

**Priority:** high

All code must pass golangci-lint checks using the configuration in .golangci.yml. Pre-commit hooks (.pre-commit-config.yaml) enforce linting standards. Address all linter warnings before committing.

### code-quality-and-linting

**Priority:** high

All code must pass golangci-lint checks as configured in .golangci.yml before commits. Use pre-commit hooks (.pre-commit-config.yaml) to enforce quality gates. Maintain Go Report Card A+ rating and address all SonarCloud quality gate issues including maintainability, reliability, and security ratings.

### commit-messages

**Priority:** high

Use conventional commits: `feat: add user auth`, `fix: handle null input`, `chore: update deps`, `refactor: extract parser`, `docs: add API guide`, `test: cover edge case`. First line under 72 chars, imperative mood. Body explains *why*, not *what*. Add scope when useful: `feat(api): add pagination`.

### communication-style

**Priority:** critical

Be concise and precise — no fluff, no emojis, no unnecessary checklists. PR descriptions: state what changed and why in 1-3 sentences, not bullet-point essays. Issue comments: answer the question directly. Code review: point out the problem and suggest the fix, skip praise and filler. Commit messages: imperative mood, under 72 chars, body explains why not what. Never pad output to appear thorough — brevity is clarity.

### complexity-limits

**Priority:** medium

Enforce concrete limits: max 20 cyclomatic complexity per function, max 4 levels of nesting depth, max 50 lines per function. Use early returns to flatten conditionals. Break complex functions into well-named helpers that each do one thing.

### context-preservation

**Priority:** medium

Record key findings (file paths, function signatures, patterns discovered) before they scroll out of context. Summarize investigation results before acting on them. When working on multi-step tasks, note intermediate decisions and their rationale to avoid re-deriving them later.

### cross-package-consistency

**Priority:** medium

Maintain naming and behavior consistency across utility packages. Similar operations should use similar names (e.g., Filter, Map, ForEach exist in both sliceutils and maputils). Function signatures should follow established patterns for predicates, mappers, and reducers.

### dead-code

**Priority:** low

Remove dead code instead of commenting it out. Version control preserves history. Commented-out code creates confusion and maintenance burden.

### dependencies

**Priority:** medium

Minimize external dependencies to maintain utility library simplicity. Use Go standard library whenever possible. Dependencies are managed via go.mod and monitored by Dependabot. Justify any new dependency additions.

### dependency-awareness

**Priority:** high

Audit dependencies before adding them. Prefer well-maintained, widely-used packages with active maintenance. Pin versions and commit lock files. Use language-specific audit tools in CI:

- Rust: `cargo audit`, `cargo deny` (license + advisory policies)
- Python: `pip-audit`, `bandit` (SAST)
- JavaScript/TypeScript: `npm audit`, `pnpm audit`
- Go: `govulncheck`
- Ruby: `bundler-audit`
- PHP: `composer audit`
- Java: OWASP `dependency-check` Maven/Gradle plugin
- C#: `dotnet list package --vulnerable`
- Elixir: `mix_audit`
Zero tolerance for critical/high CVEs. Automate dependency update PRs where possible.

### documentation

**Priority:** high

Provide godoc comments for all exported functions, types, and constants. Maintain corresponding markdown documentation in docs/ directory organized by package. Documentation is published via MkDocs (mkdocs.yml) to GitHub Pages.

### documentation-standards

**Priority:** medium

Document all exported functions with godoc comments following Go conventions. Maintain corresponding markdown documentation in docs/ directory organized by package. Each function should have a dedicated .md file with usage examples for the MkDocs-generated documentation site.

### error-handling

**Priority:** high

Always wrap errors with context describing what operation failed. Never swallow errors silently — either handle, propagate, or log them. Use language-idiomatic patterns: `Result<T, E>` in Rust, `if err != nil` with `fmt.Errorf("doing X: %w", err)` in Go, typed exceptions in Python/Java. Fail fast on unrecoverable errors.

### error-handling-and-panics

**Priority:** high

Utility functions should not panic except for truly unrecoverable errors. Prefer returning error values or boolean success indicators. Document panic conditions explicitly in godoc comments if unavoidable (e.g., reflection-based operations in structutils).

### explain-reasoning

**Priority:** medium

Briefly explain your reasoning for non-obvious decisions. State trade-offs when multiple approaches exist. Be transparent about uncertainty.

### generic-type-constraints

**Priority:** critical

All utility functions must use appropriate generic type constraints (comparable, any, or custom interfaces). Follow existing patterns in sliceutils, maputils, and structutils packages. Always ensure type safety at compile time rather than runtime reflection where possible.

### generics-usage

**Priority:** critical

Leverage Go 1.18+ generics extensively for type-safe utility functions. Use constraint types (comparable, any) appropriately. All slice and map utilities must use generic type parameters to maintain API simplicity across different types.

### go-conventions

**Priority:** high

- Follow Effective Go and Go Code Review Comments guidelines.
- Handle every error return. Wrap errors with context: `fmt.Errorf("operation failed: %w", err)`.
- Linting: `golangci-lint` with strict config (enable `govet`, `staticcheck`, `errcheck`, `gosec`, `gocritic`). Format with `gofmt`/`goimports`.
- Security: `govulncheck` for CVE scanning, `gosec` for SAST. Run both in CI.
- Testing: table-driven tests with `t.Run()`. Use `t.Parallel()` where safe. Use `testify` for assertions. Coverage with `go test -coverprofile`.
- Naming: use short, descriptive names. Receivers are 1-2 letters. Exported names are descriptive.
- Prefer composition over inheritance. Use interfaces for abstraction (accept interfaces, return structs).
- Keep packages small and focused. Avoid package-level state and `init()` functions.
- Use `context.Context` as first parameter for cancelable operations. Never store contexts in structs.
- Error types: use `errors.Is()`/`errors.As()` for comparison. Define sentinel errors with `errors.New()`.
- Modules: use Go modules. Commit `go.sum`. Use semantic version tags. Prefer stdlib over third-party when reasonable.
- Concurrency: prefer channels over mutexes. Use `sync.WaitGroup` for fan-out. Guard shared state.
- Benchmarking: use `testing.B` benchmarks. Profile with `go tool pprof`. Use `benchstat` for comparison.

### immutability-and-side-effects

**Priority:** medium

Follow functional programming patterns used in sliceutils and maputils. Functions should not mutate input parameters unless explicitly named (e.g., 'Copy' creates new instances, 'ForEach' is readonly). When mutation is necessary, document it clearly in function comments.

### incremental-approach

**Priority:** medium

Start with the smallest viable change, verify it works, then extend. Avoid generating large blocks of speculative code. Build iteratively: implement one piece, test, then move to the next. When uncertain about an approach, prototype the critical part first before committing to the full implementation.

### input-validation

**Priority:** high

Validate and sanitize all external input at system boundaries. Use allowlists over denylists. Validate types, ranges, and formats. Never trust user input.

### least-privilege

**Priority:** medium

Request only necessary permissions. Minimize file system access, network access, and API scopes. Run processes with minimal required privileges.

### linting-compliance

**Priority:** medium

All code must pass golangci-lint checks defined in .golangci.yml before committing. Run 'golangci-lint run' locally. Address all issues or add justified nolint directives with explanations. Pre-commit hooks enforce this automatically.

### meaningful-assertions

**Priority:** medium

Assert exact expected values, not just truthiness (`assert result == 42`, not `assert result`). Use snapshot testing for complex structured output. Consider property-based testing for functions with wide input ranges. Include descriptive failure messages. Always test error paths and edge cases, not just the happy path.

### minimal-changes

**Priority:** high

Make the smallest change that achieves the goal. Avoid unnecessary refactoring, reformatting, or scope creep. Don't fix what isn't broken.

### no-ai-signatures

**Priority:** critical

Never add AI attribution to commits (no Co-Authored-By AI lines, no "Generated by AI/Claude/GPT"). Never add AI attribution to PR titles or descriptions. Never add AI-generated comments or watermarks in code.

### output-awareness

**Priority:** medium

Limit explanations to 1-3 sentences unless asked for detail. Use code blocks for code, not prose. Omit unchanged code when showing diffs — use comments like `// ... existing code ...` to indicate skipped sections. Never repeat information already visible in context. Prefer short, direct answers over comprehensive walkthroughs.

### package-documentation

**Priority:** high

All exported functions, types, and constants must have godoc comments. Documentation should include usage examples and edge case behavior. Each package utility function should have a corresponding markdown file in the docs/ directory matching the existing structure (e.g., docs/sliceutils/functionname.md).

### package-organization

**Priority:** high

Organize code into focused utility packages (dateutils, sliceutils, maputils, stringutils, structutils, urlutils). Each package should contain a single Go file with implementation and a corresponding _test.go file. Keep packages flat - no nested subpackages.

### read-before-write

**Priority:** critical

Read and understand existing files before editing them. Understand the codebase conventions, patterns, and architecture before making changes. Check imports, naming styles, and project structure to ensure new code fits the existing codebase.

### readability-first

**Priority:** high

Max 120 character line width. Prefer explicit code over clever tricks — if it needs a comment to explain what it does, rewrite it. No abbreviations in public API names (`context` not `ctx` in public signatures, `repository` not `repo`). Keep functions short and focused on a single responsibility.

### safe-git-operations

**Priority:** critical

Never force-push to shared branches. Always pull before pushing. Use `--force-with-lease` instead of `--force` when necessary. Confirm destructive operations with the user.

### secrets-handling

**Priority:** critical

Never hardcode secrets, API keys, tokens, or passwords. Use environment variables or secret management systems. Never log or expose sensitive values. Reject commits containing secrets.

### systematic-debugging

**Priority:** high

Never guess at bugs. Trace the root cause backward through the call stack to find the original trigger. Analyze patterns — is this a one-off or systemic? Form a hypothesis and verify it before implementing a fix. No shotgun debugging, no random changes hoping something works.

### task-runner

**Priority:** high

Prefer `task` commands over raw build/test/lint commands when a Taskfile.yaml exists. Task runners provide consistent, documented workflows. Use `task --list` to discover available tasks. Always check for a Taskfile before running manual commands. Standard task names: setup, build, test, lint, format, bench — prefer these conventions. Lock files always committed for reproducible builds.

### tdd-workflow

**Priority:** high

Write tests before writing code, update tests when modifying behavior. When fixing bugs, write a failing test first — RED (failing test) → GREEN (minimal code to pass) → REFACTOR. Wrote production code before the test? Delete it, start over — no exceptions, don't keep as reference. Integration tests for API surfaces, unit tests for business logic, property tests for edge-case-heavy code. Run the full test suite before committing — never push untested code.

### test-alongside-code

**Priority:** high

Write tests when writing code, update tests when modifying behavior. When fixing bugs, write a failing test first (TDD). Use integration tests for the public API surface and unit tests for complex internal logic. Run the full test suite before committing.

### test-independence

**Priority:** high

Tests must be independent and idempotent — runnable in any order, in parallel. No shared mutable state between tests. Use factories or fixtures for setup. Clean up created resources (files, DB rows, env vars) after each test. Never rely on test execution order.

### test-naming

**Priority:** medium

Name tests to describe behavior: `should_return_error_when_input_is_empty`, `test_parse_handles_nested_objects`. Use `describe`/`it` blocks for grouping in languages that support them. Follow `given_when_then` or `should_when` patterns. Test names are specifications — a reader should understand the expected behavior without reading the test body.

### testing-and-coverage

**Priority:** critical

Write comprehensive unit tests for all utility functions in parallel test files (*_test.go). Maintain high coverage standards as monitored by SonarCloud. Each utility package (sliceutils, maputils, dateutils, stringutils, structutils, urlutils) must have corresponding test coverage with table-driven tests for multiple input scenarios.

### testing-anti-patterns

**Priority:** high

Do not test mock behavior instead of real behavior. Do not add test-only methods to production code. Do not mock what you don't own — wrap it and test the wrapper. Do not test implementation details — test observable behavior. Do not write tests that pass when the code is broken. If a test never fails, it's not testing anything.

### testing-coverage

**Priority:** critical

Every exported function must have corresponding test cases in *_test.go files. Follow the table-driven test pattern used throughout the codebase. Run 'go test ./...' to verify all tests pass before committing. Maintain high coverage as monitored by SonarCloud.

### testing-standards

**Priority:** critical

Write comprehensive table-driven tests in *_test.go files for every exported function. Maintain high code coverage as monitored by SonarCloud. Run tests with 'go test ./...' and ensure all tests pass before committing.

### verification-before-completion

**Priority:** critical

Never claim success without fresh verification. Run the test and see it pass. Check the file exists. Verify the build succeeds. Evidence before assertions — always. If you can't verify, say so explicitly rather than claiming success.

### verify-before-acting

**Priority:** critical

Verify assumptions before taking action. Check current state (branch, working directory, running processes) before making changes. Confirm file existence before editing. Test that build passes before committing. Never assume — confirm.

## Context

### owasp-quick-reference

1. **Broken Access Control** — enforce authorization checks on every request, deny by default.
2. **Cryptographic Failures** — use strong standard algorithms, never roll your own crypto.
3. **Injection** — parameterize all queries, sanitize and validate all inputs.
4. **Insecure Design** — threat model early, validate business logic at every layer.
5. **Security Misconfiguration** — harden defaults, disable unnecessary features and endpoints.
6. **Vulnerable Components** — keep dependencies updated, audit regularly with language-specific tools.
7. **Authentication Failures** — require MFA, enforce strong passwords, implement rate limiting.
8. **Data Integrity Failures** — verify software updates, use signed artifacts and checksums.
9. **Logging Failures** — log all security events with context, protect log data from tampering.
10. **SSRF** — validate and allowlist URLs, restrict outbound network requests.
