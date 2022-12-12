# InSlice

`func InSlice[T comparable](input []T, item T) bool`

InSlice receives a slice of type T and an item of type T and returns true if item is in slice, false otherwise.
It's a shortcut around Find for the case where it's more convenient to manipulate a bool instead of the requested items

```go
package main

import (
	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	items := []string{
		"Item 1",
		"Item 2",
	}

	found := sliceutils.InSlice(items, "Item 1") // true
	found = sliceutils.InSlice(items, "Item 3) // false
}
```
