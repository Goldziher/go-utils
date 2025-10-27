# Capitalize

`func Capitalize(str string) string`

Capitalize - Capitalizes a string by changing the casing format of the first letter of the string.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/stringutils"
)

func main() {
	result := stringutils.Capitalize("coffee") // "Coffee"
	fmt.Print(result) // "Coffee"
}
```
