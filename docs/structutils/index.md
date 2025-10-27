# structutils

Reflection-based struct utilities with struct tag support.

## Overview

The `structutils` package provides utilities for working with structs through reflection, including conversion to maps, iteration over fields, and struct tag-aware operations.

## Functions

**Conversion**: ToMap
**Iteration**: ForEach
**Inspection**: Fields, Values, FieldNames, HasField, GetField

## Example

```go
import "github.com/Goldziher/go-utils/structutils"

type Config struct {
    Host     string `json:"host" yaml:"host"`
    Port     int    `json:"port" yaml:"port"`
    Database string `json:"database" yaml:"database"`
    Password string `json:"-"`  // Omit from JSON
}

cfg := Config{
    Host:     "localhost",
    Port:     5432,
    Database: "mydb",
    Password: "secret",
}

// Convert to map (respects struct tags)
jsonMap := structutils.ToMap(cfg, "json")
// Result: map[string]any{
//     "host": "localhost",
//     "port": 5432,
//     "database": "mydb",
//     // Password omitted due to json:"-"
// }

yamlMap := structutils.ToMap(cfg, "yaml")
// Result: map[string]any{
//     "host": "localhost",
//     "port": 5432,
//     "database": "mydb",
//     "Password": "secret",  // Included as field name
// }

// Iterate over fields
structutils.ForEach(cfg, func(key string, value any, tag reflect.StructTag) {
    fmt.Printf("%s: %v (json:%s)\n", key, value, tag.Get("json"))
})

// Get field names with tag support
fieldNames := structutils.FieldNames(cfg, "json")
// Result: []string{"host", "port", "database"}
```
