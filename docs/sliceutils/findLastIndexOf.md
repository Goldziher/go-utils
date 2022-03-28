# FindLastIndexOf

`func FindLastIndexOf[T comparable](slice []T, value T) int`

FindLastIndexOf takes a slice of type T and a value of type T. If any element in the slice equals the given value, the
last index of is occurrence is returned. If no element is found, `-1` is returned.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindLastIndexOf(friends, "John")

	fmt.Print(result) // 4
}
```
