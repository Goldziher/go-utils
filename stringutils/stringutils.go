package stringutils

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

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
		if opt.Format != uint8(0) {
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
//	Options.Base: a number between 2-36 ad the base when converting ints and uints to strings, defaults to Base 10.
//	Options.Precision: number of digits to include when converting floats and complex numbers to strings, defaults to 2.
//	Options.Format: the number notation format, using the stlib ftoa functionalities, defaults to 'f':
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
			elements = append(elements, Stringify(mapKey.Interface(), options)+": "+Stringify(mapValue.Interface(), options))
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

// PadLeft - Pad a string to a certain length with another string on the left side.
func PadLeft(str string, padWith string, padTo int) string {
	strLen := len(str)

	return getPaddingString(padWith, padTo-strLen) + str
}

// PadRight - Pad a string to a certain length with another string on the right side.
func PadRight(str string, padWith string, padTo int) string {
	strLen := len(str)

	return str + getPaddingString(padWith, padTo-strLen)
}

func getPaddingString(padWith string, padLength int) string {
	padding := ""
	for i := 0; i <= padLength; i++ {
		padding += padWith
	}

	return padding[:padLength]
}
