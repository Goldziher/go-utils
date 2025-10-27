# DistinctBy

`func DistinctBy[T any, K comparable](slice []T, keySelector func(T) K) []T`

DistinctBy returns a slice containing only the first occurrence of each element, where uniqueness is determined by the key returned from the keySelector function.

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

	// Get users with distinct ages (first occurrence only)
	distinctByAge := sliceutils.DistinctBy(users, func(u User) int {
		return u.Age
	})

	fmt.Printf("%v\n", distinctByAge)
	// [{Alice 30} {Bob 25}]
}
```
