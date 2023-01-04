# PadRight

`func PadRight(str string, padWith string, padTo int) string`

PadRight - Pad a string to a certain length with another string on the right side.
If padding string is more than one char, it might be trucated to fit padTo size.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/stringutils"
)

func main() {
	result := stringutils.PadRight("Azer", "_", 7) // "Azer___"
	fmt.Print(result) // "Azer___"


	result = stringutils.PadRight("Azer", "_-", 7) // "Azer_-_"
	fmt.Print(result) // "Azer_-_"
}
```
