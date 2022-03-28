# FindIndex

`func FindIndex[T any](slice []T, predicate func(value T, index int, slice []T) bool) int`

FindIndex takes a slice of type T and executes the passed in predicate function for each element in the slice. If the
predicate returns true, the element's index is returned. If no element is found, `-1` is returned. The function is
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

	result := sliceutils.FindIndex(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})

	fmt.Print(result) // 3
}
```
