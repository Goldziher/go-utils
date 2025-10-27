// This package includes utility functions for working with pointers.
// It provides helpers for creating pointers and safely dereferencing them.

package ptr

// To returns a pointer to the given value.
// This is useful for creating pointers to literals or values in a single expression.
func To[T any](v T) *T {
	return &v
}

// Deref dereferences ptr and returns the value it points to if not nil, or else
// returns def (the default value).
// This provides safe pointer dereferencing without nil checks.
func Deref[T any](ptr *T, def T) T {
	if ptr != nil {
		return *ptr
	}
	return def
}

// Equal compares two pointers for equality.
// Returns true if both are nil or both point to equal values.
// Returns false if one is nil and the other is not.
func Equal[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// Coalesce returns the first non-nil pointer from the list.
// Returns nil if all pointers are nil.
func Coalesce[T any](ptrs ...*T) *T {
	for _, ptr := range ptrs {
		if ptr != nil {
			return ptr
		}
	}
	return nil
}

// IsNil checks if a pointer is nil.
// This is a convenience function for readability.
func IsNil[T any](ptr *T) bool {
	return ptr == nil
}

// ToSlice converts a pointer to a slice.
// Returns an empty slice if ptr is nil, otherwise returns a slice with the single element.
func ToSlice[T any](ptr *T) []T {
	if ptr == nil {
		return []T{}
	}
	return []T{*ptr}
}

// DerefSlice dereferences all pointers in a slice, using def for nil pointers.
// Returns a new slice with dereferenced values.
func DerefSlice[T any](ptrs []*T, def T) []T {
	result := make([]T, len(ptrs))
	for i, ptr := range ptrs {
		result[i] = Deref(ptr, def)
	}
	return result
}

// NonNilSlice returns a new slice containing only non-nil pointers from the input.
func NonNilSlice[T any](ptrs []*T) []*T {
	result := make([]*T, 0, len(ptrs))
	for _, ptr := range ptrs {
		if ptr != nil {
			result = append(result, ptr)
		}
	}
	return result
}

// ToPointerSlice converts a slice of values to a slice of pointers.
// Each element in the result points to the corresponding element in the input.
func ToPointerSlice[T any](values []T) []*T {
	result := make([]*T, len(values))
	for i := range values {
		result[i] = &values[i]
	}
	return result
}
