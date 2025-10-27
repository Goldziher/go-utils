# GroupBy

`func GroupBy[K comparable, V any, G comparable](mapInstance map[K]V, grouper func(key K, value V) G) map[G]map[K]V`

GroupBy groups map entries by a grouping key generated from each entry. Returns a map where keys are group identifiers and values are maps of entries belonging to that group.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/maputils"
)

func main() {
	scores := map[string]int{
		"Alice":   95,
		"Bob":     67,
		"Charlie": 88,
		"David":   72,
		"Eve":     91,
	}

	// Group by grade (A: 90+, B: 80-89, C: 70-79, etc.)
	byGrade := maputils.GroupBy(scores, func(name string, score int) string {
		switch {
		case score >= 90:
			return "A"
		case score >= 80:
			return "B"
		case score >= 70:
			return "C"
		default:
			return "F"
		}
	})

	fmt.Printf("%v\n", byGrade)
	// map[A:map[Alice:95 Eve:91] B:map[Charlie:88] C:map[Bob:67 David:72]]
}
```
