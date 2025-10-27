# Compact

`func Compact[T comparable](slice []T) []T`

Compact removes all zero values from the slice. A zero value is determined by the zero value of type T (0 for numbers, "" for strings, nil for pointers, false for bools, etc.).

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	// Remove zero numbers
	numbers := []int{1, 0, 2, 0, 3, 0, 4, 5}
	compacted := sliceutils.Compact(numbers)
	fmt.Printf("%v\n", compacted) // [1 2 3 4 5]

	// Remove empty strings
	strings := []string{"hello", "", "world", "", "!"}
	compactedStrings := sliceutils.Compact(strings)
	fmt.Printf("%v\n", compactedStrings) // [hello world !]
}
```
