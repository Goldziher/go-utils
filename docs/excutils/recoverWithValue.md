# RecoverWithValue

`func RecoverWithValue[T any](fn func() T, defaultValue T) (result T)`

RecoverWithValue recovers from a panic and returns the specified default value.

```go
package main

import (
	"fmt"

	exc "github.com/Goldziher/go-utils/excutils"
)

func main() {
	result := exc.RecoverWithValue(func() int {
		panic("something went wrong")
		return 42
	}, 0)

	fmt.Println(result)  // 0 (default value returned due to panic)

	// No panic
	result = exc.RecoverWithValue(func() int {
		return 42
	}, 0)

	fmt.Println(result)  // 42
}
```
