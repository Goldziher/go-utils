# FirstErr

`func FirstErr(errs ...error) error`

FirstErr returns the first non-nil error from a list of errors. Returns nil if all errors are nil.

```go
package main

import (
	"errors"
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	err1 := errors.New("first error")
	err2 := errors.New("second error")

	// Get first non-nil error
	err := exc.FirstErr(nil, nil, err1, err2)
	fmt.Println(err)  // first error

	// All nil
	err = exc.FirstErr(nil, nil, nil)
	fmt.Println(err)  // <nil>
}
```
