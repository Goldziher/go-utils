# Partition

`func Partition[T any](slice []T, predicate func(T) bool) (truthy []T, falsy []T)`

Partition splits a slice into two slices based on a predicate function. The first slice contains elements for which the predicate returns true, the second contains elements for which it returns false.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Partition into evens and odds
	evens, odds := sliceutils.Partition(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Printf("Evens: %v\n", evens) // [2 4 6 8 10]
	fmt.Printf("Odds: %v\n", odds)   // [1 3 5 7 9]
}
```
