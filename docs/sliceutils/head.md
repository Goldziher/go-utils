# Head

`func Head[T any](slice []T) *T`

Head returns a pointer to the first element of the slice. Returns nil if the slice is empty.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	first := sliceutils.Head(numbers)
	if first != nil {
		fmt.Printf("First: %d\n", *first) // First: 1
	}

	empty := []int{}
	noFirst := sliceutils.Head(empty)
	fmt.Printf("Empty slice head: %v\n", noFirst) // nil
}
```
