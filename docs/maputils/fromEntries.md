# FromEntries

`func FromEntries[K comparable, V any](entries [][2]any) map[K]V`

FromEntries creates a map from a slice of key-value pairs. If duplicate keys exist, later values overwrite earlier ones.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	entries := [][2]any{
		{"name", "Alice"},
		{"age", 30},
		{"role", "admin"},
	}

	user := maputils.FromEntries[string, any](entries)

	fmt.Printf("%v\n", user)
	// map[age:30 name:Alice role:admin]
}
```
