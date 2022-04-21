# QueryStringifyStruct

`func QueryStringifyStruct[T interface{}](values T, structTags ...string) string`

Creates a query string from a given struct instance. Takes struct tag names as optional parameters.

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/urlutils"
)

func main() {
	values := struct {
		User    string
		Active  bool
		Age     int `qs:"age"`
		Friends []int
	}{
		User:    "moishe",
		Active:  true,
		Age:     100,
		Friends: []int{1, 2, 3, 4, 5, 6},
	}

	result := urlutils.QueryStringifyStruct(values, "qs")

	fmt.Print(result) // "Active=true&Friends=1&Friends=2&Friends=3&Friends=4&Friends=5&Friends=6&User=moishe&age=100"
}
```

You can pass as many struct tags as you deem necessary, for example the following will also work:

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/urlutils"
)

func main() {
	values := struct {
		User    string
		Active  bool `json:"active"`
		Age     int `qs:"age"`
		Friends []int
	}{
		User:    "moishe",
		Active:  true,
		Age:     100,
		Friends: []int{1, 2, 3, 4, 5, 6},
	}

	result := urlutils.QueryStringifyStruct(values, "json", "qs")

	fmt.Print(result) // "Friends=1&Friends=2&Friends=3&Friends=4&Friends=5&Friends=6&User=moishe&active=true&age=100"
}
```

If you want to ignore a field, simply use the conventional tag value of `"-"`:

```go
package main

import (
	"fmt"

	"github.com/Goldziher/go-utils/urlutils"
)

func main() {
	values := struct {
		User    string
		Active  bool
		Age     int `qs:"-"`
		Friends []int
	}{
		User:    "moishe",
		Active:  true,
		Age:     100,
		Friends: []int{1, 2, 3, 4, 5, 6},
	}

	result := urlutils.QueryStringifyStruct(values, "qs")

	fmt.Print(result) // "Active=true&Friends=1&Friends=2&Friends=3&Friends=4&Friends=5&Friends=6&User=moishe"
}
```
