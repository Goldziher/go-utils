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

// Parse parses a raw URL and returns a URL struct or an error.
// This is a convenience wrapper around url.Parse.
func Parse(rawURL string) (*url.URL, error) {
	return url.Parse(rawURL)
}

// MustParse parses a raw URL and panics if parsing fails.
// Use this when you're certain the URL is valid.
func MustParse(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}

// IsAbsolute returns true if the URL is absolute (has a scheme).
func IsAbsolute(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return u.IsAbs()
}

// GetDomain extracts the domain (host) from a URL string.
// Returns empty string if parsing fails or no host is present.
func GetDomain(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Host
}

// GetScheme extracts the scheme from a URL string.
// Returns empty string if parsing fails or no scheme is present.
func GetScheme(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Scheme
}

// GetPath extracts the path from a URL string.
// Returns empty string if parsing fails.
func GetPath(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Path
}
