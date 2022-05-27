# Chunk

`func Chunk[T any](input []T, size int) [][]T `

Unique takes a slice of type T and size N and returns a slice of slices T of size N.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result1 := sliceutils.Chunk(numbers, 2) // [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	result2 := sliceutils.Chunk(numbers, 3) // [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
}
```
