# stringutils

String manipulation and type-safe conversion utilities.

## Overview

The `stringutils` package provides utilities for string manipulation and a powerful type-safe `Stringify` function that converts any value to a string with configurable options.

## Functions

**Type Conversion**: Stringify (with Options)
**Case Conversion**: Capitalize, ToCamelCase, ToSnakeCase, ToKebabCase
**Padding**: PadLeft, PadRight
**Manipulation**: Reverse, Truncate, RemoveWhitespace, EllipsisMiddle
**Utilities**: Contains, SplitAndTrim, JoinNonEmpty, DefaultIfEmpty

## Example

```go
import "github.com/Goldziher/go-utils/stringutils"

// Type-safe stringify with options
str := stringutils.Stringify(42)  // "42"
hex := stringutils.Stringify(42, stringutils.Options{Base: 16})  // "2a"
float := stringutils.Stringify(3.14159, stringutils.Options{Precision: 2})  // "3.14"

// Stringify complex types
m := map[string]int{"a": 1, "b": 2}
str := stringutils.Stringify(m)  // "{a: 1, b: 2}"

// String manipulation
capitalized := stringutils.Capitalize("hello world")  // "Hello world"
padded := stringutils.PadLeft("42", "0", 5)  // "00042"
camel := stringutils.ToCamelCase("hello_world")  // "helloWorld"
snake := stringutils.ToSnakeCase("HelloWorld")  // "hello_world"
kebab := stringutils.ToKebabCase("HelloWorld")  // "hello-world"

// Utilities
truncated := stringutils.Truncate("very long string", 10, "...")  // "very long ..."
cleaned := stringutils.RemoveWhitespace("hello world")  // "helloworld"
parts := stringutils.SplitAndTrim("a, b, c", ",")  // []string{"a", "b", "c"}
```

## Stringify Options

```go
type Options struct {
    Base           int    // Number base for integers (2-36), default 10
    Format         byte   // Float format ('f', 'e', 'E', 'g', 'G'), default 'f'
    Precision      int    // Float precision, default 2
    NilFormat      string // Format for nil values, default "<nil>"
    NilMapFormat   string // Format for nil maps, default "{}"
    NilSliceFormat string // Format for nil slices, default "[]"
}
```
