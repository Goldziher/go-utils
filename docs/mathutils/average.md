# Average

`func Average[T Number](values []T) float64`

Average returns the average (mean) of all values in a slice. Returns 0 for empty slice.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	integers := []int{1, 2, 3, 4, 5}
	fmt.Println(mathutils.Average(integers))  // 3.0

	floats := []float64{2.5, 7.5}
	fmt.Println(mathutils.Average(floats))  // 5.0

	empty := []int{}
	fmt.Println(mathutils.Average(empty))  // 0.0
}
```
