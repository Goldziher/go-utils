# MapKeys

`func MapKeys[K comparable, V any, R comparable](mapInstance map[K]V, mapper func(key K, value V) R) map[R]V`

MapKeys transforms map keys using a mapper function, returning a new map with transformed keys. If the mapper produces duplicate keys, later values will overwrite earlier ones.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	data := map[string]int{
		"apple":  5,
		"BANANA": 3,
		"Cherry": 8,
	}

	// Normalize keys to lowercase
	normalized := maputils.MapKeys(data, func(k string, v int) string {
		return strings.ToLower(k)
	})

	fmt.Printf("%v\n", normalized)
	// map[apple:5 banana:3 cherry:8]
}
```
