# Insert

> **⚠️ DEPRECATED:** This function is deprecated as of Go 1.21. Use `slices.Insert` from the standard library instead.

`func Insert[T any](slice []T, i int, value T) []T`

Insert takes a slice of type T, an index and a value of type T, inserting the value at the given index and shifting any
existing elements to the right.

## Migration to stdlib

```go
import "slices"

// Old
result := sliceutils.Insert(items, 2, "newItem")

// New
result := slices.Insert(items, 2, "newItem")

// Stdlib version also supports multiple insertions
result := slices.Insert(items, 2, "item1", "item2", "item3")
```

## Original Example

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 2}

	numerals = sliceutils.Insert(numerals, 1, 1)

	fmt.Print(numerals) // [0, 1, 2]
}
```
