# Invert

`func Invert[K, V comparable](mapInstance map[K]V) map[V]K`

Invert swaps keys and values in a map. Both keys and values must be comparable types. If multiple keys have the same value, only one will remain (non-deterministic which one).

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	roles := map[string]int{
		"admin": 1,
		"user":  2,
		"guest": 3,
	}

	// Invert to map from ID to role name
	byID := maputils.Invert(roles)

	fmt.Printf("%v\n", byID)
	// map[1:admin 2:user 3:guest]
}
```
