# Insert

`func Insert[T any](slice []T, i int, value T) []T`

Insert takes a slice of type T, an index and a value of type T, inserting the value at the given index and shifting any
existing elements to the right.

Deprecated: prefer `slices.Insert(slice, i, value)` from the standard library.

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
