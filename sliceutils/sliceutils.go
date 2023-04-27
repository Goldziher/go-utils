// This package includes utility functions for handling and manipulating a slice.
// It draws inspiration from JavaScript and Python and uses Go generics as a basis.

package sliceutils

import (
	"golang.org/x/exp/constraints"
)

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
	if len(slice) > 0 {
		mapped = make([]R, len(slice))
		for i, el := range slice {
			mapped[i] = mapper(el, i, slice)
		}
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
// Note: The elements are merged in their order in a slice,
// i.e. first the elements of the first slice, then that of the second and so forth.
func Merge[T any](slices ...[]T) (mergedSlice []T) {
	if len(slices) > 0 {
		mergedSliceCap := 0

		for _, slice := range slices {
			mergedSliceCap += len(slice)
		}

		if mergedSliceCap > 0 {
			mergedSlice = make([]T, 0, mergedSliceCap)

			for _, slice := range slices {
				mergedSlice = append(mergedSlice, slice...)
			}
		}
	}
	return mergedSlice
}

// Sum - receives a slice of type T and returns a value T that is the sum of the numbers.
// Note: T is constrained to be a number type.
func Sum[T constraints.Complex | constraints.Integer | constraints.Float](slice []T) (result T) {
	for _, el := range slice {
		result += el
	}
	return result
}

// Remove - receives a slice of type T and an index, removing the element at the given index.
// Note: this function does not modify the input slice.
func Remove[T any](slice []T, i int) []T {
	if len(slice) == 0 || i > len(slice)-1 {
		return slice
	}
	copied := Copy(slice)
	if i == 0 {
		return copied[1:]
	}
	if i != len(copied)-1 {
		return append(copied[:i], copied[i+1:]...)
	}
	return copied[:i]
}

// Insert - receives a slice of type T, an index and a value.
// The value is inserted at the given index. If there is an existing value at this index, it is shifted to the next index.
// Note: this function does not modify the input slice.
func Insert[T any](slice []T, i int, value T) []T {
	if len(slice) == i {
		return append(slice, value)
	}
	slice = append(slice[:i+1], slice[i:]...)
	slice[i] = value
	return slice
}

// Copy - receives a slice of type T and copies it.
func Copy[T any](slice []T) []T {
	duplicate := make([]T, len(slice), cap(slice))
	copy(duplicate, slice)
	return duplicate
}

// Intersection - takes a variadic number of slices of type T and returns a slice of type T containing any values that exist in all the slices.
// For example, given []int{1, 2, 3}, []int{1, 7, 3}, the intersection would be []int{1, 3}.
func Intersection[T comparable](slices ...[]T) []T {
	possibleIntersections := map[T]int{}
	for i, slice := range slices {
		for _, el := range slice {
			if i == 0 {
				possibleIntersections[el] = 0
			} else if _, elementExists := possibleIntersections[el]; elementExists {
				possibleIntersections[el] = i
			}
		}
	}

	intersected := make([]T, 0)
	for _, el := range slices[0] {
		if lastVisitorIndex, exists := possibleIntersections[el]; exists && lastVisitorIndex == len(slices)-1 {
			intersected = append(intersected, el)
			delete(possibleIntersections, el)
		}
	}

	return intersected
}

// Difference - takes a variadic number of slices of type T and returns a slice of type T containing the elements that are different between the slices.
// For example, given []int{1, 2, 3}, []int{2, 3, 4}, []{3, 4, 5}, the difference would be []int{1, 5}.
func Difference[T comparable](slices ...[]T) []T {
	possibleDifferences := map[T]int{}
	nonDifferentElements := map[T]int{}

	for i, slice := range slices {
		for _, el := range slice {
			if lastVisitorIndex, elementExists := possibleDifferences[el]; elementExists && lastVisitorIndex != i {
				nonDifferentElements[el] = i
			} else if !elementExists {
				possibleDifferences[el] = i
			}
		}
	}

	differentElements := make([]T, 0)

	for _, slice := range slices {
		for _, el := range slice {
			if _, exists := nonDifferentElements[el]; !exists {
				differentElements = append(differentElements, el)
			}
		}
	}

	return differentElements
}

// Union - takes a variadic number of slices of type T and returns a slice of type T containing the unique elements in the different slices
// For example, given []int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5}, the union would be []int{1, 2, 3, 4, 5}.
func Union[T comparable](slices ...[]T) []T {
	return Unique(Merge(slices...))
}

// Reverse - takes a slice of type T and returns a slice of type T with a reverse order of elements.
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

// Unique - receives a slice of type T and returns a slice of type T containing all unique elements.
func Unique[T comparable](slice []T) []T {
	unique := make([]T, 0)
	visited := map[T]bool{}

	for _, value := range slice {
		if exists := visited[value]; !exists {
			unique = append(unique, value)
			visited[value] = true
		}
	}
	return unique
}

// Chunk - receives a slice of type T and size N and returns a slice of slices T of size N.
func Chunk[T any](input []T, size int) [][]T {
	var chunks [][]T

	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

// Pluck - receives a slice of type I and a getter func to a field
// and returns a slice containing the requested field's value from each item in the slice.
func Pluck[I any, O any](input []I, getter func(I) *O) []O {
	var output []O

	for _, item := range input {
		field := getter(item)

		if field != nil {
			output = append(output, *field)
		}
	}

	return output
}

// Flatten - receives a slice of slices of type I and flattens it to a slice of type I.
func Flatten[I any](input [][]I) (output []I) {
	if len(input) > 0 {
		var outputSize int

		for _, item := range input {
			outputSize += len(item)
		}

		if outputSize > 0 {
			output = make([]I, 0, outputSize)

			for _, item := range input {
				output = append(output, item...)
			}
		}
	}
	return output
}
