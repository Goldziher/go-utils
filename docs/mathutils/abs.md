# Abs

`func Abs[T Number](value T) T`

Abs returns the absolute value of a number. For unsigned types, returns the value as-is.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	fmt.Println(mathutils.Abs(5))     // 5
	fmt.Println(mathutils.Abs(-5))    // 5
	fmt.Println(mathutils.Abs(0))     // 0
	fmt.Println(mathutils.Abs(-3.5))  // 3.5
}
```
