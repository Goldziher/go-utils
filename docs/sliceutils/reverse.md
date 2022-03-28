# Reverse

`func Reverse[T any](slice []T) []T`

Reverse takes a slice of type T and returns a slice of type T with a reverse order of elements.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	reversed := sliceutils.Reverse(numerals)

	fmt.Print(reversed) // [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
}
```
