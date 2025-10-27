# maputils

Utilities for map transformations and operations.

## Overview

The `maputils` package provides type-safe generic functions for working with maps. All functions leverage Go generics to work with any map type `map[K]V` where `K` is `comparable`.

## Functions

**Extraction**: Keys, Values
**Transformation**: Filter, Map
**Iteration**: ForEach
**Combination**: Merge
**Manipulation**: Drop, Invert, Pick, Omit

## Example

```go
import "github.com/Goldziher/go-utils/maputils"

data := map[string]int{
    "apple":  5,
    "banana": 3,
    "cherry": 8,
}

// Extract keys and values
keys := maputils.Keys(data)     // []string{"apple", "banana", "cherry"}
values := maputils.Values(data) // []int{5, 3, 8}

// Filter entries
filtered := maputils.Filter(data, func(k string, v int) bool {
    return v > 4
})
// Result: map[string]int{"apple": 5, "cherry": 8}

// Transform map
transformed := maputils.Map(data, func(k string, v int) (string, string) {
    return k, fmt.Sprintf("count: %d", v)
})
// Result: map[string]string{"apple": "count: 5", ...}

// Merge maps
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
merged := maputils.Merge(m1, m2)
// Result: map[string]int{"a": 1, "b": 3, "c": 4}
```
