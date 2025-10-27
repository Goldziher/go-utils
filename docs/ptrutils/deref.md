# Deref

`func Deref[T any](ptr *T, def T) T`

Deref dereferences ptr and returns the value it points to if not nil, or else returns def (the default value). This provides safe pointer dereferencing without nil checks.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	numPtr := ptr.To(42)
	value := ptr.Deref(numPtr, 0)
	fmt.Println(value)  // 42

	// Safe handling of nil pointers
	var nilPtr *int
	defaultValue := ptr.Deref(nilPtr, 99)
	fmt.Println(defaultValue)  // 99
}
```
