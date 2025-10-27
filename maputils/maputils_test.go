package maputils_test

import (
	"sort"
	"strconv"
	"testing"

	"github.com/Goldziher/go-utils/maputils"
	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	var daysMap = map[string]int{
		"Sunday":    1,
		"Monday":    2,
		"Tuesday":   3,
		"Wednesday": 4,
		"Thursday":  5,
		"Friday":    6,
		"Saturday":  7,
	}
	keys := maputils.Keys(daysMap)
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	sort.Strings(days)
	sort.Strings(keys)
	assert.Equal(t, days, keys)
}

func TestValues(t *testing.T) {
	var daysMap = map[string]int{
		"Sunday":    1,
		"Monday":    2,
		"Tuesday":   3,
		"Wednesday": 4,
		"Thursday":  5,
		"Friday":    6,
		"Saturday":  7,
	}
	values := maputils.Values(daysMap)
	sort.Ints(values)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, values)
}

func TestMerge(t *testing.T) {
	first := map[string]string{
		"George": "Harrison",
		"Betty":  "Davis",
	}
	second := map[string]string{
		"Ronald": "Reagen",
		"Betty":  "Stewart",
	}

	assert.Equal(t,
		map[string]string{
			"George": "Harrison",
			"Betty":  "Stewart",
			"Ronald": "Reagen",
		}, maputils.Merge(first, second))
}

func TestForEach(t *testing.T) {
	var daysMap = map[string]int{
		"Sunday":    1,
		"Monday":    2,
		"Tuesday":   3,
		"Wednesday": 4,
		"Thursday":  5,
		"Friday":    6,
		"Saturday":  7,
	}

	sum := 0

	maputils.ForEach(daysMap, func(_ string, day int) {
		sum += day
	})

	assert.Equal(t, 28, sum)
}

func TestDrop(t *testing.T) {
	var daysMap = map[string]int{
		"Sunday":    1,
		"Monday":    2,
		"Tuesday":   3,
		"Wednesday": 4,
		"Thursday":  5,
		"Friday":    6,
		"Saturday":  7,
	}

	expectedResult := map[string]int{
		"Sunday": 1,
		"Friday": 6,
	}
	assert.Equal(
		t,
		expectedResult,
		maputils.Drop(daysMap, []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Saturday"}),
	)
	// ensure we do not modify the original value
	assert.Equal(t, expectedResult, daysMap)
}

func TestCopy(t *testing.T) {
	var daysMap = map[string]int{
		"Sunday":    1,
		"Monday":    2,
		"Tuesday":   3,
		"Wednesday": 4,
		"Thursday":  5,
		"Friday":    6,
		"Saturday":  7,
	}
	daysCopy := maputils.Copy(daysMap)
	assert.Equal(t, daysMap, daysCopy)
	delete(daysCopy, "Sunday")
	assert.NotEqual(t, daysMap, daysCopy)
}

func TestFilter(t *testing.T) {
	var daysMap = map[string]int{
		"Sunday":    1,
		"Monday":    2,
		"Tuesday":   3,
		"Wednesday": 4,
		"Thursday":  5,
		"Friday":    6,
		"Saturday":  7,
	}

	var expectedResult = map[string]int{
		"Monday":    2,
		"Wednesday": 4,
		"Friday":    6,
	}

	filteredDays := maputils.Filter(daysMap, func(_ string, value int) bool {
		return value%2 == 0
	})

	assert.Equal(t, expectedResult, filteredDays)
}

func TestMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}

	result := maputils.Map(input, func(k string, v int) string {
		return k + strconv.Itoa(v*2)
	})

	assert.Equal(t, "a2", result["a"])
	assert.Equal(t, "b4", result["b"])
	assert.Equal(t, "c6", result["c"])
}

func TestMapKeys(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}

	result := maputils.MapKeys(input, func(k string, v int) int {
		return v * 10
	})

	assert.Equal(t, 1, result[10])
	assert.Equal(t, 2, result[20])
	assert.Equal(t, 3, result[30])
}

func TestInvert(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}

	result := maputils.Invert(input)

	assert.Equal(t, "a", result[1])
	assert.Equal(t, "b", result[2])
	assert.Equal(t, "c", result[3])
	assert.Len(t, result, 3)
}

func TestPick(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	result := maputils.Pick(input, []string{"a", "c", "e"})

	assert.Equal(t, 1, result["a"])
	assert.Equal(t, 3, result["c"])
	assert.NotContains(t, result, "b")
	assert.NotContains(t, result, "d")
	assert.NotContains(t, result, "e")
	assert.Len(t, result, 2)
}

func TestOmit(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	result := maputils.Omit(input, []string{"b", "d"})

	assert.Equal(t, 1, result["a"])
	assert.Equal(t, 3, result["c"])
	assert.NotContains(t, result, "b")
	assert.NotContains(t, result, "d")
	assert.Len(t, result, 2)
}

func TestHas(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}

	assert.True(t, maputils.Has(input, "a"))
	assert.True(t, maputils.Has(input, "b"))
	assert.False(t, maputils.Has(input, "c"))
}

func TestGet(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}

	assert.Equal(t, 1, maputils.Get(input, "a", 0))
	assert.Equal(t, 2, maputils.Get(input, "b", 0))
	assert.Equal(t, 99, maputils.Get(input, "c", 99))
}

func TestToEntries(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}

	result := maputils.ToEntries(input)

	assert.Len(t, result, 2)

	// Check that entries exist (order is non-deterministic)
	hasA := false
	hasB := false
	for _, entry := range result {
		if entry[0] == "a" && entry[1] == 1 {
			hasA = true
		}
		if entry[0] == "b" && entry[1] == 2 {
			hasB = true
		}
	}
	assert.True(t, hasA)
	assert.True(t, hasB)
}

func TestFromEntries(t *testing.T) {
	entries := [][2]any{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	result := maputils.FromEntries[string, int](entries)

	assert.Equal(t, 1, result["a"])
	assert.Equal(t, 2, result["b"])
	assert.Equal(t, 3, result["c"])
	assert.Len(t, result, 3)
}

func TestMapGroupBy(t *testing.T) {
	input := map[string]int{
		"apple":  10,
		"banana": 5,
		"cherry": 15,
		"date":   5,
	}

	result := maputils.GroupBy(input, func(k string, v int) string {
		if v < 10 {
			return "small"
		}
		return "large"
	})

	assert.Len(t, result, 2)
	assert.Len(t, result["small"], 2)
	assert.Len(t, result["large"], 2)
	assert.Equal(t, 5, result["small"]["banana"])
	assert.Equal(t, 5, result["small"]["date"])
	assert.Equal(t, 10, result["large"]["apple"])
	assert.Equal(t, 15, result["large"]["cherry"])
}
