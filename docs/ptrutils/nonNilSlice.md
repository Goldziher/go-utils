# NonNilSlice

`func NonNilSlice[T any](ptrs []*T) []*T`

NonNilSlice returns a new slice containing only non-nil pointers from the input.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	ptrs := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}

	// Filter out nil pointers
	nonNil := ptr.NonNilSlice(ptrs)
	fmt.Printf("Length: %d\n", len(nonNil))  // 3

	for _, p := range nonNil {
		fmt.Println(*p)  // 1, 3, 5
	}
}
```
