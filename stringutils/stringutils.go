// This package includes utility functions for string manipulation and formatting.
// It provides helpers for padding, capitalization, and type-safe string conversion.

package stringutils

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Default options.
const (
	DefaultPrecision      = 2
	DefaultBase           = 10
	DefaultFormat         = 'f'
	DefaultNilMapFormat   = "{}"
	DefaultNilSliceFormat = "[]"
	DefaultNilFormat      = "<nil>"
)

// Options - stringification options.
type Options struct {
	Base           int
	Format         byte
	Precision      int
	NilFormat      string
	NilMapFormat   string
	NilSliceFormat string
}

func parseOptions(opts ...Options) Options {
	options := Options{
		Base:           DefaultBase,
		Precision:      DefaultPrecision,
		Format:         DefaultFormat,
		NilFormat:      DefaultNilFormat,
		NilMapFormat:   DefaultNilMapFormat,
		NilSliceFormat: DefaultNilSliceFormat,
	}

	for _, opt := range opts {
		if opt.Base != 0 {
			options.Base = opt.Base
		}
		if opt.Precision != 0 {
			options.Precision = opt.Precision
		}
		if opt.Format != 0 {
			options.Format = opt.Format
		}
		if opt.NilFormat != "" {
			options.NilFormat = opt.NilFormat
		}
		if opt.NilMapFormat != "" {
			options.NilMapFormat = opt.NilMapFormat
		}
		if opt.NilSliceFormat != "" {
			options.NilSliceFormat = opt.NilSliceFormat
		}
	}

	return options
}

// Stringify receives an arbitrary value and converts it into a string.
// Stringify also accepts an options object with the following properties:
//
//	Options.NilFormat: the string format for nil values, defaults to "<nil>".
//	Options.NilMapFormat: the string format for nil map objects, defaults to "{}".
//	Options.NilSliceFormat: the string format for nil slice objects, defaults to "[]".
//	Options.Base: a number between 2-36 as the base when converting ints and uints to strings, defaults to Base 10.
//	Options.Precision: number of digits to include when converting floats and complex numbers to strings, defaults to 2.
//	Options.Format: the number notation format, using the stdlib ftoa functionalities, defaults to 'f':
//		'b' (-ddddp±ddd, a binary exponent),
//		'e' (-d.dddde±dd, a decimal exponent),
//		'E' (-d.ddddE±dd, a decimal exponent),
//		'f' (-ddd.dddd, no exponent),
//		'g' ('e' for large exponents, 'f' otherwise),
//		'G' ('E' for large exponents, 'f' otherwise),
//		'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
//		'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
func Stringify(value any, opts ...Options) string {
	options := parseOptions(opts...)

	// first use a type switch, because its more performant
	switch assigned := value.(type) {
	case string:
		return assigned
	case bool:
		return strconv.FormatBool(assigned)
	case nil:
		return options.NilFormat
	case []byte:
		return string(assigned)
	case int:
		return strconv.FormatInt(int64(assigned), options.Base)
	case int8:
		return strconv.FormatInt(int64(assigned), options.Base)
	case int16:
		return strconv.FormatInt(int64(assigned), options.Base)
	case int32:
		return strconv.FormatInt(int64(assigned), options.Base)
	case int64:
		return strconv.FormatInt(assigned, options.Base)
	case uint:
		return strconv.FormatUint(uint64(assigned), options.Base)
	case uint8:
		return strconv.FormatUint(uint64(assigned), options.Base)
	case uint16:
		return strconv.FormatUint(uint64(assigned), options.Base)
	case uint32:
		return strconv.FormatUint(uint64(assigned), options.Base)
	case uint64:
		return strconv.FormatUint(assigned, options.Base)
	case float32:
		return strconv.FormatFloat(float64(assigned), options.Format, options.Precision, 32)
	case float64:
		return strconv.FormatFloat(assigned, options.Format, options.Precision, 64)
	case complex64:
		return strconv.FormatComplex(complex128(assigned), options.Format, options.Precision, 64)
	case complex128:
		return strconv.FormatComplex(assigned, options.Format, options.Precision, 128)
	case fmt.Stringer:
		return assigned.String()
	case fmt.GoStringer:
		return assigned.GoString()
	}

	// fallback now to using reflection
	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		v := reflect.ValueOf(value)
		if v.IsNil() {
			return options.NilMapFormat
		}

		elements := make([]string, 0)
		for _, mapKey := range v.MapKeys() {
			mapValue := v.MapIndex(mapKey)
			elements = append(
				elements,
				Stringify(
					mapKey.Interface(),
					options,
				)+": "+Stringify(
					mapValue.Interface(),
					options,
				),
			)
		}
		// we sort to ensure deterministic results, given that map keys are arbitrarily ordered
		sort.Strings(elements)
		return "{" + strings.Join(elements, ", ") + "}"
	case reflect.Slice:
		s := reflect.ValueOf(value)
		if s.IsNil() {
			return options.NilSliceFormat
		}

		elements := make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			elements[i] = Stringify(s.Index(i).Interface(), options)
		}
		return "[" + strings.Join(elements, ", ") + "]"
	}

	// finally, we use the least performant solution, which is to fmt.Sprintf the value
	return fmt.Sprintf("%v", value)
}

