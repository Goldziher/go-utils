// This package includes utility functions for handling and manipulating a slice.
// It draws inspiration from JavaScript and Python and uses Go generics as a basis.

package sliceutils

// Filter - given a slice of type T, executes the given predicate function on each element in the slice.
// The predicate is passed the current element, the current index and the slice itself as function arguments.
// If the predicate returns true, the value is included in the result, otherwise it is filtered out.
func Filter[T any](slice []T, predicate func(value T, index int, slice []T) bool) (filtered []T) {
	for i, el := range slice {
		if ok := predicate(el, i, slice); ok {
			filtered = append(filtered, el)
		}
	}
	return filtered
}

// ForEach - given a slice of type T, executes the passed in function for each element in the slice.
// The function is passed the current element, the current index and the slice itself as function arguments.
func ForEach[T any](slice []T, function func(value T, index int, slice []T)) {
	for i, el := range slice {
		function(el, i, slice)
	}
}

// Map - given a slice of type T, executes the passed in mapper function for each element in the slice, returning a slice of type R.
// The function is passed the current element, the current index and the slice itself as function arguments.
func Map[T any, R any](slice []T, mapper func(value T, index int, slice []T) R) (mapped []R) {
	for i, el := range slice {
		mapped = append(mapped, mapper(el, i, slice))
	}
	return mapped
}

// Reduce - given a slice of type T, executes the passed in reducer function for each element in the slice, returning a result of type R.
// The function is passed the accumulator, current element, current index and the slice itself as function arguments.
// The third argument to reduce is the initial value of type R to use.
func Reduce[T any, R any](slice []T, reducer func(acc R, value T, index int, slice []T) R, initial R) R {
	acc := initial
	for i, el := range slice {
		acc = reducer(acc, el, i, slice)
	}
	return acc
}

// Find - given a slice of type T, executes the passed in predicate function for each element in the slice.
// If the predicate returns true - a pointer to the element is returned. If no element is found, nil is returned.
// The function is passed the current element, the current index and the slice itself as function arguments.
func Find[T any](slice []T, predicate func(value T, index int, slice []T) bool) *T {
	for i, el := range slice {
		if ok := predicate(el, i, slice); ok {
			return &el
		}
	}
	return nil
}

// FindIndex - given a slice of type T, executes the passed in predicate function for each element in the slice.
// If the predicate returns true - the index of the element is returned. If no element is found, -1 is returned.
// The function is passed the current element, the current index and the slice itself as function arguments.
func FindIndex[T any](slice []T, predicate func(value T, index int, slice []T) bool) int {
	for i, el := range slice {
		if ok := predicate(el, i, slice); ok {
			return i
		}
	}
	return -1
}

// FindIndexOf - given a slice of type T and a value of type T, return ths first index of an element equal to value.
// If no element is found, -1 is returned.
func FindIndexOf[T comparable](slice []T, value T) int {
	for i, el := range slice {
		if el == value {
			return i
		}
	}
	return -1
}

// FindLastIndex - given a slice of type T, executes the passed in predicate function for each element in the slice starting from its end.
// If no element is found, -1 is returned. The function is passed the current element, the current index and the slice itself as function arguments.
func FindLastIndex[T any](slice []T, predicate func(value T, index int, slice []T) bool) int {
	for i := len(slice) - 1; i > 0; i-- {
		el := slice[i]
		if ok := predicate(el, i, slice); ok {
			return i
		}
	}
	return -1
}

// FindLastIndexOf - given a slice of type T and a value of type T, returning the last index matching the given value.
// If no element is found, -1 is returned.
func FindLastIndexOf[T comparable](slice []T, value T) int {
	for i := len(slice) - 1; i > 0; i-- {
		el := slice[i]
		if el == value {
			return i
		}
	}
	return -1
}

