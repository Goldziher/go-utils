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
go get -u github.com/Goldziher/go-utils
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
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

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
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

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
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

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
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

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
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

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
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

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
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

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
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

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
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

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
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

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
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

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
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Includes(friends, "John")

	fmt.Print(result) // true
}
```

### Some

`func Some[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool`

Some takes a slice of type T and a predicate function, returning true if the predicate returned true for **some**
elements. The function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Some(friends, func(value string, index int, slice []string) bool {
		return value == "Mandy"
	})

	fmt.Print(result) // true
}
```

### Every

`func Every[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool`

Every takes a slice of type T and a predicate function, returning true if the predicate returned true for **every**
elements. The function is passed the current element, the current index and the slice itself as function arguments.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	friends := []string{"John", "Bob", "Mendy", "Suzy", "John"}

	result := sliceutils.Every(friends, func(value string, index int, slice []string) bool {
		return value == "Mandy"
	})

	fmt.Print(result) // false
}
```

### Merge

`func Merge[T any](slices ...[]T) (mergedSlice []T)`

Merge takes slices of type T and merges them into a single slice of type T, preserving their order.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}

	result := sliceutils.Merge(first, second, third)

	fmt.Print(result) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
}
```

### Sum

`func Sum[T numbers](slice []T) (result T)`

Sum takes a slice of numbers T, which can be any of the number types, and returns a sum of their values.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := sliceutils.Sum(numerals)

	fmt.Print(result) // 45
}
```

### Remove

`func Remove[T any](slice []T, i int) []T`

Remove takes a slice of type T and an index, removing the element at the given index.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	numerals = sliceutils.Remove(numerals, 3)

	fmt.Print(numerals)   // [0, 1, 2, 4, 5, 6, 7, 8, 9]
}
```

### Insert

`func Insert[T any](slice []T, i int, value T) []T`

Insert takes a slice of type T, an index and a value of type T, inserting the value at the given index and shifting any
existing elements to the right.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 2}

	numerals = sliceutils.Insert(numerals, 1, 1)

	fmt.Print(numerals)   // [0, 1, 2]
}
```

### Copy

`func Copy[T any](slice []T) []T`

Copy receives a slice of type T and returns a copy.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1}

	numeralsCopy := sliceutils.Copy(numerals)
	numeralsCopy[0] = 1

	fmt.Print(numerals)   // [0, 1]
	fmt.Print(numeralsCopy)   // [1, 1]
}
```

### Intersection

`func Intersection[T comparable](slices ...[]T) []T`

Intersection takes a variadic number of slices of type T and returns a slice of type T containing any values that are common to all slices.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first, second := []int{1, 2, 3}, []int{1, 7, 3}

	result := sliceutils.Intersection(first, second)

	fmt.Print(result)   // [1, 3]
}
```

### Difference

`func Difference[T comparable](slices ...[]T) []T`

Difference takes a variadic number of slices of type T and returns a slice of type T containing the elements that are different between the slices.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first, second := []int{1, 2, 3}, []int{1, 7, 3}

	result := sliceutils.Difference(first, second)

	fmt.Print(result)   // [2, 7]
}
```

### Union

`func Union[T comparable](slices ...[]T) []T`

Union takes a variadic number of slices of type T and returns a slice of type T containing the unique elements in the different slices.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	first, second := []int{1, 2, 3}, []int{1, 7, 3}

	result := sliceutils.Union(first, second)

	fmt.Print(result)   // [1, 2, 3, 7]
}
```

### Reverse

`func Reverse[T any](slice []T) []T`

Reverse takes a slice of type T and returns a slice of type T with a reverse order of elements.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	reversed := sliceutils.Reverse(numerals)

	fmt.Print(reversed)   // [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
}
```

### Unique

`func Unique[T comparable](slice []T) []T`

Unique tales a slice of type T and returns a slice of type T containing all unique elements.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/sliceutils"
)

func main() {
	numerals := []int{0, 1, 2, 3, 3, 1}

	result := sliceutils.Unique(numerals)

	fmt.Print(result)   // [0, 1, 2, 3]
}
```
