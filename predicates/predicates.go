package predicates

// IsNotEmpty determines whether a given value T has no length.
func IsNotEmpty[T ~map[any]any | ~[]any | ~string | ~chan<- any](value T) bool {
	if len(value) > 0 {
		return true
	}
	return false
}
