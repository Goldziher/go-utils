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
		assert.Equal(t, "active=true&age=100&friends=1&friends=2&friends=3&friends=4&friends=5&friends=6&user=moishe", actualOutput)
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
			assert.Equal(t, testCase.expectedOutput, urlutils.QueryStringifyStruct(testCase.input, testCase.structTag))
		})
	}
}
