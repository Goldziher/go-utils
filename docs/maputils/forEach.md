# ForEach

`func ForEach[K comparable, V any](mapInstance map[K]V, function func(key K, value V))`

ForEach given a map with keys K and values V, executes the passed in function for each key-value pair.

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

	maputils.ForEach(vegetables, func(key string, value int) {
		fmt.Printf("Buy %d Kg of %s", value, key) //  "Buy 5 Kg of potatoes", "Buy 10 Kg of carrots"
	})
}
```
