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
