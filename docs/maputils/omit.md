# Omit

`func Omit[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V`

Omit creates a new map excluding the specified keys. This is the opposite of Pick.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	user := map[string]any{
		"name":     "Alice",
		"age":      30,
		"email":    "alice@example.com",
		"password": "secret",
		"role":     "admin",
	}

	// Remove sensitive fields
	safe := maputils.Omit(user, []string{"password"})

	fmt.Printf("%v\n", safe)
	// map[age:30 email:alice@example.com name:Alice role:admin]
}
```
