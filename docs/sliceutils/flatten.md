# Flatten

`Flatten[I any](input [][]I) []I`

Flatten - receives a slice of slice of type I and flattens it to a slice of type I.

```go
package main

import (
	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	items := [][]int{
		{1, 2, 3, 4},
		{5, 6},
		{7, 8},
		{9, 10, 11},
	}

	flattened := sliceutils.Flatten(items) //[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
}
```
