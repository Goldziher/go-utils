# Merge

`func Merge[K comparable, V any](mapInstances ...map[K]V) map[K]V`

Merge takes a variadic numbers of maps with keys K and values V and returns a merged map. Merging is done from left to
right. If a key already exists in a previous map, its value is over-written.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	vegetables := map[string]int{
		"potatoes": 5,
		"carrots":  10,
		"tomatoes": 3,
	}

	fruits := map[string]int{
		"bananas":  3,
		"tomatoes": 5,
	}

	result := maputils.Merge(vegetables, fruits)

	fmt.Print(result) //{ "potatoes": 5, "carrots":  10, "tomatoes": 5, "bananas": 3 }
}
```
