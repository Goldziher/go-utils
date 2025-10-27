# TakeLast

`func TakeLast[T any](slice []T, n int) []T`

TakeLast returns the last n elements from the slice. If n is greater than the slice length, returns the entire slice. If n is zero or negative, returns an empty slice.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	last3 := sliceutils.TakeLast(numbers, 3)
	fmt.Printf("%v\n", last3) // [8 9 10]
}
```
