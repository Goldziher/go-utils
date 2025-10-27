# DerefSlice

`func DerefSlice[T any](ptrs []*T, def T) []T`

DerefSlice dereferences all pointers in a slice, using def for nil pointers. Returns a new slice with dereferenced values.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	ptrs := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}

	// Dereference with default value of 0 for nils
	values := ptr.DerefSlice(ptrs, 0)
	fmt.Println(values)  // [1 0 3 0 5]
}
```
