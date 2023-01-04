# PadLeft

`func PadLeft(str string, padWith string, padTo int) string`

PadLeft - Pad a string to a certain length with another string on the left side.
If padding string is more than one char, it might be trucated to fit padTo size.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/stringutils"
)

func main() {
	result := stringutils.PadLeft("Azer", "_", 7) // "___Azer"
	fmt.Print(result) // "___Azer"

	result = stringutils.PadLeft("Azer", "_-", 7) // "_-_Azer"
	fmt.Print(result) // "_-_Azer"
}
```
