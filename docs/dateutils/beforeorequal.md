# BeforeOrEqual

`func BeforeOrEqual(milestone time.Time, date time.Time) bool`

BeforeOrEqual - returns true if a date is before or equal to another date

```go
package main

import (
	"github.com/Goldziher/go-utils/dateutils"
	"time"
)

func main() {
	milestone, _ := time.Parse("2006-01-02", "2023-01-01")

	dBefore, _ := time.Parse("2006-01-02", "2022-12-31")
	dEqual, _ := time.Parse("2006-01-02", "2023-01-01")
	dAfter, _ := time.Parse("2006-01-02", "2023-01-31")

	dateutils.BeforeOrEqual(milestone, dBefore) // true
	dateutils.BeforeOrEqual(milestone, dEqual) // true
	dateutils.BeforeOrEqual(milestone, dAfter) // false
}
```
