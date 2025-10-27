# Go Utils

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/Goldziher/go-utils)](https://goreportcard.com/report/github.com/Goldziher/go-utils)
[![Go Reference](https://pkg.go.dev/badge/github.com/Goldziher/go-utils.svg)](https://pkg.go.dev/github.com/Goldziher/go-utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

</div>

Functional programming utilities for **Go 1.21+** designed to complement stdlib's `slices`, `maps`, and `cmp` packages.

## Why Go Utils?

**Complementary to stdlib**: Use `slices.Index`, `slices.Contains`, `slices.Clone` for basic operations. Use go-utils for functional patterns (Map, Filter, Reduce), LINQ-style operations (GroupBy, Partition), and utilities not in stdlib.

**Type-safe with generics**: All functions use Go 1.21+ generics with appropriate constraints for compile-time type safety.

**Immutable by default**: Functions don't mutate inputs. Operations return new values.

**100% test coverage**: Every function is comprehensively tested.

## Installation

```bash
go get -u github.com/Goldziher/go-utils
```

## Quick Examples

### Functional Slice Operations

```go
import "github.com/Goldziher/go-utils/sliceutils"

numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Chain functional operations
evens := sliceutils.Filter(numbers, func(v, i int, s []int) bool {
    return v%2 == 0
})

doubled := sliceutils.Map(evens, func(v, i int, s []int) int {
    return v * 2
})

sum := sliceutils.Reduce(doubled, func(acc, v, i int, s []int) int {
    return acc + v
}, 0)
// Result: sum = 60 (2+4+6+8+10 doubled = 4+8+12+16+20 = 60)
```

### LINQ-Style Operations

```go
type User struct {
    Name string
    Age  int
    Role string
}

users := []User{
    {"Alice", 30, "admin"},
    {"Bob", 25, "user"},
    {"Charlie", 30, "user"},
}

// Group by age
byAge := sliceutils.GroupBy(users, func(u User) int { return u.Age })
// Result: map[int][]User{25: [{Bob 25 user}], 30: [{Alice 30 admin}, {Charlie 30 user}]}

// Partition by predicate
admins, regularUsers := sliceutils.Partition(users, func(u User) bool {
    return u.Role == "admin"
})

// Get distinct ages
ages := sliceutils.DistinctBy(users, func(u User) int { return u.Age })
```

### Map Transformations

```go
import "github.com/Goldziher/go-utils/maputils"

data := map[string]int{"apple": 5, "banana": 3, "cherry": 8}

// Extract keys and values
keys := maputils.Keys(data)     // []string{"apple", "banana", "cherry"}
values := maputils.Values(data) // []int{5, 3, 8}

// Filter map entries
filtered := maputils.Filter(data, func(k string, v int) bool {
    return v > 4
})
// Result: map[string]int{"apple": 5, "cherry": 8}
```

### Type-Safe String Conversion

```go
import "github.com/Goldziher/go-utils/stringutils"

// Stringify any type with options
hex := stringutils.Stringify(42, stringutils.Options{Base: 16})  // "2a"
float := stringutils.Stringify(3.14159, stringutils.Options{Precision: 2})  // "3.14"

// Stringify complex types
m := map[string]int{"a": 1, "b": 2}
str := stringutils.Stringify(m)  // "{a: 1, b: 2}"
```

### Business Date Calculations

```go
import "github.com/Goldziher/go-utils/dateutils"

// Add business days (skips weekends)
future := dateutils.AddBusinessDays(time.Now(), 5)

// Check date range overlap
overlaps := dateutils.Overlap(start1, end1, start2, end2)

// Calculate age
age := dateutils.Age(birthdate)
```

## Packages

| Package | Description | Key Functions |
|---------|-------------|---------------|
| **[sliceutils](https://pkg.go.dev/github.com/Goldziher/go-utils/sliceutils)** | Functional and LINQ-style slice operations | Map, Filter, Reduce, GroupBy, Partition, Find |
| **[maputils](https://pkg.go.dev/github.com/Goldziher/go-utils/maputils)** | Map transformations | Keys, Values, Filter, Merge, Invert |
| **[stringutils](https://pkg.go.dev/github.com/Goldziher/go-utils/stringutils)** | String manipulation | Stringify, ToCamelCase, ToSnakeCase, Truncate |
| **[structutils](https://pkg.go.dev/github.com/Goldziher/go-utils/structutils)** | Struct reflection utilities | ToMap, ForEach, FieldNames (tag-aware) |
| **[dateutils](https://pkg.go.dev/github.com/Goldziher/go-utils/dateutils)** | Date/time utilities | AddBusinessDays, Overlap, Age, StartOfWeek |
| **[urlutils](https://pkg.go.dev/github.com/Goldziher/go-utils/urlutils)** | URL and query string builders | QueryStringifyMap, QueryStringifyStruct |
| **[mathutils](https://pkg.go.dev/github.com/Goldziher/go-utils/mathutils)** | Generic math operations | Clamp, InRange, Gcd, Lcm, IsPrime |
| **[ptrutils](https://pkg.go.dev/github.com/Goldziher/go-utils/ptrutils)** | Pointer utilities | ToPtr, Deref |
| **[excutils](https://pkg.go.dev/github.com/Goldziher/go-utils/excutils)** | Exception-style error handling | Panic, Try, Must |

## Documentation

Full documentation available at [goldziher.github.io/go-utils](https://goldziher.github.io/go-utils/)

API reference at [pkg.go.dev/github.com/Goldziher/go-utils](https://pkg.go.dev/github.com/Goldziher/go-utils)

## Contributing

Contributions are welcome! Please read the [Contributing Guide](CONTRIBUTING.md) before submitting PRs.

## License

MIT License - see [LICENSE](LICENSE) for details.

## Author

[Na'aman Hirschfeld](https://github.com/Goldziher)
