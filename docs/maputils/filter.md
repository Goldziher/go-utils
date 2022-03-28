# Filter

`func Filter[K comparable, V any](mapInstance map[K]V, function func(key K, value V) bool) map[K]V`

Filter takes a map with keys K and values V, and executes the passed in function for each key-value pair. If the filter
function returns true, the key-value pair will be included in the output, otherwise it is filtered out.

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
	}

	result := maputils.Filter(vegetables, func(key string, value int) bool {
        return value == 5
	})

	fmt.Print(result) // { "potatoes": 5 }
}
```
