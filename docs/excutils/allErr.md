# AllErr

`func AllErr(errs ...error) error`

AllErr returns an error containing all non-nil errors from the list. Returns nil if all errors are nil. Multiple errors are joined using errors.Join (Go 1.20+).

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

	// Join all non-nil errors
	err := exc.AllErr(nil, err1, nil, err2)
	fmt.Println(err)
	// first error
	// second error
}
```
