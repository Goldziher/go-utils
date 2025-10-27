# IsNil

`func IsNil[T any](ptr *T) bool`

IsNil checks if a pointer is nil. This is a convenience function for readability.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	numPtr := ptr.To(42)
	var nilPtr *int

	fmt.Println(ptr.IsNil(numPtr))  // false
	fmt.Println(ptr.IsNil(nilPtr))  // true
}
```
