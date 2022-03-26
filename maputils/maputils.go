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

// Merge - gives a maps with key type K and values of type V, merge the separate maps into a single map.
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

func ForEach[K comparable, V any](mapInstance map[K]V, function func(key K, value V)) {
	for key, value := range mapInstance {
		function(key, value)
	}
}

func ReMap[K comparable, V any](mapInstance map[K]V, function func(key K, value V) K) map[K]V {
	for key, value := range mapInstance {
		newKey := function(key, value)
		mapInstance[newKey] = value
		delete(mapInstance, key)
	}
	return mapInstance
}

func Drop[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V {
	for _, key := range keys {
		if _, keyExists := mapInstance[key]; keyExists {
			delete(mapInstance, key)
		}
	}
	return mapInstance
}
