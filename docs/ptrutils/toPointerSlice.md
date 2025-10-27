# ToPointerSlice

`func ToPointerSlice[T any](values []T) []*T`

ToPointerSlice converts a slice of values to a slice of pointers. Each element in the result points to the corresponding element in the input.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	values := []int{1, 2, 3, 4, 5}

	// Convert to pointer slice
	ptrs := ptr.ToPointerSlice(values)

	for _, p := range ptrs {
		fmt.Println(*p)  // 1, 2, 3, 4, 5
	}
}
```
