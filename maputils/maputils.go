// This package includes utility functions for handling and manipulating maps.
// It draws inspiration from JavaScript and Python and uses Go generics as a basis.

package maputils

// Keys - given a map with keys K and values V, returns a slice of type K of the map's keys.
// Note: Go maps do not preserve insertion orde.
func Keys[K comparable, V any](mapInstance map[K]V) []K {
	keys := make([]K, len(mapInstance))

	i := 0
	for k := range mapInstance {
		keys[i] = k
		i++
	}

	return keys
}

// Values - given a map with keys K and values V, returns a slice of type V of the map's values.
// Note: Go maps do not preserve insertion order.
func Values[K comparable, V any](mapInstance map[K]V) []V {
	values := make([]V, len(mapInstance))

	i := 0
	for _, v := range mapInstance {
		values[i] = v
		i++
	}

	return values
}

// Merge - takes an arbitrary number of map instances with keys K and values V and merges them into a single map.
// Note: merge works from left most to rightmost. If a key already exists in a previous map, its value is over-written.
func Merge[K comparable, V any](mapInstances ...map[K]V) map[K]V {
	mergedMap := make(map[K]V, 0)

	for _, mapInstance := range mapInstances {
		for k, v := range mapInstance {
			mergedMap[k] = v
		}
	}
	return mergedMap
}

// ForEach - given a map with keys K and values V, executes the passed in function for each key-value pair.
func ForEach[K comparable, V any](mapInstance map[K]V, function func(key K, value V)) {
	for key, value := range mapInstance {
		function(key, value)
	}
}

// ReMap - takes a map with keys K and values V, executing the remapper function for each each key-value pair.
// The remapper function should return a key of type K, either the original key or a new key.
// If the key is new, the old key is deleted and the value is reassigned to the new key.
// Note: this function will modify the passed in map. To get a different object, use the Copy function to pass a copy to this functiont.
func ReMap[K comparable, V any](mapInstance map[K]V, function func(key K, value V) K) map[K]V {
	keysToDrop := make([]K, 0)
	for key, value := range mapInstance {
		outputKey := function(key, value)
		if outputKey != key {
			mapInstance[outputKey] = value
			keysToDrop = append(keysToDrop, key)
		}
	}
	mapInstance = Drop(mapInstance, keysToDrop)
	return mapInstance
}

// Drop - takes a map with keys K and values V, and a slice of keys K, dropping all the key-value pairs that match the keys in the slice.
// Note: this function will modify the passed in map. To get a different object, use the Copy function to pass a copy to this functiont.
func Drop[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V {
	for _, key := range keys {
		if _, keyExists := mapInstance[key]; keyExists {
			delete(mapInstance, key)
		}
	}
	return mapInstance
}

// Copy - take a map with keys K and values V and returns a copy of the map.
func Copy[K comparable, V any](mapInstance map[K]V) map[K]V {
	mapCopy := make(map[K]V, len(mapInstance))

	for key, value := range mapInstance {
		mapCopy[key] = value
	}

	return mapCopy
}
