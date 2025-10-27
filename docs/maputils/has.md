# Has

`func Has[K comparable, V any](mapInstance map[K]V, key K) bool`

Has checks if a key exists in the map. This is equivalent to the two-value form of map lookup but more expressive.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}

	if maputils.Has(config, "host") {
		fmt.Println("Host is configured")
	}

	if !maputils.Has(config, "database") {
		fmt.Println("Database not configured")
	}
}
```
