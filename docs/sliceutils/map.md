# Map

`func Map[T any, R any](slice []T, mapper func(value T, index int, slice []T) R) (mapped []R)`

Map allows transforming the values in a slice by executing the given mapper function for each element in the slice. The
function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := sliceutils.Map(numerals, func(value int, index int, slice []int) int {
		return value * 2
	})

	fmt.Print(result) // [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
}
```
