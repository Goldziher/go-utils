package stringutils_test

import (
	"fmt"
	"testing"

	"github.com/Goldziher/go-utils/stringutils"
	"github.com/stretchr/testify/assert"
)

type stringerType struct{}

func (s stringerType) String() string {
	return "Stringer"
}

type goStringerType struct{}

func (s goStringerType) GoString() string {
	return "GoStringer"
}

var nilMap map[string]interface{}
var nilSlice []string

func TestStringify(t *testing.T) {
	testCases := []struct {
		input          any
		expectedOutput string
		options        stringutils.Options
	}{
		{
			"abc",
			"abc",
			stringutils.Options{},
		},
		{
			true,
			"true",
			stringutils.Options{},
		},
		{
			false,
			"false",
			stringutils.Options{},
		},
		{
			[]byte("abc"),
			"abc",
			stringutils.Options{},
		},
		{
			nil,
			stringutils.DefaultNilFormat,
			stringutils.Options{},
		},
		{
			nil,
			"null",
			stringutils.Options{NilFormat: "null"},
		},
		{
			1,
			"1",
			stringutils.Options{},
		},
		{
			10,
			"1010",
			stringutils.Options{Base: 2},
		},
		{
			int8(1),
			"1",
			stringutils.Options{},
		},
		{
			int16(1),
			"1",
			stringutils.Options{},
		},
		{
			int32(1),
			"1",
			stringutils.Options{},
		},
		{
			int64(1),
			"1",
			stringutils.Options{},
		},
		{
			uint(1),
			"1",
			stringutils.Options{},
		},
		{
			uint8(1),
			"1",
			stringutils.Options{},
		},
		{
			uint16(1),
			"1",
			stringutils.Options{},
		},
		{
			uint32(1),
			"1",
			stringutils.Options{},
		},
		{
			uint64(1),
			"1",
			stringutils.Options{},
		},
		{
			float32(1),
			"1.00",
			stringutils.Options{},
		},
		{
			float32(1),
			"1.0000",
			stringutils.Options{Precision: 4},
		},
		{
			float32(1),
			"1.00E+00",
			stringutils.Options{Format: 'E'},
		},
		{
			float64(1),
			"1.00",
			stringutils.Options{},
		},
		{
			complex64(1),
			"(1.00+0.00i)",
			stringutils.Options{},
		},
		{
			complex128(1),
			"(1.00+0.00i)",
			stringutils.Options{},
		},
		{
			stringerType{},
			"Stringer",
			stringutils.Options{},
		},
		{
			goStringerType{},
			"GoStringer",
			stringutils.Options{},
		},
		{
			nilMap,
			stringutils.DefaultNilMapFormat,
			stringutils.Options{},
		},
		{
			nilMap,
			stringutils.DefaultNilFormat,
			stringutils.Options{NilMapFormat: stringutils.DefaultNilFormat},
		},
		{
			map[string]int{"one": 1, "two": 2},
			"{one: 1, two: 2}",
			stringutils.Options{},
		},
		{
			nilSlice,
			stringutils.DefaultNilSliceFormat,
			stringutils.Options{},
		},
		{
			nilSlice,
			stringutils.DefaultNilFormat,
			stringutils.Options{NilSliceFormat: stringutils.DefaultNilFormat},
		},
		{
			[]int{1, 2, 3},
			"[1, 2, 3]",
			stringutils.Options{},
		},
		{
			struct {
				X int
				Y int
			}{X: 1, Y: 10},
			"{1 10}",
			stringutils.Options{},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase: input: %v, expected: %v", testCase.input, testCase.expectedOutput), func(t *testing.T) {
			assert.Equal(t, testCase.expectedOutput, stringutils.Stringify(testCase.input, testCase.options))
		})
	}
}
