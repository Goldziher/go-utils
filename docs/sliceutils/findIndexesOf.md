# FindIndexesOf

`func FindIndexes[T any](slice []T, predicate func(value T, index int, slice []T) bool) []int`

FindIndexesOf takes a slice of type T and a value of type T, returning a slice containing all indexes where elements
equal the passed in value.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindIndexesOf(friends, "John")

	fmt.Print(result) // [0, 4]
}
```
