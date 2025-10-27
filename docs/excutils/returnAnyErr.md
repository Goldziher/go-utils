# ReturnAnyErr

`func ReturnAnyErr(values ...any) error`

ReturnAnyErr returns the first error found in the list of values, if any. This is useful when dealing with variadic functions that may return errors.

```go
package main

import (
	"errors"
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	err1 := errors.New("an error")

	// Find any error in mixed values
	err := exc.ReturnAnyErr(42, "hello", err1, true)
	fmt.Println(err)  // an error

	// No errors
	err = exc.ReturnAnyErr(42, "hello", true)
	fmt.Println(err)  // <nil>
}
```
