# Keys

`func Keys[K comparable, V any](mapInstance map[K]V) []K`

Keys takes a map with keys K and values V and returns a slice of type K with the map's keys.

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

	result := maputils.Keys(vegetables)

	fmt.Print(result) // ["potatoes", "carrots"]
}
```
