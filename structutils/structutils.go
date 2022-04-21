package structutils

import "reflect"

// ForEach given a struct, calls the passed in function with each visible struct field's name, value and tag.
func ForEach[T interface{}](structInstance T, function func(key string, value interface{}, tag reflect.StructTag)) {
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
func ToMap[T interface{}](structInstance T, structTags ...string) map[string]any {
	output := make(map[string]any, 0)

	ForEach(structInstance, func(key string, value interface{}, tag reflect.StructTag) {
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
