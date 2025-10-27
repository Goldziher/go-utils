# IsPrime

`func IsPrime[T constraints.Integer](n T) bool`

IsPrime checks if a number is prime. Returns false for numbers less than 2.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/mathutils"
)

func main() {
	// Prime numbers
	fmt.Println(mathutils.IsPrime(2))   // true
	fmt.Println(mathutils.IsPrime(3))   // true
	fmt.Println(mathutils.IsPrime(17))  // true
	fmt.Println(mathutils.IsPrime(19))  // true

	// Non-prime numbers
	fmt.Println(mathutils.IsPrime(0))   // false
	fmt.Println(mathutils.IsPrime(1))   // false
	fmt.Println(mathutils.IsPrime(4))   // false
	fmt.Println(mathutils.IsPrime(15))  // false
}
```
