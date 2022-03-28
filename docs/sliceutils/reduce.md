# Reduce

`func Reduce[T any, R any](slice []T, reducer func(acc R, value T, index int, slice []T) R, initial R) R`

Reduce allows transforming the slice and its values into a different value by executing the given reducer function for
each element in the slice. The function is passed the accumulator, current element, current index and the slice as
function arguments. The third argument to reduce is the initial value.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := sliceutils.Reduce(
		numerals,
		func(acc int, cur int, index int, slice []int) int {
			return acc + cur
		},
		0,
	)

	fmt.Print(sum) // 45
}
```
