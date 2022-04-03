# Stringify

`func Stringify(value any, opts ...Options) string`

Stringify receives an arbitrary value and converts it into a string.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/stringutils"
)

func main() {
	value := 1000

	result := stringutils.Stringify(value) // "1000"

	fmt.Print(result) // "1000"
}
```

Stringify also accepts an options object with the following properties:

- `NilFormat`: the string format for nil values, defaults to "<nil>".
- `NilMapFormat`: the string format for nil map objects, defaults to "{}".
- `NilSliceFormat`: the string format for nil slice objects, defaults to "[]".
- `Base`: a number between 2-36 ad the base when converting ints and uints to strings, defaults to Base 10.
- `Precision`: number of digits to include when converting floats and complex numbers to strings, defaults to 2.
- `Format`: the number notation format, using the stlib `FTOA` functionalities, defaults to 'f':
  - 'b' (-ddddp±ddd, a binary exponent),
  - 'e' (-d.dddde±dd, a decimal exponent),
  - 'E' (-d.ddddE±dd, a decimal exponent),
  - 'f' (-ddd.dddd, no exponent),
  - 'g' ('e' for large exponents, 'f' otherwise),
  - 'G' ('E' for large exponents, 'f' otherwise),
  - 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
  - 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
