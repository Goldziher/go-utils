<!--
🤖 AI-RULEZ :: GENERATED FILE — DO NOT EDIT DIRECTLY
Project: go-utils
Generated: 2026-06-17 20:42:33
Source: .ai-rulez/config.toml
Target: .windsurf/rules/minimal-changes.md
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
1. NEVER edit this file (.windsurf/rules/minimal-changes.md) - it is auto-generated

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
Content-Hash: blake3:3d13df876cf9c9c8932531d906876e5c3e62f7413fdfb016a206c86d90042886
Source-Hash: blake3:42ec88d9fbc953ee4be083da31cedbaf02a3b623edff8de5828cb9d7c5768f3e
-->

# minimal-changes

**Priority:** high

Make the smallest change that achieves the goal. Avoid unnecessary refactoring, reformatting, or scope creep. Don't fix what isn't broken.
