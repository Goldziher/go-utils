# SkipLast

`func SkipLast[T any](slice []T, n int) []T`

SkipLast returns the slice with the last n elements removed. If n is greater than or equal to the slice length, returns an empty slice. If n is zero or negative, returns the original slice.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	skipLast3 := sliceutils.SkipLast(numbers, 3)
	fmt.Printf("%v\n", skipLast3) // [1 2 3 4 5 6 7]
}
```
