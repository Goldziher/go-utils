# Merge

`func Merge[T any](slices ...[]T) (mergedSlice []T)`

Merge takes slices of type T and merges them into a single slice of type T, preserving their order.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}

	result := sliceutils.Merge(first, second, third)

	fmt.Print(result) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
}
```
