# Values

`func Values[K comparable, V any](mapInstance map[K]V) []V`

Values takes a map with keys K and values V and returns a slice of type V with the map's values.

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

	result := maputils.Values(vegetables)

	fmt.Print(result) // [5, 10]
}
```
