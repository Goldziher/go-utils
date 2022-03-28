# Copy

`func Copy[K comparable, V any](mapInstance map[K]V) map[K]V`

Copy takes a map with keys K and values V and returns a copy.

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

	copiedVegetables := maputils.Copy(vegetables)

	copiedVegetables["potatoes"] = 3

	fmt.Print(vegetables["potatoes"])       // 5
	fmt.Print(copiedVegetables["potatoes"]) // 3
}
```
