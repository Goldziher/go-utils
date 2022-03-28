# Unique

`func Unique[T comparable](slice []T) []T`

Unique tales a slice of type T and returns a slice of type T containing all unique elements.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 3, 1}

	result := sliceutils.Unique(numerals)

	fmt.Print(result) // [0, 1, 2, 3]
}
```
