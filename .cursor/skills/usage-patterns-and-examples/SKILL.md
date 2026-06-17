---
name: usage-patterns-and-examples
priority: high
description: usage-patterns-and-examples
---

# Usage Patterns and Examples

## Usage Patterns and Examples

### Installation in Your Project

```bash
go get -u github.com/Goldziher/go-utils
```

### Common Usage Patterns

#### Working with Slices (JavaScript-style)

The `sliceutils` package provides familiar operations for developers coming from JavaScript:

```go
import "github.com/Goldziher/go-utils/sliceutils"

// Filter and Map (chainable pattern)
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

filtered := sliceutils.Filter(numbers, func(val int, idx int, slice []int) bool {
    return val%2 == 0
}) // [2, 4, 6, 8, 10]

doubled := sliceutils.Map(filtered, func(val int, idx int, slice []int) int {
    return val * 2
}) // [4, 8, 12, 16, 20]

// Reduce for aggregation
sum := sliceutils.Reduce(numbers, func(acc int, cur int, idx int, slice []int) int {
    return acc + cur
}, 0) // 55

// Set operations
a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}

union := sliceutils.Union(a, b)              // [1, 2, 3, 4, 5, 6]
intersection := sliceutils.Intersection(a, b) // [3, 4]
difference := sliceutils.Difference(a, b)     // [1, 2, 5, 6]
```

#### Working with Maps

```go
import "github.com/Goldziher/go-utils/maputils"

data := map[string]int{"a": 1, "b": 2, "c": 3}

// Extract keys and values
keys := maputils.Keys(data)     // ["a", "b", "c"]
values := maputils.Values(data) // [1, 2, 3]

// Filter map entries
filtered := maputils.Filter(data, func(key string, value int) bool {
    return value > 1
}) // {"b": 2, "c": 3}

// Merge maps (right-to-left precedence)
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
merged := maputils.Merge(m1, m2) // {"a": 1, "b": 3, "c": 4}
```

#### Working with Structs

```go
import "github.com/Goldziher/go-utils/structutils"

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

// Convert struct to map (respects struct tags)
userMap := structutils.ToMap(user, "json")
// {"name": "Alice", "age": 30, "email": "alice@example.com"}
```

#### URL Query String Building

```go
import "github.com/Goldziher/go-utils/urlutils"

// From map
params := map[string]any{
    "search": "golang",
    "page":   1,
    "tags":   []string{"generics", "utils"},
}
query := urlutils.QueryStringifyMap(params)
// "page=1&search=golang&tags=generics&tags=utils"

// From struct
type QueryParams struct {
    Search string   `qs:"q"`
    Page   int      `qs:"page"`
    Tags   []string `qs:"tag"`
}
params := QueryParams{Search: "golang", Page: 1, Tags: []string{"generics", "utils"}}
query := urlutils.QueryStringifyStruct(params, "qs")
// "page=1&q=golang&tag=generics&tag=utils"
```

### Type Safety with Generics

All functions provide compile-time type safety:

```go
// Compile error: type mismatch
numbers := []int{1, 2, 3}
result := sliceutils.Map(numbers, func(val int, idx int, slice []int) string {
    return strconv.Itoa(val) // Returns []string, not []int
})
```
