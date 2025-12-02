# Copy

`func Copy[T any](slice []T) []T`

Copy receives a slice of type T and returns a copy.

Deprecated: prefer `slices.Clone(slice)` from the standard library.

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
