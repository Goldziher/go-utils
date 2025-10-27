# ptrutils

Pointer utilities for safe pointer operations and conversions.

## Overview

The `ptrutils` package (imported as `ptr`) provides utilities for working with pointers, including creating pointers to literals, safe dereferencing, and pointer slice operations.

## Functions

**Creation**: To
**Dereferencing**: Deref, DerefSlice
**Comparison**: Equal, IsNil
**Utilities**: Coalesce, ToSlice, NonNilSlice, ToPointerSlice

## Example

```go
import ptr "github.com/Goldziher/go-utils/ptrutils"

// Create pointers to literals
numPtr := ptr.To(42)
strPtr := ptr.To("hello")

// Safe dereferencing with default
value := ptr.Deref(numPtr, 0)  // 42
nilValue := ptr.Deref((*int)(nil), 99)  // 99

// Coalesce - get first non-nil
first := ptr.Coalesce(nil, nil, ptr.To(5), ptr.To(10))  // *5

// Convert slice to pointer slice
values := []int{1, 2, 3}
ptrs := ptr.ToPointerSlice(values)  // []*int
```
