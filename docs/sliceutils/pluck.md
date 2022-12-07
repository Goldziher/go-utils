# Pluck

`func Pluck[I any, O any](input []I, getter func(I) *O) []O `

Pluck receives a slice of type I and a getter func to a field and returns an array containing requested field from each slice's item.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	items := []Pluckable{
		{
			Code:  "azer",
			Value: "Azer",
		},
		{
			Code:  "tyuio",
			Value: "Tyuio",
		},
	}

	result1 := sliceutils.Pluck(items, func(item Pluckable) *string {
		return &item.Code
	}) // []string{"azer", "tyuio""}
	result2 := sliceutils.Pluck(items, func(item Pluckable) *string {
		return &item.Value
	}) // []string{"Azer", "Tyuio"}
}
```
