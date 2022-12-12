# PadLeft

`func PadLeft(str string, padWith string, padTo int) string`

PadLeft - Pad a string to a certain length with another string on the left side.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/stringutils"
)

func main() {
	result := stringutils.PadLeft("Azer", "_", 7) // "___Azer"

	fmt.Print(result) // "___Azer"
}
```
