# FindIndexOf

`func FindIndexOf[T comparable](slice []T, value T) int`

FindIndexOf takes a slice of type T and a value of type T. If any element in the slice equals the given value, its index
is returned. If no element is found, `-1` is returned.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	result := sliceutils.FindIndexOf(days, "Wednesday")

	fmt.Print(result) // 3
}
```
