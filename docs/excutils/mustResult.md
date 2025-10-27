# MustResult

`func MustResult[T any](value T, err error, messages ...string) T`

MustResult panics if the error is not nil, otherwise returns the result. This is useful for cases where an error should never occur and you want to fail fast. The optional messages are formatted and included in the panic message.

```go
package main

import (
	"os"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	// Panic if error, otherwise return file
	file := exc.MustResult(os.Open("config.json"), "failed to open config")
	defer file.Close()

	// Use file...
}
```
