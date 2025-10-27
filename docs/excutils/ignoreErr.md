# IgnoreErr

`func IgnoreErr(fn func() error)`

IgnoreErr executes a function that returns an error and ignores the error. This is useful for cleanup operations where errors can be safely ignored.

```go
package main

import (
	"os"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	file, _ := os.Open("temp.txt")
	defer exc.IgnoreErr(file.Close)

	// Use file...
	// Close error will be ignored
}
```
