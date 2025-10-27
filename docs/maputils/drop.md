# Drop

`func Drop[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V`

Drop takes a map with keys K and values V, and a slice of keys K, dropping all the key-value pairs that match the keys in the slice.

Note: this function will modify the passed in map. To get a different object, use the Copy function to pass a copy to this function.

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
		"tomatoes": 3,
	}

	// Drop will modify the original map
	result := maputils.Drop(vegetables, []string{"carrots", "tomatoes"})

	fmt.Print(result) // { "potatoes": 5 }
	fmt.Print(vegetables) // { "potatoes": 5 } - original map was modified

	// To avoid modifying the original, use Copy first
	vegetables2 := map[string]int{
		"potatoes": 5,
		"carrots":  10,
		"tomatoes": 3,
	}

	copied := maputils.Copy(vegetables2)
	result2 := maputils.Drop(copied, []string{"carrots"})

	fmt.Print(result2) // { "potatoes": 5, "tomatoes": 3 }
	fmt.Print(vegetables2) // { "potatoes": 5, "carrots": 10, "tomatoes": 3 } - original unchanged
}
```
