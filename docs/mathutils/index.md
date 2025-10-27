# mathutils

Generic math operations with type constraints.

## Overview

The `mathutils` package provides mathematical utilities that work with generic numeric types through Go generics and constraints. All functions use appropriate type constraints (`cmp.Ordered`, `constraints.Integer`, etc.) for compile-time type safety.

## Functions

**Range Operations**: Clamp, InRange
**Numeric Operations**: Abs, Average
**Number Theory**: Gcd, Lcm, IsPrime

## Example

```go
import "github.com/Goldziher/go-utils/mathutils"

// Clamp values to a range
clamped := mathutils.Clamp(150, 0, 100)  // 100

// Check if value is in range
inRange := mathutils.InRange(50, 0, 100)  // true

// Number theory
gcd := mathutils.Gcd(48, 18)   // 6
lcm := mathutils.Lcm(4, 6)     // 12
prime := mathutils.IsPrime(17)  // true

// Statistics
values := []int{1, 2, 3, 4, 5}
avg := mathutils.Average(values)  // 3.0
```
