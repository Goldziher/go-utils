# Windows

`func Windows[T any](slice []T, size int) [][]T`

Windows returns a slice of sliding windows of the specified size. Each window is a slice of consecutive elements. If size is greater than the slice length or less than 1, returns an empty slice.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Get all windows of size 3
	windows := sliceutils.Windows(numbers, 3)
	fmt.Printf("%v\n", windows)
	// [[1 2 3] [2 3 4] [3 4 5]]

	// Useful for calculating rolling averages, moving statistics, etc.
}
```
