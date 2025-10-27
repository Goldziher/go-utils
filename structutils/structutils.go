// This package includes utility functions for working with structs using reflection.
// It provides helpers for iterating over struct fields and converting structs to maps.

package structutils

import "reflect"

// ForEach given a struct, calls the passed in function with each visible struct field's name, value and tag.
func ForEach[T any](structInstance T, function func(key string, value any, tag reflect.StructTag)) {
	typeOf := reflect.TypeOf(structInstance)
	valueOf := reflect.ValueOf(structInstance)
	fields := reflect.VisibleFields(typeOf)

	for _, field := range fields {
		value := valueOf.FieldByName(field.Name)
		function(field.Name, value.Interface(), field.Tag)
	}
}

// ToMap given a struct, converts it to a map[string]any.
// This function also takes struct tag names as optional parameters - if passed in, the struct tags will be used to remap or omit values.
func ToMap[T any](structInstance T, structTags ...string) map[string]any {
	output := make(map[string]any, 0)

	ForEach(structInstance, func(key string, value any, tag reflect.StructTag) {
		if tag != "" {
			for _, s := range structTags {
				if tagValue, isPresent := tag.Lookup(s); isPresent && tagValue != "" {
					key = tagValue
					break
				}
			}
		}
		if key != "-" {
			output[key] = value
		}
	})

	return output
}

// Fields returns a slice of field names from a struct.
func Fields[T any](structInstance T) []string {
	typeOf := reflect.TypeOf(structInstance)
	fields := reflect.VisibleFields(typeOf)
	result := make([]string, len(fields))
	for i, field := range fields {
		result[i] = field.Name
	}
	return result
}

// Values returns a slice of field values from a struct.
func Values[T any](structInstance T) []any {
	typeOf := reflect.TypeOf(structInstance)
	valueOf := reflect.ValueOf(structInstance)
	fields := reflect.VisibleFields(typeOf)
	result := make([]any, len(fields))
	for i, field := range fields {
		value := valueOf.FieldByName(field.Name)
		result[i] = value.Interface()
	}
	return result
}

// HasField checks if a struct has a field with the given name.
func HasField[T any](structInstance T, fieldName string) bool {
	typeOf := reflect.TypeOf(structInstance)
	_, found := typeOf.FieldByName(fieldName)
	return found
}

// GetField retrieves the value of a field by name.
// Returns the value and true if found, zero value and false if not found.
func GetField[T any](structInstance T, fieldName string) (any, bool) {
	valueOf := reflect.ValueOf(structInstance)
	field := valueOf.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, false
	}
	return field.Interface(), true
}

// FieldNames returns field names, optionally using struct tags for naming.
// If struct tags are provided, uses the first matching tag value as the field name.
// Omits fields with tag value "-".
func FieldNames[T any](structInstance T, structTags ...string) []string {
	typeOf := reflect.TypeOf(structInstance)
	fields := reflect.VisibleFields(typeOf)
	result := make([]string, 0, len(fields))

	for _, field := range fields {
		name := field.Name

		if field.Tag != "" && len(structTags) > 0 {
			for _, tag := range structTags {
				if tagValue, isPresent := field.Tag.Lookup(tag); isPresent && tagValue != "" {
					name = tagValue
					break
				}
			}
		}

		if name != "-" {
			result = append(result, name)
		}
	}

	return result
}
