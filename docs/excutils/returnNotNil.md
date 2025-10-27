# ReturnNotNil

`func ReturnNotNil[T any](value *T, messages ...string) *T`

ReturnNotNil panics if the value is nil, otherwise returns the value. This is useful for asserting that a pointer must not be nil. The optional messages are formatted and included in the panic message.

```go
package main

import (
	ptr "github.com/Goldziher/go-utils/ptrutils"
	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	numPtr := ptr.To(42)

	// Returns value if not nil
	result := exc.ReturnNotNil(numPtr, "number must not be nil")

	// Panics if nil
	var nilPtr *int
	exc.ReturnNotNil(nilPtr, "unexpected nil value")  // panics
}
```
