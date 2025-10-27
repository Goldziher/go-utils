# ToEntries

`func ToEntries[K comparable, V any](mapInstance map[K]V) [][2]any`

ToEntries converts a map to a slice of key-value pairs. Order is non-deterministic due to map iteration order.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Charlie": 92,
	}

	entries := maputils.ToEntries(scores)

	for _, entry := range entries {
		name := entry[0].(string)
		score := entry[1].(int)
		fmt.Printf("%s: %d\n", name, score)
	}
}
```
