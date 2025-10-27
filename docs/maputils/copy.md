# Copy

> **⚠️ DEPRECATED:** This function is deprecated as of Go 1.21. Use `maps.Clone` from the standard library instead.

`func Copy[K comparable, V any](mapInstance map[K]V) map[K]V`

Copy takes a map with keys K and values V and returns a copy.

## Migration to stdlib

```go
import "maps"

// Old
copied := maputils.Copy(original)

// New
copied := maps.Clone(original)
```

## Original Example

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	vegetables := map[string]int{
		"potatoes": 5,
		"carrots":  10,
	}

	copiedVegetables := maputils.Copy(vegetables)

	copiedVegetables["potatoes"] = 3

	fmt.Print(vegetables["potatoes"])       // 5
	fmt.Print(copiedVegetables["potatoes"]) // 3
}
```
