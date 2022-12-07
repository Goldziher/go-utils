# PadRight

`func PadRight(str string, padWith string, padTo int) string`

PadRight - Pad a string to a certain length with another string on the right side.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/stringutils"
)

func main() {
	result := stringutils.PadRight("Azer", "_", 7) // "Azer___"

	fmt.Print(result) // "Azer___"
}
```
