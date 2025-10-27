# sliceutils

Functional programming utilities for slices, inspired by JavaScript Array methods and LINQ.

## Overview

The `sliceutils` package provides type-safe generic functions for slice manipulation. All functions work with any slice type through Go generics and accept callback functions following JavaScript's Array method conventions (value, index, slice).

## Categories

**Functional Operations**: Map, Filter, Reduce, ForEach
**Search Operations**: Find, FindIndex, FindIndexes, FindLastIndex
**Predicates**: Some, Every
**Set Operations**: Union, Intersection, Difference, Unique
**Transformations**: Reverse, Flatten, FlatMap, Chunk, Pluck
**LINQ-style**: GroupBy, Partition, DistinctBy, CountBy, MinBy, MaxBy
**Utilities**: Remove, EnsureUniqueAndAppend, Sum

## When to use stdlib vs sliceutils

**Use stdlib `slices` for**:
- `slices.Index(slice, value)` - find index of value
- `slices.Contains(slice, value)` - check if slice contains value
- `slices.Clone(slice)` - copy a slice
- `slices.Concat(slices...)` - merge slices
- `slices.Sort(slice)` - sort a slice

**Use sliceutils for**:
- Functional patterns with callbacks (Map, Filter, Reduce)
- LINQ-style operations (GroupBy, Partition, DistinctBy)
- Complex search operations (FindIndex with predicate)
- Set operations (Union, Intersection, Difference)

## Example

```go
import "github.com/Goldziher/go-utils/sliceutils"

numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Functional patterns
evens := sliceutils.Filter(numbers, func(v, i int, s []int) bool {
    return v%2 == 0
})

doubled := sliceutils.Map(evens, func(v, i int, s []int) int {
    return v * 2
})

sum := sliceutils.Reduce(doubled, func(acc, v, i int, s []int) int {
    return acc + v
}, 0)

// LINQ-style operations
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
byAge := sliceutils.GroupBy(users, func(u User) int {
    return u.Age
})

// Partition by role
admins, regularUsers := sliceutils.Partition(users, func(u User) bool {
    return u.Role == "admin"
})

// Get distinct ages
ages := sliceutils.Map(users, func(u User, i int, s []User) int {
    return u.Age
})
uniqueAges := sliceutils.Unique(ages)
```
