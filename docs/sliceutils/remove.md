# Remove

`func Remove[T any](slice []T, i int) []T`

Remove takes a slice of type T and an index, removing the element at the given index.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	numerals = sliceutils.Remove(numerals, 3)

	fmt.Print(numerals) // [0, 1, 2, 4, 5, 6, 7, 8, 9]
}
```
