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

var nilMap map[string]any
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
		t.Run(
			fmt.Sprintf(
				"TestCase: input: %v, expected: %v",
				testCase.input,
				testCase.expectedOutput,
			),
			func(t *testing.T) {
				assert.Equal(
					t,
					testCase.expectedOutput,
					stringutils.Stringify(testCase.input, testCase.options),
				)
			},
		)
	}
}

func TestPadLeft(t *testing.T) {
	assert.Equal(t, "_-_Azer", stringutils.PadLeft("Azer", "_-", 7))
	assert.Equal(t, "Test", stringutils.PadLeft("Test", "Großmeister", 0))
	assert.Equal(t, "Test", stringutils.PadLeft("Test", "Großmeister", -1))
	assert.Equal(t, "Test", stringutils.PadLeft("Test", "", 7))
	assert.Equal(t, "GroTest", stringutils.PadLeft("Test", "Großmeister", 7))
	assert.Equal(t, "Gro\xc3Test", stringutils.PadLeft("Test", "Großmeister", 8))
	assert.Equal(t, "GroßTest", stringutils.PadLeft("Test", "Großmeister", 9))
	assert.Equal(t, "GroßmeisterGroßTest", stringutils.PadLeft("Test", "Großmeister", 21))
}

func TestPadRight(t *testing.T) {
	assert.Equal(t, "Azer-_-", stringutils.PadRight("Azer", "-_", 7))
	assert.Equal(t, "Test", stringutils.PadRight("Test", "Großmeister", 0))
	assert.Equal(t, "Test", stringutils.PadRight("Test", "Großmeister", -1))
	assert.Equal(t, "Test", stringutils.PadRight("Test", "", 7))
	assert.Equal(t, "TestGro", stringutils.PadRight("Test", "Großmeister", 7))
	assert.Equal(t, "TestGro\xc3", stringutils.PadRight("Test", "Großmeister", 8))
	assert.Equal(t, "TestGroß", stringutils.PadRight("Test", "Großmeister", 9))
	assert.Equal(t, "TestGroßmeisterGroß", stringutils.PadRight("Test", "Großmeister", 21))
}

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "store",
			expected: "Store",
		},
		{
			input:    "batman",
			expected: "Batman",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "rusty",
			expected: "Rusty",
		},
	}

	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("TestCase: input: %v, expected: %v", testCase.input, testCase.expected),
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, stringutils.Capitalize(testCase.input))
			},
		)
	}
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "olleh", stringutils.Reverse("hello"))
	assert.Equal(t, "!dlrow", stringutils.Reverse("world!"))
	assert.Equal(t, "", stringutils.Reverse(""))
	assert.Equal(t, "a", stringutils.Reverse("a"))
	// UTF-8 test
	assert.Equal(t, "👋🌍", stringutils.Reverse("🌍👋"))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "hello", stringutils.Truncate("hello", 10, "..."))
	assert.Equal(t, "hello...", stringutils.Truncate("hello world", 5, "..."))
	assert.Equal(t, "...", stringutils.Truncate("test", 0, "..."))
	assert.Equal(t, "hello [more]", stringutils.Truncate("hello world", 5, " [more]"))
}

func TestContains(t *testing.T) {
	assert.True(t, stringutils.Contains("hello world", "world", false))
	assert.False(t, stringutils.Contains("hello world", "WORLD", false))
	assert.True(t, stringutils.Contains("hello world", "WORLD", true))
	assert.True(t, stringutils.Contains("Hello World", "hello", true))
	assert.False(t, stringutils.Contains("test", "xyz", false))
}

func TestSplitAndTrim(t *testing.T) {
	result := stringutils.SplitAndTrim("  hello , world ,  foo  ", ",")
	assert.Equal(t, []string{"hello", "world", "foo"}, result)

	result = stringutils.SplitAndTrim("a,  , b,  ,c", ",")
	assert.Equal(t, []string{"a", "b", "c"}, result)

	result = stringutils.SplitAndTrim("", ",")
	assert.Empty(t, result)
}

