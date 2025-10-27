# To

`func To[T any](v T) *T`

To returns a pointer to the given value. This is useful for creating pointers to literals or values in a single expression.

```go
package main

import (
	"fmt"

	ptr "github.com/Goldziher/go-utils/ptrutils"
)

func main() {
	// Create pointers to literals
	numPtr := ptr.To(42)
	strPtr := ptr.To("hello")
	boolPtr := ptr.To(true)

	fmt.Printf("%v, %v, %v\n", *numPtr, *strPtr, *boolPtr)
	// 42, hello, true

	// Useful for struct fields that require pointers
	type Config struct {
		MaxRetries *int
	}
	config := Config{MaxRetries: ptr.To(5)}
}
```
