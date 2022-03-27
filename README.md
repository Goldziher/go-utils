# Go Utils

<div align="center">

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)

[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Goldziher_go-utils&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Goldziher_go-utils)

</div>

Go Utils is a utility library inspired by JavaScript and Python. It makes extensive use of Go 1.18+ generics to offer a
straightforward API.

## Installation

```shell
go get github.com/Goldziher/go-utils
```

## Slice Utilities

### Filter

`func Filter[T any](slice []T, predicate func(value T, index int, slice []T) bool) []T`

Filter takes a slice of type `T` and filters it using the given predicate function. The predicate is passed the current
element, the current index and the slice as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddNumbers := sliceutils.Filter(numerals, func(value int, index int, slice []int) bool {
		return value%2 != 0
	})

	fmt.Printf("%v", oddNumbers) // [1 3 5 7 9]
}
```

### ForEach

`func ForEach[T any](slice []T, function func(value T, index int, slice []T))`

ForEach executes the passed in function for each element in the given slice. The function is passed the current element,
the current index and the slice as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	result := 0
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceutils.ForEach(numerals, func(value int, index int, slice []int) {
		result += value
	})

	fmt.Print(result) // 45
}
```

### Map

`func Map[T any, R any](slice []T, mapper func(value T, index int, slice []T) R) (mapped []R)`

Map allows transforming the values in a slice by executing the given mapper function for each element in the slice. The
function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := sliceutils.Map(numerals, func(value int, index int, slice []int) int {
		return value * 2
	})

	fmt.Print(result) // [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
}
```

### Reduce

`func Reduce[T any, R any](slice []T, reducer func(acc R, value T, index int, slice []T) R, initial R) R`

Reduce allows transforming the slice and its values into a different value by executing the given reducer function for
each element in the slice. The function is passed the accumulator, current element, current index and the slice as
function arguments. The third argument to reduce is the initial value.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := sliceutils.Reduce(
		numerals,
		func(acc int, cur int, index int, slice []int) int {
			return acc + cur
		},
		0,
	)

	fmt.Print(sum) // 45
}
```

### Find

`func Find[T any](slice []T, predicate func(value T, index int, slice []T) bool) *T`

Find takes a slice of type T and executes the passed in predicate function for each element in the slice. If the
predicate returns true, a pointer to the element is returned. If no element is found, nil is returned. The function is
passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	result := sliceutils.Find(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})

	fmt.Print(result) // "Wednesday"
}
```

### FindIndex

`func FindIndex[T any](slice []T, predicate func(value T, index int, slice []T) bool) int`

FindIndex takes a slice of type T and executes the passed in predicate function for each element in the slice. If the
predicate returns true, the element's index is returned. If no element is found, `-1` is returned. The function is
passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	result := sliceutils.FindIndex(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})

	fmt.Print(result) // 3
}
```

### FindIndexOf

`func FindIndexOf[T comparable](slice []T, value T) int`

FindIndexOf takes a slice of type T and a value of type T. If any element in the slice equals the given value, its index
is returned. If no element is found, `-1` is returned.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	result := sliceutils.FindIndexOf(days, "Wednesday")

	fmt.Print(result) // 3
}
```

### FindLastIndex

`func FindLastIndex[T any](slice []T, predicate func(value T, index int, slice []T) bool) int`

FindLastIndex takes a slice of type T and executes the passed in predicate function for each element in the slice
starting from its end. If the predicate returns true, the element's index is returned. If no element is found, `-1` is
returned. The function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindLastIndex(friends, func(value string, index int, slice []string) bool {
		return value == "John"
	})

	fmt.Print(result) // 4
}
```

### FindLastIndexOf

`func FindLastIndexOf[T comparable](slice []T, value T) int`

FindLastIndexOf takes a slice of type T and a value of type T. If any element in the slice equals the given value, the
last index of is occurrence is returned. If no element is found, `-1` is returned.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindLastIndexOf(friends, "John")

	fmt.Print(result) // 4
}
```

### FindIndexes

`func FindIndexes[T any](slice []T, predicate func(value T, index int, slice []T) bool) []int`

FindIndexes takes a slice of type T and executes the passed in predicate function for each element in the slice,
returning a slice containing all indexes for which the predicate returned true. The function is passed the current
element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindIndexes(friends, func(value string, index int, slice []string) bool {
		return value == "John"
	})

	fmt.Print(result) // [0, 4]
}
```

### FindIndexesOf

`func FindIndexes[T any](slice []T, predicate func(value T, index int, slice []T) bool) []int`

FindIndexesOf takes a slice of type T and a value of type T, returning a slice containing all indexes where elements
equal the passed in value.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.FindIndexesOf(friends, "John")

	fmt.Print(result) // [0, 4]
}
```

### Includes

`func Includes[T comparable](slice []T, value T) bool`

Includes takes a slice of type T and a value of type T, determining whether or not the value is included in the slice.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Includes(friends, "John")

	fmt.Print(result) // true
}
```

### Some

`func Some[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool`

Some takes a slice of type T and a predicate function, returning true if the predicate returned true for **some** elements.
The function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Some(friends, func(value string, index int, slice []string) bool {
        return value == "Mandy"
	})

	fmt.Print(result) // true
}
```

### Every

`func Every[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool`

Every takes a slice of type T and a predicate function, returning true if the predicate returned true for **every** elements.
The function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	var friends = []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Every(friends, func(value string, index int, slice []string) bool {
        return value == "Mandy"
	})

	fmt.Print(result) // false
}
```
