# Floor

`func Floor(t time.Time) time.Time`

Floor - takes a datetime and return a datetime from the same day at 00:00:00.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/dateutils"
	"time"
)

func main() {
	d, _ := time.Parse("2006-01-02", "2009-10-13")
	fmt.Println(dateutils.Floor(d)) // 2009-10-13 00:00:00 +0000 UTC
}
```
