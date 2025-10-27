// This package includes utility functions for URL and query string manipulation.
// It provides helpers for building query strings from maps and structs.

package urlutils

import (
	"net/url"
	"reflect"

	"github.com/Goldziher/go-utils/stringutils"
)

// QueryStringifyMap creates a query string from a given map instance.
func QueryStringifyMap[K comparable, V any](values map[K]V) string {
	query := url.Values{}

	for key, value := range values {
		stringifiedKey := stringutils.Stringify(key)
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			s := reflect.ValueOf(value)
			if s.IsNil() {
				query.Add(stringifiedKey, "")
			}

			for i := 0; i < s.Len(); i++ {
				query.Add(stringifiedKey, stringutils.Stringify(s.Index(i).Interface()))
			}
		} else {
			query.Add(stringifiedKey, stringutils.Stringify(value))
		}
	}

	return query.Encode()
}

// QueryStringifyStruct creates a query string from a given struct instance. Takes struct tag names as optional parameters.
func QueryStringifyStruct[T any](values T, structTags ...string) string {
	query := url.Values{}

	typeOf := reflect.TypeOf(values)
	valueOf := reflect.ValueOf(values)
	fields := reflect.VisibleFields(typeOf)

	for _, field := range fields {
		key := field.Name

		if field.Tag != "" {
			for _, s := range structTags {
				if tagValue, isPresent := field.Tag.Lookup(s); isPresent && tagValue != "" {
					key = tagValue
					break
				}
			}
		}

		if key == "-" {
			continue
		}

		value := valueOf.FieldByName(field.Name)

		if value.Kind() == reflect.Slice {
			if value.IsNil() {
				query.Add(key, "")
			}

			for i := 0; i < value.Len(); i++ {
				query.Add(key, stringutils.Stringify(value.Index(i).Interface()))
			}
		} else {
			query.Add(key, stringutils.Stringify(value.Interface()))
		}
	}

	return query.Encode()
}
