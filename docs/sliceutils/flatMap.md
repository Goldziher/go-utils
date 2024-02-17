# FlatMap

`func FlatMap[T any, R any](slice []T, mapper func(value T, index int, slice []T) []R) []R`

FlatMap receives a slice of type T, executes the passed in slice-mapper function for each element in the slice, and
returns a flattened slice containing all the elements from all the mapped slices.
The function is passed the current element, the current index and the slice itself as function arguments.


```go
package main

import (
	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	items := []int{1, 2, 3, 4}

	flatMapped := sliceutils.FlatMap(items, func(value int, index int, slice []int) []int {
		return []int{value, value * 2}
	}) // []int{1, 2, 2, 4, 3, 6, 4, 8}
}
```
