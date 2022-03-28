# FindIndexes

`func FindIndexes[T any](slice []T, predicate func(value T, index int, slice []T) bool) []int`

FindIndexes takes a slice of type T and executes the passed in predicate function for each element in the slice,
returning a slice containing all indexes for which the predicate returned true. The function is passed the current
element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindIndexes(friends, func(value string, index int, slice []string) bool {
		return value == "John"
	})

	fmt.Print(result) // [0, 4]
}
```
