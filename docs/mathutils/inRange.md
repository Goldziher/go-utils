# InRange

`func InRange[T cmp.Ordered](value, min, max T) bool`

InRange checks if a value is within a range [min, max] (inclusive).

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	fmt.Println(mathutils.InRange(5, 0, 10))    // true
	fmt.Println(mathutils.InRange(0, 0, 10))    // true (inclusive)
	fmt.Println(mathutils.InRange(10, 0, 10))   // true (inclusive)
	fmt.Println(mathutils.InRange(-1, 0, 10))   // false
	fmt.Println(mathutils.InRange(11, 0, 10))   // false
}
```
