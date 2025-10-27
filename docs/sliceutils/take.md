# Take

`func Take[T any](slice []T, n int) []T`

Take returns the first n elements from the slice. If n is greater than the slice length, returns the entire slice. If n is zero or negative, returns an empty slice.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	first3 := sliceutils.Take(numbers, 3)
	fmt.Printf("%v\n", first3) // [1 2 3]

	tooMany := sliceutils.Take(numbers, 100)
	fmt.Printf("%v\n", tooMany) // [1 2 3 4 5 6 7 8 9 10]
}
```
