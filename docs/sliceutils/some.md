# Some

`func Some[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool`

Some takes a slice of type T and a predicate function, returning true if the predicate returned true for **some**
elements. The function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Some(friends, func(value string, index int, slice []string) bool {
		return value == "Mandy"
	})

	fmt.Print(result) // true
}
```
