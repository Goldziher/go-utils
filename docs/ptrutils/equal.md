# Equal

`func Equal[T comparable](a, b *T) bool`

Equal compares two pointers for equality. Returns true if both are nil or both point to equal values. Returns false if one is nil and the other is not.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	a := ptr.To(5)
	b := ptr.To(5)
	c := ptr.To(10)
	var d *int

	fmt.Println(ptr.Equal(a, b))  // true (same values)
	fmt.Println(ptr.Equal(a, c))  // false (different values)
	fmt.Println(ptr.Equal(a, d))  // false (one nil, one not)
	fmt.Println(ptr.Equal((*int)(nil), (*int)(nil)))  // true (both nil)
}
```
