package urlutils_test

import (
	"fmt"
	"testing"

	"github.com/Goldziher/go-utils/urlutils"
	"github.com/stretchr/testify/assert"
)

var nilSlice []int

var nilMap map[string]any

func TestQueryStringifyMap(t *testing.T) {
	t.Run("Test nil map", func(t *testing.T) {
		actualOutput := urlutils.QueryStringifyMap(nilMap)
		assert.Equal(t, "", actualOutput)
	})
	t.Run("Test string any map", func(t *testing.T) {
		actualOutput := urlutils.QueryStringifyMap(map[string]any{
			"user":    "moishe",
			"active":  true,
			"age":     100,
			"friends": []int{1, 2, 3, 4, 5, 6},
		})
		assert.Equal(
			t,
			"active=true&age=100&friends=1&friends=2&friends=3&friends=4&friends=5&friends=6&user=moishe",
			actualOutput,
		)
	})
	t.Run("Test string any map with nil slice", func(t *testing.T) {
		actualOutput := urlutils.QueryStringifyMap(map[string]any{
			"user":    "moishe",
			"active":  true,
			"age":     100,
			"friends": nilSlice,
		})
		assert.Equal(t, "active=true&age=100&friends=&user=moishe", actualOutput)
	})
	t.Run("Test string int map", func(t *testing.T) {
		actualOutput := urlutils.QueryStringifyMap(map[string]int{
			"user": 1,
		})
		assert.Equal(t, "user=1", actualOutput)
	})
	t.Run("Test int int map", func(t *testing.T) {
		actualOutput := urlutils.QueryStringifyMap(map[int]int{
			1: 2,
		})
		assert.Equal(t, "1=2", actualOutput)
	})
}

func TestQueryStringifyStruct(t *testing.T) {
	testCases := []struct {
		input          any
		expectedOutput string
		structTag      string
	}{
		{
			struct {
				User    string
				Active  bool
				Age     int
				Friends []int
			}{
				User:    "moishe",
				Active:  true,
				Age:     100,
				Friends: []int{1, 2, 3, 4, 5, 6},
			},
			"Active=true&Age=100&Friends=1&Friends=2&Friends=3&Friends=4&Friends=5&Friends=6&User=moishe",
			"",
		},
		{
			struct {
				User    string
				Active  bool
				Age     int
				Friends []int
			}{
				User:    "moishe",
				Active:  true,
				Age:     100,
				Friends: nilSlice,
			},
			"Active=true&Age=100&Friends=&User=moishe",
			"",
		},
		{
			struct {
				User    string
				Active  bool
				Age     int `qs:"-"`
				Friends []int
			}{
				User:    "moishe",
				Active:  true,
				Age:     100,
				Friends: []int{1, 2, 3, 4, 5, 6},
			},
			"Active=true&Friends=1&Friends=2&Friends=3&Friends=4&Friends=5&Friends=6&User=moishe",
			"qs",
		},
		{
			struct {
				User    string
				Active  bool
				Age     int `qs:"age"`
				Friends []int
			}{
				User:    "moishe",
				Active:  true,
				Age:     100,
				Friends: []int{1, 2, 3, 4, 5, 6},
			},
			"Active=true&Friends=1&Friends=2&Friends=3&Friends=4&Friends=5&Friends=6&User=moishe&age=100",
			"qs",
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test: %s", testCase.expectedOutput), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.expectedOutput,
				urlutils.QueryStringifyStruct(testCase.input, testCase.structTag),
			)
		})
	}
}

func TestParse(t *testing.T) {
	u, err := urlutils.Parse("https://example.com/path?key=value")
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "https", u.Scheme)
	assert.Equal(t, "example.com", u.Host)
	assert.Equal(t, "/path", u.Path)

	_, err = urlutils.Parse("://invalid")
	assert.Error(t, err)
}

func TestMustParse(t *testing.T) {
	u := urlutils.MustParse("https://example.com/path")
	assert.NotNil(t, u)
	assert.Equal(t, "https", u.Scheme)

	assert.Panics(t, func() {
		urlutils.MustParse("://invalid")
	})
}

func TestIsAbsolute(t *testing.T) {
	assert.True(t, urlutils.IsAbsolute("https://example.com"))
	assert.True(t, urlutils.IsAbsolute("http://example.com/path"))
	assert.False(t, urlutils.IsAbsolute("/path"))
	assert.False(t, urlutils.IsAbsolute("path"))
	assert.False(t, urlutils.IsAbsolute("://invalid"))
}

func TestGetDomain(t *testing.T) {
	assert.Equal(t, "example.com", urlutils.GetDomain("https://example.com/path"))
	assert.Equal(t, "example.com:8080", urlutils.GetDomain("https://example.com:8080/path"))
	assert.Equal(t, "", urlutils.GetDomain("/path"))
	assert.Equal(t, "", urlutils.GetDomain("://invalid"))
}

func TestGetScheme(t *testing.T) {
	assert.Equal(t, "https", urlutils.GetScheme("https://example.com"))
	assert.Equal(t, "http", urlutils.GetScheme("http://example.com"))
	assert.Equal(t, "", urlutils.GetScheme("/path"))
	assert.Equal(t, "", urlutils.GetScheme("://invalid"))
}

func TestGetPath(t *testing.T) {
	assert.Equal(t, "/path/to/resource", urlutils.GetPath("https://example.com/path/to/resource"))
	assert.Equal(t, "/", urlutils.GetPath("https://example.com/"))
	assert.Equal(t, "", urlutils.GetPath("https://example.com"))
	assert.Equal(t, "", urlutils.GetPath("://invalid"))
}

func TestAddQuery(t *testing.T) {
	u := urlutils.MustParse("https://example.com/path")
	result := urlutils.AddQuery(u, "key", "value")

	assert.Equal(t, "value", result.Query().Get("key"))
	// Original URL should be unchanged
	assert.Empty(t, u.RawQuery)

	// Add to existing query
	u2 := urlutils.MustParse("https://example.com/path?existing=param")
	result2 := urlutils.AddQuery(u2, "new", "value")
	assert.Equal(t, "param", result2.Query().Get("existing"))
	assert.Equal(t, "value", result2.Query().Get("new"))
}

func TestSetQuery(t *testing.T) {
	u := urlutils.MustParse("https://example.com/path?key=old")
	result := urlutils.SetQuery(u, "key", "new")

	assert.Equal(t, "new", result.Query().Get("key"))
	// Original URL should be unchanged
	assert.Equal(t, "old", u.Query().Get("key"))
}

func TestRemoveQuery(t *testing.T) {
	u := urlutils.MustParse("https://example.com/path?key1=value1&key2=value2")
	result := urlutils.RemoveQuery(u, "key1")

	assert.False(t, result.Query().Has("key1"))
	assert.True(t, result.Query().Has("key2"))
	// Original URL should be unchanged
	assert.True(t, u.Query().Has("key1"))
}

func TestGetQuery(t *testing.T) {
	u := urlutils.MustParse("https://example.com/path?key=value")
	assert.Equal(t, "value", urlutils.GetQuery(u, "key"))
	assert.Equal(t, "", urlutils.GetQuery(u, "nonexistent"))
}

func TestHasQuery(t *testing.T) {
	u := urlutils.MustParse("https://example.com/path?key=value")
	assert.True(t, urlutils.HasQuery(u, "key"))
	assert.False(t, urlutils.HasQuery(u, "nonexistent"))
}
