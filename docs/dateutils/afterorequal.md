# AfterOrEqual

`func AfterOrEqual(milestone time.Time, date time.Time) bool`

AfterOrEqual returns true if a date is equal to or after another date.

Deprecated: prefer `!date.Before(milestone)` from the standard library.

```go
package main

import (
	"time"

	"github.com/Goldziher/go-utils/dateutils"
)

func main() {
	milestone, _ := time.Parse("2006-01-02", "2023-01-01")

	dBefore, _ := time.Parse("2006-01-02", "2022-12-31")
	dEqual, _ := time.Parse("2006-01-02", "2023-01-01")
	dAfter, _ := time.Parse("2006-01-02", "2023-01-31")

	dateutils.AfterOrEqual(milestone, dBefore) // false
	dateutils.AfterOrEqual(milestone, dEqual)  // true
	dateutils.AfterOrEqual(milestone, dAfter)  // true
}
```
