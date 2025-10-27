// This package includes utility functions for mathematical operations.
// It provides helpers for common math operations with generic support.

package mathutils

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Number is a constraint for numeric types.
type Number interface {
	constraints.Integer | constraints.Float
}

// Min returns the minimum of two values.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// MinN returns the minimum value from a slice.
// Returns zero value if slice is empty.
func MinN[T constraints.Ordered](values []T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}

	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

// MaxN returns the maximum value from a slice.
// Returns zero value if slice is empty.
func MaxN[T constraints.Ordered](values []T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}

	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

// Clamp restricts a value to be within a specified range.
// If value < min, returns min. If value > max, returns max. Otherwise returns value.
func Clamp[T constraints.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Abs returns the absolute value of an integer.
// For unsigned types, returns the value as-is.
func Abs[T Number](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

// InRange checks if a value is within a range [min, max] (inclusive).
func InRange[T constraints.Ordered](value, min, max T) bool {
	return value >= min && value <= max
}

// Sum returns the sum of all values in a slice.
func Sum[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// Product returns the product of all values in a slice.
// Returns 1 for empty slice.
func Product[T Number](values []T) T {
	if len(values) == 0 {
		return 1
	}

	product := values[0]
	for i := 1; i < len(values); i++ {
		product *= values[i]
	}
	return product
}

// Average returns the average (mean) of all values in a slice.
// Returns 0 for empty slice.
func Average[T Number](values []T) float64 {
	if len(values) == 0 {
		return 0
	}
	return float64(Sum(values)) / float64(len(values))
}

// Round rounds a float to the nearest integer.
func Round(value float64) float64 {
	return math.Round(value)
}

// RoundToDecimal rounds a float to a specified number of decimal places.
func RoundToDecimal(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}

// Floor rounds a float down to the nearest integer.
func Floor(value float64) float64 {
	return math.Floor(value)
}

// Ceil rounds a float up to the nearest integer.
func Ceil(value float64) float64 {
	return math.Ceil(value)
}

// Gcd returns the greatest common divisor of two integers using Euclidean algorithm.
func Gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs(a)
}

// Lcm returns the least common multiple of two integers.
func Lcm[T constraints.Integer](a, b T) T {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / Gcd(a, b)
}

// Pow returns base raised to the power of exponent for integers.
func Pow[T constraints.Integer](base, exponent T) T {
	if exponent == 0 {
		return 1
	}

	result := base
	var i T = 1
	for i < exponent {
		result *= base
		i++
	}
	return result
}

// IsPrime checks if a number is prime.
// Returns false for numbers less than 2.
func IsPrime[T constraints.Integer](n T) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Check odd divisors up to sqrt(n)
	var i T = 3
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += 2
	}
	return true
}

// IsEven checks if a number is even.
func IsEven[T constraints.Integer](n T) bool {
	return n%2 == 0
}

// IsOdd checks if a number is odd.
func IsOdd[T constraints.Integer](n T) bool {
	return n%2 != 0
}

// Sign returns the sign of a number: -1 for negative, 0 for zero, 1 for positive.
func Sign[T Number](value T) int {
	if value < 0 {
		return -1
	}
	if value > 0 {
		return 1
	}
	return 0
}
