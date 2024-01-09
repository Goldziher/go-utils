# EnsureUniqueAndAppend

`func EnsureUniqueAndAppend[T comparable](slice []T, item T) []T`

EnsureUniqueAndAppend appends an item to a slice if it does not already exists.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	slice := []string{}
	item := "go-utils"

	slice = sliceutils.EnsureUniqueAndAppend(slice, item) // ["go-utils"]
	slice = sliceutils.EnsureUniqueAndAppend(slice, item) // ["go-utils"]
}
```
