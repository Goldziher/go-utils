# MinBy

`func MinBy[T any, O cmp.Ordered](slice []T, selector func(T) O) *T`

MinBy returns a pointer to the element with the minimum value as determined by the selector function. Returns nil if the slice is empty.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	type Product struct {
		Name  string
		Price float64
	}

	products := []Product{
		{"Apple", 1.50},
		{"Banana", 0.75},
		{"Cherry", 2.00},
	}

	cheapest := sliceutils.MinBy(products, func(p Product) float64 {
		return p.Price
	})

	if cheapest != nil {
		fmt.Printf("Cheapest: %s at $%.2f\n", cheapest.Name, cheapest.Price)
		// Cheapest: Banana at $0.75
	}
}
```
