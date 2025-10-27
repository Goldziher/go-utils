# Skip

`func Skip[T any](slice []T, n int) []T`

Skip returns the slice with the first n elements removed. If n is greater than or equal to the slice length, returns an empty slice. If n is zero or negative, returns the original slice.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	skip3 := sliceutils.Skip(numbers, 3)
	fmt.Printf("%v\n", skip3) // [4 5 6 7 8 9 10]

	skipAll := sliceutils.Skip(numbers, 100)
	fmt.Printf("%v\n", skipAll) // []
}
```
