# MustNotNil

`func MustNotNil[T any](value *T, messages ...string) T`

MustNotNil panics with a formatted error message if the value is nil. Returns the dereferenced value if not nil.

```go
package main

import (
	ptr "github.com/Goldziher/go-utils/ptrutils"
	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	numPtr := ptr.To(42)

	// Returns dereferenced value if not nil
	value := exc.MustNotNil(numPtr, "number must not be nil")
	println(value)  // 42

	// Panics if nil
	var nilPtr *int
	exc.MustNotNil(nilPtr, "unexpected nil value")  // panics
}
```
