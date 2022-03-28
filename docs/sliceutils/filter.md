# Filter

`func Filter[T any](slice []T, predicate func(value T, index int, slice []T) bool) []T`

Filter takes a slice of type `T` and filters it using the given predicate function. The predicate is passed the current
element, the current index and the slice as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddNumbers := sliceutils.Filter(numerals, func(value int, index int, slice []int) bool {
		return value%2 != 0
	})

	fmt.Printf("%v", oddNumbers) // [1 3 5 7 9]
}
```
