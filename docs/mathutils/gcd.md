# Gcd

`func Gcd[T constraints.Integer](a, b T) T`

Gcd returns the greatest common divisor of two integers using the Euclidean algorithm.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	fmt.Println(mathutils.Gcd(48, 18))   // 6
	fmt.Println(mathutils.Gcd(17, 13))   // 1
	fmt.Println(mathutils.Gcd(10, 5))    // 5
	fmt.Println(mathutils.Gcd(-12, 24))  // 12 (always returns positive)
	fmt.Println(mathutils.Gcd(7, 0))     // 7
}
```
