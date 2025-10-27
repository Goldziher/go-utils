# ToSlice

`func ToSlice[T any](ptr *T) []T`

ToSlice converts a pointer to a slice. Returns an empty slice if ptr is nil, otherwise returns a slice with the single element.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	numPtr := ptr.To(42)
	slice := ptr.ToSlice(numPtr)
	fmt.Println(slice)  // [42]

	var nilPtr *int
	emptySlice := ptr.ToSlice(nilPtr)
	fmt.Println(emptySlice)  // []
}
```
