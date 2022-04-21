# QueryStringifyMap

`func QueryStringifyMap[K comparable, V any](values map[K]V) string`

Creates a query string from a given map instance.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/urlutils"
)

func main() {
	values := map[string]any{
		"user":    "moishe",
		"active":  true,
		"age":     100,
		"friends": []int{1, 2, 3, 4, 5, 6},
	}

	result := urlutils.QueryStringifyMap(values)

	fmt.Print(result) // "active=true&age=100&friends=1&friends=2&friends=3&friends=4&friends=5&friends=6&user=moishe"
}
```
