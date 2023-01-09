# Ceil

`func Ceil(t time.Time) time.Time`

Ceil - takes a datetime and return a datetime from the same day at 23:59:59.

```go
package main

import (
  "fmt"
	"github.com/Goldziher/go-utils/dateutils"
	"time"
)

func main() {
	d, _ := time.Parse("2006-01-02", "2009-10-13")
	fmt.Println(dateutils.Ceil(d)) // 2009-10-13 23:59:59 +0000 UTC
}
```
