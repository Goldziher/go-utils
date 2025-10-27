# Map

`func Map[K comparable, V any, R any](mapInstance map[K]V, mapper func(key K, value V) R) map[K]R`

Map transforms map values using a mapper function, returning a new map with transformed values while preserving keys.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	prices := map[string]int{
		"apple":  100,
		"banana": 75,
		"cherry": 200,
	}

	// Transform values to strings
	priceStrings := maputils.Map(prices, func(k string, v int) string {
		return fmt.Sprintf("$%d.%02d", v/100, v%100)
	})

	fmt.Printf("%v\n", priceStrings)
	// map[apple:$1.00 banana:$0.75 cherry:$2.00]
}
```
