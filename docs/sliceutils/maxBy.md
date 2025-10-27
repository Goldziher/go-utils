# MaxBy

`func MaxBy[T any, O cmp.Ordered](slice []T, selector func(T) O) *T`

MaxBy returns a pointer to the element with the maximum value as determined by the selector function. Returns nil if the slice is empty.

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

	expensive := sliceutils.MaxBy(products, func(p Product) float64 {
		return p.Price
	})

	if expensive != nil {
		fmt.Printf("Most expensive: %s at $%.2f\n", expensive.Name, expensive.Price)
		// Most expensive: Cherry at $2.00
	}
}
```
