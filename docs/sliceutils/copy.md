# Copy

> **⚠️ DEPRECATED:** This function is deprecated as of Go 1.21. Use `slices.Clone` from the standard library instead.

`func Copy[T any](slice []T) []T`

Copy receives a slice of type T and returns a copy.

## Migration to stdlib

```go
import "slices"

// Old
copied := sliceutils.Copy(original)

// New
copied := slices.Clone(original)
```

## Original Example

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1}

	numeralsCopy := sliceutils.Copy(numerals)
	numeralsCopy[0] = 1

	fmt.Print(numerals)     // [0, 1]
	fmt.Print(numeralsCopy) // [1, 1]
}
```
