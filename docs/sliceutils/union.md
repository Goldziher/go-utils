# Union

`func Union[T comparable](slices ...[]T) []T`

Union takes a variadic number of slices of type T and returns a slice of type T containing the unique elements in the
different slices.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first, second := []int{1, 2, 3}, []int{1, 7, 3}

	result := sliceutils.Union(first, second)

	fmt.Print(result) // [1, 2, 3, 7]
}
```
