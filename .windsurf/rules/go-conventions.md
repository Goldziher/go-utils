<!--
🤖 AI-RULEZ :: GENERATED FILE — DO NOT EDIT DIRECTLY
Project: go-utils
Generated: 2026-06-17 20:42:33
Source: .ai-rulez/config.toml
Target: .windsurf/rules/go-conventions.md
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
1. NEVER edit this file (.windsurf/rules/go-conventions.md) - it is auto-generated

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
Content-Hash: blake3:a1d7065c1b928113ebac756d0d3c69f4d60938612fdbf071684c2ab310f237eb
Source-Hash: blake3:42ec88d9fbc953ee4be083da31cedbaf02a3b623edff8de5828cb9d7c5768f3e
-->

# go-conventions

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
