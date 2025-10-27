# Must

`func Must(err error, messages ...string)`

Must panics if the error is not nil. This is useful for initialization or setup code where errors should never occur. The optional messages are formatted and included in the panic message.

```go
package main

import (
	"os"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	// Panic if error occurs
	err := os.Chdir("/tmp")
	exc.Must(err, "failed to change directory")

	// Continue if no error
	println("Successfully changed directory")
}
```
