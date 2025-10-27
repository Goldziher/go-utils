# Clamp

`func Clamp[T cmp.Ordered](value, min, max T) T`

Clamp restricts a value to be within a specified range. If value < min, returns min. If value > max, returns max. Otherwise returns value.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	// Clamp integers
	fmt.Println(mathutils.Clamp(5, 0, 10))   // 5
	fmt.Println(mathutils.Clamp(-5, 0, 10))  // 0
	fmt.Println(mathutils.Clamp(15, 0, 10))  // 10

	// Clamp floats
	fmt.Println(mathutils.Clamp(2.5, 0.0, 5.0))  // 2.5
	fmt.Println(mathutils.Clamp(6.0, 0.0, 5.0))  // 5.0
}
```
