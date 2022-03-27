package maputils_test

import (
	"go-utils/asserts"
	"go-utils/maputils"
	"sort"
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
	asserts.Equal(t, days, keys)
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
	asserts.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, values)
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

	asserts.Equal(t,
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

	asserts.Equal(t, 28, sum)
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
	asserts.Equal(t, expectedResult, maputils.Drop(daysMap, []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Saturday"}))
	// ensure we do not modify the original value
	asserts.Equal(t, expectedResult, daysMap)
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
	asserts.Equal(t, daysMap, daysCopy)
	delete(daysCopy, "Sunday")
	asserts.NotEqual(t, daysMap, daysCopy)
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

	asserts.Equal(t, expectedResult, filteredDays)
}
