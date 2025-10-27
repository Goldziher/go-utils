# Retry

`func Retry(fn func() error, maxAttempts int) error`

Retry executes a function up to maxAttempts times, returning the first successful result. Returns the last error if all attempts fail. Returns an error if maxAttempts is less than 1.

```go
package main

import (
	"errors"
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	attempts := 0

	err := exc.Retry(func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	}, 5)

	if err == nil {
		fmt.Printf("Succeeded after %d attempts\n", attempts)
	}
}
```
