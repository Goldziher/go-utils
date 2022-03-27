package maputils_test

import (
	"go-utils/assert"
	"go-utils/maputils"
	"sort"
	"strconv"
	"testing"
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

func TestReMap(t *testing.T) {
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
		"Sunday":   1,
		"2":        2,
		"Tuesday":  3,
		"4":        4,
		"Thursday": 5,
		"6":        6,
		"Saturday": 7,
	}

	remappedDays := maputils.ReMap(daysMap, func(key string, value int) string {
		if value%2 == 0 {
			return strconv.Itoa(value)
		}
		return key
	})

	assert.Equal(t, expectedResult, remappedDays)
	assert.Equal(t, expectedResult, daysMap)
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
	assert.Equal(t, expectedResult, maputils.Drop(daysMap, []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Saturday"}))
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
