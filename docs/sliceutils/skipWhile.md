# SkipWhile

`func SkipWhile[T any](slice []T, predicate func(T) bool) []T`

SkipWhile skips elements from the beginning of the slice while the predicate returns true. Returns the remaining elements starting from the first element where the predicate returns false.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Skip while less than 5
	result := sliceutils.SkipWhile(numbers, func(n int) bool {
		return n < 5
	})

	fmt.Printf("%v\n", result) // [5 6 7 8 9 10]
}
```
