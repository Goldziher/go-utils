// This package includes utility functions for handling and manipulating maputils.
// It draws inspiration from JavaScript and Python and uses Go generics as a basis.

package maputils

// Keys - takes a map with keys K and values V, returns a slice of type K of the map's keys.
// Note: Go maps do not preserve insertion order.
func Keys[K comparable, V any](mapInstance map[K]V) []K {
	keys := make([]K, len(mapInstance))

	i := 0
	for k := range mapInstance {
		keys[i] = k
		i++
	}

	return keys
}

// Values - takes a map with keys K and values V, returns a slice of type V of the map's values.
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
// Note: merge works from left to right. If a key already exists in a previous map, its value is over-written.
func Merge[K comparable, V any](mapInstances ...map[K]V) map[K]V {
	var mergedMapSize int

	for _, mapInstance := range mapInstances {
		mergedMapSize += len(mapInstance)
	}

	mergedMap := make(map[K]V, mergedMapSize)

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

// Drop - takes a map with keys K and values V, and a slice of keys K, dropping all the key-value pairs that match the keys in the slice.
// Note: this function will modify the passed in map. To get a different object, use the Copy function to pass a copy to this function.
func Drop[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V {
	for _, key := range keys {
		delete(mapInstance, key)
	}

	return mapInstance
}

// Copy - takes a map with keys K and values V and returns a copy of the map.
//
// Deprecated: Copy is deprecated as of Go 1.21. Use maps.Clone from the standard library instead:
//
//	import "maps"
//	copied := maps.Clone(original)
//
// This function will be removed in a future major version.
func Copy[K comparable, V any](mapInstance map[K]V) map[K]V {
	mapCopy := make(map[K]V, len(mapInstance))

	for key, value := range mapInstance {
		mapCopy[key] = value
	}

	return mapCopy
}

// Filter - takes a map with keys K and values V, and executes the passed in function for each key-value pair.
// If the filter function returns true, the key-value pair will be included in the output, otherwise it is filtered out.
func Filter[K comparable, V any](mapInstance map[K]V, function func(key K, value V) bool) map[K]V {
	mapCopy := make(map[K]V, len(mapInstance))

	for key, value := range mapInstance {
		if function(key, value) {
			mapCopy[key] = value
		}
	}

	return mapCopy
}

// Map transforms map values using a mapper function, returning a new map with transformed values.
func Map[K comparable, V any, R any](mapInstance map[K]V, mapper func(key K, value V) R) map[K]R {
	result := make(map[K]R, len(mapInstance))

	for key, value := range mapInstance {
		result[key] = mapper(key, value)
	}

	return result
}

// MapKeys transforms map keys using a mapper function, returning a new map with transformed keys.
// Note: If the mapper produces duplicate keys, later values will overwrite earlier ones.
func MapKeys[K comparable, V any, R comparable](
	mapInstance map[K]V,
	mapper func(key K, value V) R,
) map[R]V {
	result := make(map[R]V, len(mapInstance))

	for key, value := range mapInstance {
		newKey := mapper(key, value)
		result[newKey] = value
	}

	return result
}

// Invert swaps keys and values in a map.
// Note: If multiple keys have the same value, only one will remain (non-deterministic which one).
// Both keys and values must be comparable types.
func Invert[K, V comparable](mapInstance map[K]V) map[V]K {
	result := make(map[V]K, len(mapInstance))

	for key, value := range mapInstance {
		result[value] = key
	}

	return result
}

// Pick creates a new map containing only the specified keys.
// Keys that don't exist in the original map are ignored.
func Pick[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	for _, key := range keys {
		if value, exists := mapInstance[key]; exists {
			result[key] = value
		}
	}

	return result
}

// Omit creates a new map excluding the specified keys.
func Omit[K comparable, V any](mapInstance map[K]V, keys []K) map[K]V {
	keysToOmit := make(map[K]bool, len(keys))
	for _, key := range keys {
		keysToOmit[key] = true
	}

	result := make(map[K]V)
	for key, value := range mapInstance {
		if !keysToOmit[key] {
			result[key] = value
		}
	}

	return result
}

// Has checks if a key exists in the map.
func Has[K comparable, V any](mapInstance map[K]V, key K) bool {
	_, exists := mapInstance[key]
	return exists
}

// Get safely gets a value from the map with a default fallback if the key doesn't exist.
func Get[K comparable, V any](mapInstance map[K]V, key K, defaultValue V) V {
	if value, exists := mapInstance[key]; exists {
		return value
	}
	return defaultValue
}

// ToEntries converts a map to a slice of key-value pairs.
// Note: Order is non-deterministic due to map iteration order.
func ToEntries[K comparable, V any](mapInstance map[K]V) [][2]any {
	entries := make([][2]any, 0, len(mapInstance))

	for key, value := range mapInstance {
		entries = append(entries, [2]any{key, value})
	}

	return entries
}

// FromEntries creates a map from a slice of key-value pairs.
// If duplicate keys exist, later values overwrite earlier ones.
func FromEntries[K comparable, V any](entries [][2]any) map[K]V {
	result := make(map[K]V, len(entries))

	for _, entry := range entries {
		if key, ok := entry[0].(K); ok {
			if value, ok := entry[1].(V); ok {
				result[key] = value
			}
		}
	}

	return result
}

// GroupBy groups map entries by a grouping key generated from each entry.
// Returns a map where keys are group identifiers and values are maps of entries belonging to that group.
func GroupBy[K comparable, V any, G comparable](
	mapInstance map[K]V,
	grouper func(key K, value V) G,
) map[G]map[K]V {
	result := make(map[G]map[K]V)

	for key, value := range mapInstance {
		group := grouper(key, value)
		if result[group] == nil {
			result[group] = make(map[K]V)
		}
		result[group][key] = value
	}

	return result
}
