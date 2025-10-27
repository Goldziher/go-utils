# Go Utils

Functional programming utilities for Go 1.21+ designed to complement the standard library's `slices`, `maps`, and `cmp` packages.

## Installation

```bash
go get -u github.com/Goldziher/go-utils
```

## Design Philosophy

**Complementary to stdlib**: Use `slices.Index`, `slices.Contains`, `slices.Clone` for basic operations. Use this library for functional patterns (Map, Filter, Reduce), LINQ-style operations (GroupBy, Partition), and utilities not available in stdlib.

**Type-safe generics**: All functions leverage Go 1.21+ generics with appropriate constraints (`comparable`, `cmp.Ordered`, etc.) for compile-time type safety.

**Immutable by default**: Functions don't mutate inputs unless explicitly named (e.g., `Remove` creates new slices).

**100% test coverage**: Every function has comprehensive test coverage maintained by CI.

## Packages

### sliceutils
Functional operations and LINQ-style utilities for slices.

```go
import "github.com/Goldziher/go-utils/sliceutils"

// Functional patterns
numbers := []int{1, 2, 3, 4, 5}
doubled := sliceutils.Map(numbers, func(v, i int, s []int) int { return v * 2 })
evens := sliceutils.Filter(numbers, func(v, i int, s []int) bool { return v%2 == 0 })
sum := sliceutils.Reduce(numbers, func(acc, v, i int, s []int) int { return acc + v }, 0)

// LINQ-style operations
type User struct { Name string; Age int }
users := []User{{"Alice", 30}, {"Bob", 25}, {"Charlie", 30}}
byAge := sliceutils.GroupBy(users, func(u User) int { return u.Age })
adults, minors := sliceutils.Partition(users, func(u User) bool { return u.Age >= 18 })
```

### maputils
Map transformations and utilities.

```go
import "github.com/Goldziher/go-utils/maputils"

m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := maputils.Keys(m)           // []string{"a", "b", "c"}
values := maputils.Values(m)       // []int{1, 2, 3}
filtered := maputils.Filter(m, func(k string, v int) bool { return v > 1 })
```

### stringutils
String manipulation with type-safe conversion.

```go
import "github.com/Goldziher/go-utils/stringutils"

// Type-safe stringify with options
str := stringutils.Stringify(42, stringutils.Options{Base: 16})  // "2a"

// String manipulation
padded := stringutils.PadLeft("42", "0", 5)  // "00042"
capitalized := stringutils.Capitalize("hello")  // "Hello"
```

### structutils
Reflection-based struct utilities with tag support.

```go
import "github.com/Goldziher/go-utils/structutils"

type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

cfg := Config{Host: "localhost", Port: 8080}
m := structutils.ToMap(cfg, "json")  // map[string]any{"host": "localhost", "port": 8080}
```

### dateutils
Time and date utilities for business logic.

```go
import "github.com/Goldziher/go-utils/dateutils"

// Business day calculations
future := dateutils.AddBusinessDays(time.Now(), 5)

// Date overlap detection
overlaps := dateutils.Overlap(start1, end1, start2, end2)

// Age calculation
age := dateutils.Age(birthdate)
```

### urlutils
URL parsing and query string builders.

```go
import "github.com/Goldziher/go-utils/urlutils"

// Query string from map
params := map[string]any{"page": 1, "tags": []string{"go", "utils"}}
query := urlutils.QueryStringifyMap(params)  // "page=1&tags=go&tags=utils"

// Query string from struct with tags
type Query struct {
    Page int      `qs:"page"`
    Tags []string `qs:"tag"`
}
q := Query{Page: 1, Tags: []string{"go", "utils"}}
query := urlutils.QueryStringifyStruct(q, "qs")
```

### mathutils
Generic math operations with type constraints.

```go
import "github.com/Goldziher/go-utils/mathutils"

clamped := mathutils.Clamp(value, min, max)
inRange := mathutils.InRange(value, min, max)
isPrime := mathutils.IsPrime(17)  // true
gcd := mathutils.Gcd(48, 18)      // 6
```

## Links

- [GitHub Repository](https://github.com/Goldziher/go-utils)
- [Go Package Documentation](https://pkg.go.dev/github.com/Goldziher/go-utils)
- [Issue Tracker](https://github.com/Goldziher/go-utils/issues)