// PadLeft - Pad a string to a certain length (in bytes) with another string on the left side.
func PadLeft(str string, padWith string, padTo int) string {
	strLen := len(str)

	return getPaddingString(padWith, padTo-strLen) + str
}

// PadRight - Pad a string to a certain length (in bytes) with another string on the right side.
func PadRight(str string, padWith string, padTo int) string {
	strLen := len(str)

	return str + getPaddingString(padWith, padTo-strLen)
}

func getPaddingString(padWith string, padLength int) string {
	if padWith == "" || padLength <= 0 {
		return ""
	}

	var builder strings.Builder

	for builder.Len() < padLength {
		builder.WriteString(padWith)
	}

	return builder.String()[:padLength]
}

// Capitalize capitalizes the first letter of the given string.
// Returns the string unchanged if empty.
func Capitalize(str string) string {
	if str == "" {
		return str
	}
	firstLetter := rune(str[0])
	return string(unicode.ToUpper(firstLetter)) + str[1:]
}

// Reverse reverses a string (UTF-8 safe).
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Truncate truncates a string to maxLength with optional suffix.
// If the string is longer than maxLength, it's cut and suffix is appended.
func Truncate(str string, maxLength int, suffix string) string {
	if len(str) <= maxLength {
		return str
	}
	return str[:maxLength] + suffix
}

// Contains checks if string contains substring with optional case-insensitive matching.
func Contains(str, substr string, caseInsensitive bool) bool {
	if caseInsensitive {
		return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
	}
	return strings.Contains(str, substr)
}

// SplitAndTrim splits a string by separator and trims whitespace from each part.
// Empty parts are removed from the result.
func SplitAndTrim(str, sep string) []string {
	parts := strings.Split(str, sep)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// JoinNonEmpty joins strings with a separator, skipping empty strings.
func JoinNonEmpty(strs []string, sep string) string {
	nonEmpty := make([]string, 0, len(strs))
	for _, s := range strs {
		if s != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}
	return strings.Join(nonEmpty, sep)
}

// IsEmpty checks if string is empty or only whitespace.
func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// DefaultIfEmpty returns default if string is empty or only whitespace.
// Returns the trimmed string if not empty.
func DefaultIfEmpty(str, def string) string {
	trimmed := strings.TrimSpace(str)
	if trimmed == "" {
		return def
	}
	return trimmed
}

// ToTitle converts string to Title Case.
func ToTitle(str string) string {
	return strings.Title(str) //nolint:staticcheck // Title is deprecated but simple for basic use
}

// ToCamelCase converts string to camelCase.
// Splits on spaces, hyphens, and underscores.
func ToCamelCase(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return ""
	}

	// Split on space, hyphen, underscore
	words := strings.FieldsFunc(str, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	if len(words) == 0 {
		return ""
	}

	// First word lowercase, rest capitalized
	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		result += Capitalize(strings.ToLower(words[i]))
	}

	return result
}

// ToSnakeCase converts string to snake_case.
func ToSnakeCase(str string) string {
	var result strings.Builder
	str = strings.TrimSpace(str)
	runes := []rune(str)

	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 {
			prevChar := runes[i-1]
			// Add underscore if:
			// 1. Previous char is not uppercase/separator, OR
			// 2. Next char exists and is lowercase (handles acronyms like HTTPResponse)
			if (!unicode.IsUpper(prevChar) && prevChar != '_' && prevChar != ' ' && prevChar != '-') ||
				(i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
				result.WriteRune('_')
			}
		}

		if r == ' ' || r == '-' {
			result.WriteRune('_')
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

// ToKebabCase converts string to kebab-case.
func ToKebabCase(str string) string {
	var result strings.Builder
	str = strings.TrimSpace(str)
	runes := []rune(str)

	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 {
			prevChar := runes[i-1]
			// Add hyphen if:
			// 1. Previous char is not uppercase/separator, OR
			// 2. Next char exists and is lowercase (handles acronyms like HTTPResponse)
			if (!unicode.IsUpper(prevChar) && prevChar != '-' && prevChar != ' ' && prevChar != '_') ||
				(i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
				result.WriteRune('-')
			}
		}

		if r == ' ' || r == '_' {
			result.WriteRune('-')
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}

// Repeat repeats a string n times.
func Repeat(str string, n int) string {
	return strings.Repeat(str, n)
}

// RemoveWhitespace removes all whitespace from string.
func RemoveWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// EllipsisMiddle truncates middle of string with ellipsis if longer than maxLength.
// Keeps approximately equal parts from start and end.
func EllipsisMiddle(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}

	ellipsis := "..."
	if maxLength < len(ellipsis) {
		return ellipsis[:maxLength]
	}

	sideLength := (maxLength - len(ellipsis)) / 2
	return str[:sideLength] + ellipsis + str[len(str)-sideLength:]
}
