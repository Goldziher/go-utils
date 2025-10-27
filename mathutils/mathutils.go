// This package includes utility functions for mathematical operations.
// It provides helpers for common math operations with generic support.

package mathutils

import "golang.org/x/exp/constraints"

// Number is a constraint for numeric types.
type Number interface {
	constraints.Integer | constraints.Float
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

// Average returns the average (mean) of all values in a slice.
// Returns 0 for empty slice.
func Average[T Number](values []T) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum T
	for _, v := range values {
		sum += v
	}
	return float64(sum) / float64(len(values))
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
