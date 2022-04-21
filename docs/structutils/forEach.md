# ForEach

`func ForEach[T interface{}](structInstance T, function func(key string, value interface{}, tag reflect.StructTag))`

Takes a struct and calls the passed in function for each of its **visible** fields, passing to in the field's name,
value and tag.

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

	structutils.ForEach(personInstance, func(key string, value interface{}, tag reflect.StructTag) {
		fmt.Printf("%v - %v - %v\n", key, value, tag.Get("myTag"))
	})

	// FirstName - Moishe
	// LastName - Zuchmir
	// Age - 100 - myValue
}
```
