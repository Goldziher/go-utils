# Initial

`func Initial[T any](slice []T) []T`

Initial returns all elements of the slice except the last. Returns an empty slice if the slice has 0 or 1 elements.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	allButLast := sliceutils.Initial(numbers)
	fmt.Printf("%v\n", allButLast) // [1 2 3 4]

	single := []int{1}
	emptyInitial := sliceutils.Initial(single)
	fmt.Printf("%v\n", emptyInitial) // []
}
```
