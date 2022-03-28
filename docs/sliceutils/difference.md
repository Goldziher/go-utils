# Difference

`func Difference[T comparable](slices ...[]T) []T`

Difference takes a variadic number of slices of type T and returns a slice of type T containing the elements that are
different between the slices.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first, second := []int{1, 2, 3}, []int{1, 7, 3}

	result := sliceutils.Difference(first, second)

	fmt.Print(result) // [2, 7]
}
```
