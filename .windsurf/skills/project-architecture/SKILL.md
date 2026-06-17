---
name: project-architecture
description: "project-architecture"
# Content-Hash: blake3:eadb7b473a9b816ed8fbcdf29275e462b1b27ae3f87a49689c4e468fe5264ca9
# Source-Hash: blake3:42ec88d9fbc953ee4be083da31cedbaf02a3b623edff8de5828cb9d7c5768f3e
---

# Project Architecture

## Project Architecture

### Package Structure

go-utils is organized into specialized utility packages, each focused on a specific data type:

```
go-utils/
├── dateutils/      # Time and date manipulation utilities
├── maputils/       # Generic map operations
├── sliceutils/     # Generic slice operations (inspired by JS Array methods)
├── stringutils/    # String manipulation and formatting
├── structutils/    # Struct reflection and conversion utilities
└── urlutils/       # URL and query string utilities
```

### Design Principles

**1. Generics-First Approach**

All functions leverage Go 1.18+ generics for type safety:

```go
// Example: sliceutils.Filter signature
func Filter[T any](slice []T, predicate func(value T, index int, slice []T) bool) []T

// Example: maputils.Keys signature
func Keys[K comparable, V any](mapInstance map[K]V) []K
```

**2. Functional Programming Style**

Many functions accept callback functions with consistent signatures:

```go
// Slice callbacks receive: value, index, slice
sliceutils.Map(items, func(value T, index int, slice []T) R { ... })

// Map callbacks receive: key, value
maputils.ForEach(m, func(key K, value V) { ... })
```

**3. Zero-Cost Abstractions**

Functions are designed for performance with minimal overhead:
- Pre-allocated slices where sizes are known
- In-place operations where appropriate
- Efficient iteration patterns

### Adding New Utilities

When contributing a new utility function:

1. **Choose the correct package** based on the primary data type
2. **Add the function** to `<package>/<package>.go`
3. **Add comprehensive tests** to `<package>/<package>_test.go`
4. **Create documentation** at `docs/<package>/<function>.md` following existing patterns
5. **Update mkdocs.yml** to include the new documentation page

### Documentation Generation

Documentation is built with MkDocs and deployed automatically:

```bash
# Install MkDocs (if not already installed)
pip install mkdocs mkdocs-material

# Serve documentation locally
mkdocs serve
# Visit http://127.0.0.1:8000

# Build static documentation
mkdocs build
```

Documentation structure in `mkdocs.yml` mirrors the package structure for easy navigation.
