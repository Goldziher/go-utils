# Try

`func Try(fn func() error) (err error)`

Try executes a function and returns its error. This is useful for defer statements or cleanup operations where you want to capture errors. Recovers from panics and converts them to errors.

```go
package main

import (
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	err := exc.Try(func() error {
		// Code that might panic or return error
		return riskyOperation()
	})

	if err != nil {
		fmt.Printf("Caught error: %v\n", err)
	}
}

func riskyOperation() error {
	// ... may return error or panic
	return nil
}
```
