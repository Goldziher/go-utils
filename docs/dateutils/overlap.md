# Overlap

`Overlap(start1 time.Time, end1 time.Time, start2 time.Time, end2 time.Time) bool`

Overlap - returns true if two date intervals overlap.

```go
package main

import (
	"github.com/Goldziher/go-utils/dateutils"
	"time"
)

func main() {
	s1, _ := time.Parse("2006-01-02", "2022-12-28")
	e1, _ := time.Parse("2006-01-02", "2022-12-31")

	s2, _ := time.Parse("2006-01-02", "2022-12-30")
	e2, _ := time.Parse("2006-01-02", "2023-01-01")

	s3, _ := time.Parse("2006-01-02", "2023-01-02")
	e3, _ := time.Parse("2006-01-02", "2023-01-04")

	dateutils.Overlap(s1, e1, s2, e2) // true
	dateutils.Overlap(s1, e1, s3, e3) // false
}
```
