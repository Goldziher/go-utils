# Tail

`func Tail[T any](slice []T) []T`

Tail returns all elements of the slice except the first. Returns an empty slice if the slice has 0 or 1 elements.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	rest := sliceutils.Tail(numbers)
	fmt.Printf("%v\n", rest) // [2 3 4 5]

	single := []int{1}
	emptyTail := sliceutils.Tail(single)
	fmt.Printf("%v\n", emptyTail) // []
}
```