// FindIndexes - given a slice of type T, executes the passed in predicate function for each element in the slice.
// Returns a slice containing all indexes of elements for which the predicate returned true. If no element are found, returns a nil int slice.
// The function is passed the current element, the current index and the slice itself as function arguments.
func FindIndexes[T any](slice []T, predicate func(value T, index int, slice []T) bool) []int {
	var indexes []int
	for i, el := range slice {
		if ok := predicate(el, i, slice); ok {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// FindIndexesOf - given a slice of type T and a value of type T, returns a slice containing all indexes match the given value.
// If no element are found, returns a nil int slice.
func FindIndexesOf[T comparable](slice []T, value T) []int {
	var indexes []int
	for i, el := range slice {
		if el == value {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Includes - given a slice of type T and a value of type T, determines whether the value is contained by the slice.
// Note: T is constrained to comparable types only and comparison is determined using the equality operator.
func Includes[T comparable](slice []T, value T) bool {
	for _, el := range slice {
		if el == value {
			return true
		}
	}
	return false
}

// Some - given a slice of type T, executes the given predicate for each element of the slice.
// If the predicate returns true for any element, it returns true, otherwise it returns false.
// The function is passed the current element, the current index and the slice itself as function arguments.
func Some[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool {
	for i, el := range slice {
		if ok := predicate(el, i, slice); ok {
			return true
		}
	}
	return false
}

// Every - given a slice of type T, executes the given predicate for each element of the slice.
// If the predicate returns true for all elements, it returns true, otherwise it returns false.
// The function is passed the current element, the current index and the slice itself as function arguments.
func Every[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool {
	for i, el := range slice {
		if ok := predicate(el, i, slice); !ok {
			return false
		}
	}
	return true
}

// Merge - receives slices of type T and merges them into a single slice of type T.
// Note: The elements are merged in their order in the a slice,
// i.e. first the elements of the first slice, then that of the second and so forth.
func Merge[T any](slices ...[]T) (mergedSlice []T) {
	for _, slice := range slices {
		for _, el := range slice {
			mergedSlice = append(mergedSlice, el)
		}
	}
	return mergedSlice
}

// numbers - an interface for all number types.
type numbers interface {
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | complex64 | complex128
}

// Sum - receives a slice of type T and returns a value T that is the sum of the numbers.
// Note: T is constrained to be a number type.
func Sum[T numbers](slice []T) (result T) {
	for _, el := range slice {
		result += el
	}
	return result
}

// Remove - receives a slice of type T and an index, removing the element at the given index.
// Note: this function returns a copy of the passed in slice and does not modify it in place.
func Remove[T any](slice []T, i int) []T {
	result := make([]T, len(slice), cap(slice))
	copy(result, slice)

	if len(result) == 0 {
		return result
	} else if i != len(result)-1 {
		return append(result[:i], result[i+1:]...)
	}
	return result[:i]
}

// Insert - receives a slice of type T, an index and a value.
// The value is inserted at the given index. If there is an existing value at this index, it is shifted to the next index.
func Insert[T any](slice []T, i int, value T) []T {
	if len(slice) == i {
		return append(slice, value)
	}
	slice = append(slice[:i+1], slice[i:]...)
	slice[i] = value
	return slice
}

// Intersection - receives a slice of type T and returns a slice of type T containing any values that exist in all the a slice.
// For example, given []int{1, 2, 3}, []{1, 7, 3}, the intersection would be []int{1, 3}.
func Intersection[T comparable](slices ...[]T) []T {
	return Filter(Merge(slices...), func(value T, i int, slice []T) bool {
		indexes := FindIndexesOf(slice, value)
		return len(indexes) == len(slices) && indexes[0] == i
	})
}

// Difference - receives a slice of type T, and returns the difference between the a slice.
// For example, given []int{1, 2, 3}, []{2, 3, 4}, []{3, 4, 5}, the difference would be []int{1, 5}.
func Difference[T comparable](slices ...[]T) []T {
	return Filter(Merge(slices...), func(value T, _ int, slice []T) bool {
		indexes := FindIndexesOf(slice, value)
		return len(indexes) == 1
	})
}

// Union - receives slices of type T, and returns a slice composed of the unique elements of the slices.
// For example, given []int{1, 2, 3}, []{2, 3, 4}, []{3, 4, 5}, the difference would be []int{1, 2, 3, 4, 5}.
func Union[T comparable](slices ...[]T) []T {
	return Filter(Merge(slices...), func(value T, i int, slice []T) bool {
		indexes := FindIndexesOf(slice, value)
		return indexes[0] == i
	})
}

// Reverse - receives a slice of type T and reverses it, returning a slice of type T with a reverse order of elements.
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	itemCount := len(slice)
	middle := itemCount / 2
	result[middle] = slice[middle]
	for i := 0; i < middle; i++ {
		mirrorIdx := itemCount - i - 1
		result[i], result[mirrorIdx] = slice[mirrorIdx], slice[i]
	}
	return result
}

// Sort - receives a slice of type T and a sorter function.
func Sort[T any](slice []T, sorter func(a T, b T) int) []T {
	result := make([]T, len(slice), cap(slice))
	for i, a := range slice {
		if i <= len(slice)-2 {
			b := slice[i+1]
			priority := sorter(a, b)
			if priority <= 0 {
				result[i] = a
				result[i+1] = b
			} else {
				result[i] = b
				result[i+1] = a
			}
		}
	}
	return result
}

// Unique - receives a slice of type T and returns a slice of type T
// containing all unique elements
func Unique[T comparable](slice []T) []T {
	var unique []T
	visited := map[T]bool{}

	for _, value := range slice {
		if exists := visited[value]; !exists {
			unique = append(unique, value)
			visited[value] = true
		}
	}
	return unique
}
