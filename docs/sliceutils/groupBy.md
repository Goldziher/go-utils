# GroupBy

`func GroupBy[T any, K comparable](slice []T, keySelector func(T) K) map[K][]T`

GroupBy groups elements of a slice by a key extracted using the keySelector function. Returns a map where keys are the grouping keys and values are slices of elements that share that key. This is a LINQ-style operation useful for categorizing or organizing data.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 25},
	}

	// Group users by age
	byAge := sliceutils.GroupBy(users, func(u User) int {
		return u.Age
	})

	fmt.Printf("%v\n", byAge)
	// map[25:[{Bob 25} {David 25}] 30:[{Alice 30} {Charlie 30}]]
}
```
