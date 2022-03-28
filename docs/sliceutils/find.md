# Find

`func Find[T any](slice []T, predicate func(value T, index int, slice []T) bool) *T`

Find takes a slice of type T and executes the passed in predicate function for each element in the slice. If the
predicate returns true, a pointer to the element is returned. If no element is found, nil is returned. The function is
passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	result := sliceutils.Find(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})

	fmt.Print(result) // "Wednesday"
}
```
