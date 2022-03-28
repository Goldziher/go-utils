# Sum

`func Sum[T numbers](slice []T) (result T)`

Sum takes a slice of numbers T, which can be any of the number types, and returns a sum of their values.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := sliceutils.Sum(numerals)

	fmt.Print(result) // 45
}
```
