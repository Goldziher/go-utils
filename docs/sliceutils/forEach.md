# ForEach

`func ForEach[T any](slice []T, function func(value T, index int, slice []T))`

ForEach executes the passed in function for each element in the given slice. The function is passed the current element,
the current index and the slice as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	result := 0
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceutils.ForEach(numerals, func(value int, index int, slice []int) {
		result += value
	})

	fmt.Print(result) // 45
}
```
