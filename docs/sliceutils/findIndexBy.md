# FindIndexBy

`func FindIndexBy[T comparable](input []T, predicate func(value T) bool) int`

FindIndexBy takes a slice of type T and a predicate and returns the index of the
first element that satisfies the predicate, and `-1` otherwise.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1,2,3,4,5,6}

	result := sliceutils.FindIndexBy(numbers, func(value int) bool { return value % 2 == 0 })
	fmt.Print(result) // 1
}
```
