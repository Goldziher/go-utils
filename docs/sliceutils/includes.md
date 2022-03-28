# Includes

`func Includes[T comparable](slice []T, value T) bool`

Includes receives a slice of type T and a value of type T, determining whether or not the value is included in the
slice.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Includes(friends, "John")

	fmt.Print(result) // true
}
```