func TestJoinNonEmpty(t *testing.T) {
	assert.Equal(t, "hello,world", stringutils.JoinNonEmpty([]string{"hello", "", "world"}, ","))
	assert.Equal(t, "a-b-c", stringutils.JoinNonEmpty([]string{"a", "b", "c"}, "-"))
	assert.Equal(t, "test", stringutils.JoinNonEmpty([]string{"", "test", ""}, ","))
	assert.Equal(t, "", stringutils.JoinNonEmpty([]string{}, ","))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, stringutils.IsEmpty(""))
	assert.True(t, stringutils.IsEmpty("   "))
	assert.True(t, stringutils.IsEmpty("\t\n"))
	assert.False(t, stringutils.IsEmpty("hello"))
	assert.False(t, stringutils.IsEmpty(" hello "))
}

func TestDefaultIfEmpty(t *testing.T) {
	assert.Equal(t, "default", stringutils.DefaultIfEmpty("", "default"))
	assert.Equal(t, "default", stringutils.DefaultIfEmpty("   ", "default"))
	assert.Equal(t, "hello", stringutils.DefaultIfEmpty("hello", "default"))
	assert.Equal(t, "hello", stringutils.DefaultIfEmpty(" hello ", "default"))
}

func TestToTitle(t *testing.T) {
	assert.Equal(t, "Hello World", stringutils.ToTitle("hello world"))
	assert.Equal(t, "HELLO WORLD", stringutils.ToTitle("HELLO WORLD"))
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "helloWorld", stringutils.ToCamelCase("hello world"))
	assert.Equal(t, "helloWorldFoo", stringutils.ToCamelCase("hello-world-foo"))
	assert.Equal(t, "helloWorldBar", stringutils.ToCamelCase("hello_world_bar"))
	assert.Equal(t, "test", stringutils.ToCamelCase("test"))
	assert.Equal(t, "", stringutils.ToCamelCase(""))
	assert.Equal(t, "firstName", stringutils.ToCamelCase("first name"))
	assert.Equal(t, "", stringutils.ToCamelCase("___"))
	assert.Equal(t, "", stringutils.ToCamelCase("---"))
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "hello_world", stringutils.ToSnakeCase("HelloWorld"))
	assert.Equal(t, "hello_world", stringutils.ToSnakeCase("helloWorld"))
	assert.Equal(t, "hello_world", stringutils.ToSnakeCase("hello-world"))
	assert.Equal(t, "hello_world", stringutils.ToSnakeCase("hello world"))
	assert.Equal(t, "test", stringutils.ToSnakeCase("test"))
	assert.Equal(t, "first_name", stringutils.ToSnakeCase("firstName"))
	assert.Equal(t, "http_response", stringutils.ToSnakeCase("HTTPResponse"))
}

func TestToKebabCase(t *testing.T) {
	assert.Equal(t, "hello-world", stringutils.ToKebabCase("HelloWorld"))
	assert.Equal(t, "hello-world", stringutils.ToKebabCase("helloWorld"))
	assert.Equal(t, "hello-world", stringutils.ToKebabCase("hello_world"))
	assert.Equal(t, "hello-world", stringutils.ToKebabCase("hello world"))
	assert.Equal(t, "test", stringutils.ToKebabCase("test"))
	assert.Equal(t, "first-name", stringutils.ToKebabCase("firstName"))
	assert.Equal(t, "http-response", stringutils.ToKebabCase("HTTPResponse"))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaa", stringutils.Repeat("a", 3))
	assert.Equal(t, "hihihi", stringutils.Repeat("hi", 3))
	assert.Equal(t, "", stringutils.Repeat("test", 0))
	assert.Equal(t, "x", stringutils.Repeat("x", 1))
}

func TestRemoveWhitespace(t *testing.T) {
	assert.Equal(t, "helloworld", stringutils.RemoveWhitespace("hello world"))
	assert.Equal(t, "test", stringutils.RemoveWhitespace(" t e s t "))
	assert.Equal(t, "foo", stringutils.RemoveWhitespace("\tfoo\n"))
	assert.Equal(t, "", stringutils.RemoveWhitespace("   "))
}

func TestEllipsisMiddle(t *testing.T) {
	assert.Equal(t, "hello", stringutils.EllipsisMiddle("hello", 10))
	assert.Equal(t, "hel...rld", stringutils.EllipsisMiddle("hello world", 9))
	assert.Equal(t, "he...ld", stringutils.EllipsisMiddle("hello world", 7))
	assert.Equal(t, "...", stringutils.EllipsisMiddle("hello", 3))
	assert.Equal(t, "..", stringutils.EllipsisMiddle("hello", 2))
}
