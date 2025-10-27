# Last

`func Last[T any](slice []T) *T`

Last returns a pointer to the last element of the slice. Returns nil if the slice is empty.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	lastElem := sliceutils.Last(numbers)
	if lastElem != nil {
		fmt.Printf("Last: %d\n", *lastElem) // Last: 5
	}

	empty := []int{}
	noLast := sliceutils.Last(empty)
	fmt.Printf("Empty slice last: %v\n", noLast) // nil
}
```
