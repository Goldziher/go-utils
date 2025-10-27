# RetryWithResult

`func RetryWithResult[T any](fn func() (T, error), maxAttempts int) (T, error)`

RetryWithResult executes a function up to maxAttempts times, returning the first successful result. Returns the result and error from the last attempt if all fail.

```go
package main

import (
	"errors"
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	attempts := 0

	result, err := exc.RetryWithResult(func() (int, error) {
		attempts++
		if attempts < 3 {
			return 0, errors.New("temporary error")
		}
		return 42, nil
	}, 5)

	if err == nil {
		fmt.Printf("Got result: %d after %d attempts\n", result, attempts)
	}
}
```
