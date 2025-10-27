# CountBy

`func CountBy[T any, K comparable](slice []T, keySelector func(T) K) map[K]int`

CountBy counts the occurrences of each key extracted using the keySelector function. Returns a map where keys are the extracted keys and values are the counts.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	words := []string{"apple", "banana", "apricot", "blueberry", "avocado"}

	// Count words by first letter
	counts := sliceutils.CountBy(words, func(word string) rune {
		return rune(word[0])
	})

	fmt.Printf("%v\n", counts)
	// map[97:3 98:2] (a=3, b=2)
}
```
