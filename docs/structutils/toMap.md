# ToMap

`func ToMap[T any](structInstance T, structTags ...string) map[string]any`

ToMap takes a struct and converts it to into an instance of `map[string]any`.

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/Goldziher/go-utils/structutils"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int `myTag:"myValue"`
}

func main() {
	personInstance := Person{
		FirstName: "Moishe",
		LastName:  "Zuchmir",
		Age:       100,
	}

	personMap := structutils.ToMap(personInstance)

	fmt.Print(personMap)
	// { "FirstName": "Moishe", "LastName": "Zuchmir", "Age": 100 }
}
```

You can also pass in struct tags as an optional argument:

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/Goldziher/go-utils/structutils"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int `myTag:"myValue"`
}

func main() {
	personInstance := Person{
		FirstName: "Moishe",
		LastName:  "Zuchmir",
		Age:       100,
	}

	personMap := structutils.ToMap(personInstance, "myTag")

	fmt.Print(personMap)
	// { "FirstName": "Moishe", "LastName": "Zuchmir", "myTag": 100 }
}
```

To omit a value, use the standard `"-"` struct tag value:

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/Goldziher/go-utils/structutils"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int `myTag:"-"`
}

func main() {
	personInstance := Person{
		FirstName: "Moishe",
		LastName:  "Zuchmir",
		Age:       100,
	}

	personMap := structutils.ToMap(personInstance, "myTag")

	fmt.Print(personMap)
	// { "FirstName": "Moishe", "LastName": "Zuchmir" }
}
```
