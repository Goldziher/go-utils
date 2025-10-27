# Coalesce

`func Coalesce[T any](ptrs ...*T) *T`

Coalesce returns the first non-nil pointer from the list. Returns nil if all pointers are nil.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	var a, b *int
	c := ptr.To(42)
	d := ptr.To(99)

	// Get first non-nil pointer
	result := ptr.Coalesce(a, b, c, d)
	fmt.Println(*result)  // 42

	// All nil
	result = ptr.Coalesce((*int)(nil), (*int)(nil))
	fmt.Println(result)  // <nil>
}
```
