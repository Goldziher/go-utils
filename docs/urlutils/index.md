# urlutils

URL parsing and query string builder utilities.

## Overview

The `urlutils` package provides utilities for working with URLs and building query strings from maps and structs with struct tag support.

## Functions

**Query Builders**: QueryStringifyMap, QueryStringifyStruct
**URL Parsing**: Parse, MustParse
**URL Inspection**: IsAbsolute, GetDomain, GetScheme, GetPath

## Example

```go
import "github.com/Goldziher/go-utils/urlutils"

// Build query string from map
params := map[string]any{
    "search": "golang",
    "page":   1,
    "tags":   []string{"generics", "utils"},
    "active": true,
}
query := urlutils.QueryStringifyMap(params)
// Result: "active=true&page=1&search=golang&tags=generics&tags=utils"

// Build query string from struct with tags
type SearchQuery struct {
    Query    string   `qs:"q"`
    Page     int      `qs:"page"`
    PageSize int      `qs:"page_size"`
    Tags     []string `qs:"tag"`
    Active   bool     `qs:"active"`
}

q := SearchQuery{
    Query:    "golang",
    Page:     1,
    PageSize: 20,
    Tags:     []string{"generics", "utils"},
    Active:   true,
}
query := urlutils.QueryStringifyStruct(q, "qs")
// Result: "active=true&page=1&page_size=20&q=golang&tag=generics&tag=utils"

// URL parsing and inspection
u, err := urlutils.Parse("https://example.com:8080/path?key=value")
if err == nil {
    domain := urlutils.GetDomain("https://example.com:8080/path")  // "example.com:8080"
    scheme := urlutils.GetScheme("https://example.com/path")       // "https"
    path := urlutils.GetPath("https://example.com/path/to/page")   // "/path/to/page"
    isAbs := urlutils.IsAbsolute("https://example.com")            // true
}

// Panic on parse error (use when URL is known to be valid)
u := urlutils.MustParse("https://example.com/path")
```
