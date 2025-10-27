# Pick

`func Pick[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V`

Pick creates a new map containing only the specified keys. Keys that don't exist in the original map are ignored.

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

	// Pick only public fields
	public := maputils.Pick(user, []string{"name", "age", "role"})

	fmt.Printf("%v\n", public)
	// map[age:30 name:Alice role:admin]
}
```
