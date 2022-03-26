### Filter

`func Filter[T any](slice []T, predicate func(value T, index int, slice []T) bool) []T`

Filters a given slice of type `T` using a passed in predicate function. The predicate is passed the current element, the
current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"go-utils/sliceutils"
)

func main() {
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddNumbers := sliceutils.Filter(numerals, func(value int, _ int, _ []int) bool {
		return value%2 != 0
	})

	fmt.Printf("%v", oddNumbers) // [1 3 5 7 9]
}
```

### ForEach

`func ForEach[T any](slice []T, function func(value T, index int, slice []T))`

Executes the passed in function for each element in the given slice. The function is passed the current element, the
current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"go-utils/sliceutils"
)

func main() {
	result := 0
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceutils.ForEach(numerals, func(value int, _ int, _ []int) {
		result += value
	})

	fmt.Print(result) // 45
}
```
