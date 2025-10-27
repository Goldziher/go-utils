# Get

`func Get[K comparable, V any](mapInstance map[K]V, key K, defaultValue V) V`

Get safely gets a value from the map with a default fallback if the key doesn't exist.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	config := map[string]int{
		"maxConnections": 100,
		"timeout":        30,
	}

	// Get existing key
	maxConn := maputils.Get(config, "maxConnections", 50)
	fmt.Printf("Max connections: %d\n", maxConn) // 100

	// Get missing key with default
	bufferSize := maputils.Get(config, "bufferSize", 1024)
	fmt.Printf("Buffer size: %d\n", bufferSize) // 1024
}
```
