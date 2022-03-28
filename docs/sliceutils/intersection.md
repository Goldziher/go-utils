# Intersection

`func Intersection[T comparable](slices ...[]T) []T`

Intersection takes a variadic number of slices of type T and returns a slice of type T containing any values that are
common to all slices.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first, second := []int{1, 2, 3}, []int{1, 7, 3}

	result := sliceutils.Intersection(first, second)

	fmt.Print(result) // [1, 3]
}
```
