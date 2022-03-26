package maputils_test

import (
	"go-utils/assert"
	"go-utils/maputils"
	"sort"
	"testing"
)

var daysMap = map[string]int{
	"Sunday":    1,
	"Monday":    2,
	"Tuesday":   3,
	"Wednesday": 4,
	"Thursday":  5,
	"Friday":    6,
	"Saturday":  7,
}

func TestKeys(t *testing.T) {
	keys := maputils.Keys(daysMap)
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	sort.Strings(days)
	sort.Strings(keys)
	assert.Equal(t, days, keys)
}

func TestValues(t *testing.T) {
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
