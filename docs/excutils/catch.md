# Catch

`func Catch(fn func() error) (err error)`

Catch executes a function and recovers from any panics, converting them to errors. Returns the function's error if any, or an error created from a recovered panic.

```go
package main

import (
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	err := exc.Catch(func() error {
		// This might panic
		panic("something went wrong")
	})

	fmt.Printf("Caught: %v\n", err)
	// Caught: panic: something went wrong
}
```
