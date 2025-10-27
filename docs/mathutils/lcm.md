# Lcm

`func Lcm[T constraints.Integer](a, b T) T`

Lcm returns the least common multiple of two integers.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	fmt.Println(mathutils.Lcm(4, 6))    // 12
	fmt.Println(mathutils.Lcm(17, 13))  // 221
	fmt.Println(mathutils.Lcm(10, 5))   // 10
	fmt.Println(mathutils.Lcm(0, 5))    // 0
	fmt.Println(mathutils.Lcm(5, 0))    // 0
}
```
