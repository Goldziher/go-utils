# Merge

> **⚠️ DEPRECATED:** This function is deprecated as of Go 1.21. Use `slices.Concat` from the standard library instead.

`func Merge[T any](slices ...[]T) (mergedSlice []T)`

Merge takes slices of type T and merges them into a single slice of type T, preserving their order.

## Migration to stdlib

```go
import "slices"

// Old
merged := sliceutils.Merge(slice1, slice2, slice3)

// New
merged := slices.Concat(slice1, slice2, slice3)
```

## Original Example

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}

	result := sliceutils.Merge(first, second, third)

	fmt.Print(result) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
}
```
