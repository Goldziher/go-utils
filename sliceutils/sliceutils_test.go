package sliceutils_test

import (
	"go-utils/asserts"
	"go-utils/sliceutils"
	"strconv"
	"strings"
	"testing"
)

var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var lastNames = []string{"Jacobs", "Vin", "Jacobs", "Smith"}

func TestFilter(t *testing.T) {
	expectedResult := []int{0, 2, 4, 6, 8}
	actualResult := sliceutils.Filter(numerals, func(value int, _ int, _ []int) bool {
		return value%2 == 0
	})
	asserts.Equal(t, expectedResult, actualResult)
}

func TestForEach(t *testing.T) {
	result := 0
	sliceutils.ForEach(numerals, func(value int, _ int, _ []int) {
		result += value
	})
	asserts.Equal(t, 45, result)
}

func TestMap(t *testing.T) {
	expectedResult := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	actualResult := sliceutils.Map(numerals, func(value int, _ int, _ []int) string {
		return strconv.Itoa(value)
	})
	asserts.Equal(t, expectedResult, actualResult)
}

func TestReduce(t *testing.T) {
	expectedResult := map[string]string{"result": "0123456789"}
	actualResult := sliceutils.Reduce(
		numerals,
		func(acc map[string]string, cur int, _ int, _ []int) map[string]string {
			acc["result"] += strconv.Itoa(cur)
			return acc
		},
		map[string]string{"result": ""},
	)
	asserts.Equal(t, expectedResult, actualResult)
}

func TestFind(t *testing.T) {
	expectedResult := "Wednesday"
	actualResult := sliceutils.Find(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})
	asserts.Equal(t, expectedResult, *actualResult)
	asserts.Nil(t, sliceutils.Find(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Rishon")
	}))
}

func TestFindIndex(t *testing.T) {
	expectedResult := 3
	actualResult := sliceutils.FindIndex(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})
	asserts.Equal(t, expectedResult, actualResult)
	asserts.Equal(t, -1, sliceutils.FindIndex(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Rishon")
	}))
}

func TestFindIndexOf(t *testing.T) {
	expectedResult := 3
	actualResult := sliceutils.FindIndexOf(days, "Wednesday")
	asserts.Equal(t, expectedResult, actualResult)
	asserts.Equal(t, -1, sliceutils.FindIndexOf(days, "Rishon"))
}

func TestFindLastIndex(t *testing.T) {
	expectedResult := 2
	actualResult := sliceutils.FindLastIndex(lastNames, func(value string, index int, slice []string) bool {
		return value == "Jacobs"
	})
	asserts.Equal(t, expectedResult, actualResult)
	asserts.Equal(t, -1, sliceutils.FindLastIndex(lastNames, func(value string, index int, slice []string) bool {
		return value == "Hamudi"
	}))
}

func TestFindLastIndexOf(t *testing.T) {
	expectedResult := 2
	actualResult := sliceutils.FindLastIndexOf(lastNames, "Jacobs")
	asserts.Equal(t, expectedResult, actualResult)
	asserts.Equal(t, -1, sliceutils.FindLastIndexOf(lastNames, "Hamudi"))
}

func TestFindIndexes(t *testing.T) {
	expectedResult := []int{0, 2}
	actualResult := sliceutils.FindIndexes(lastNames, func(value string, index int, slice []string) bool {
		return value == "Jacobs"
	})
	asserts.Equal(t, expectedResult, actualResult)
	asserts.Nil(t, sliceutils.FindIndexes(lastNames, func(value string, index int, slice []string) bool {
		return value == "Hamudi"
	}))
}

func TestFindIndexesOf(t *testing.T) {
	expectedResult := []int{0, 2}
	actualResult := sliceutils.FindIndexesOf(lastNames, "Jacobs")
	asserts.Equal(t, expectedResult, actualResult)
	asserts.Nil(t, "Hamudi")
}

func TestIncludes(t *testing.T) {
	asserts.True(t, sliceutils.Includes(numerals, 1))
	asserts.False(t, sliceutils.Includes(numerals, 11))
}

func TestAny(t *testing.T) {
	asserts.True(t, sliceutils.Any(numerals, func(value int, _ int, _ []int) bool {
		return value%5 == 0
	}))
	asserts.False(t, sliceutils.Any(numerals, func(value int, _ int, _ []int) bool {
		return value == 11
	}))
}

func TestAll(t *testing.T) {
	asserts.True(t, sliceutils.All([]int{1, 1, 1}, func(value int, _ int, _ []int) bool {
		return value == 1
	}))
	asserts.False(t, sliceutils.All([]int{1, 1, 1, 2}, func(value int, _ int, _ []int) bool {
		return value == 1
	}))
}

func TestMerge(t *testing.T) {
	result := sliceutils.Merge(numerals[:5], numerals[5:])
	asserts.Equal(t, numerals, result)
}

func TestSum(t *testing.T) {
	result := sliceutils.Sum(numerals)
	asserts.Equal(t, 45, result)
}

func TestRemove(t *testing.T) {
	testSlice := []int{1, 2, 3}
	asserts.Equal(t, []int{2, 3}, sliceutils.Remove(testSlice, 0))
	asserts.Equal(t, []int{1, 3}, sliceutils.Remove(testSlice, 1))
	asserts.Equal(t, []int{1, 2}, sliceutils.Remove(testSlice, 2))
}

func TestInsert(t *testing.T) {
	testSlice := []int{1, 2}

	asserts.Equal(t, []int{3, 1, 2}, sliceutils.Insert(testSlice, 0, 3))
	asserts.Equal(t, []int{1, 3, 2}, sliceutils.Insert(testSlice, 1, 3))
	asserts.Equal(t, []int{1, 2, 3}, sliceutils.Insert(testSlice, 2, 3))
}

func TestIntersection(t *testing.T) {
	expectedResult := []int{3, 4, 5}

	first := []int{1, 2, 3, 4, 5}
	second := []int{2, 3, 4, 5, 6}
	third := []int{3, 4, 5, 6, 7}

	asserts.Equal(t, expectedResult, sliceutils.Intersection(first, second, third))
}

func TestDifference(t *testing.T) {
	expectedResult := []int{1, 7}

	first := []int{1, 2, 3, 4, 5}
	second := []int{2, 3, 4, 5, 6}
	third := []int{3, 4, 5, 6, 7}

	asserts.Equal(t, expectedResult, sliceutils.Difference(first, second, third))
}

func TestUnion(t *testing.T) {
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7}

	first := []int{1, 2, 3, 4, 5}
	second := []int{2, 3, 4, 5, 6}
	third := []int{3, 4, 5, 6, 7}

	asserts.Equal(t, expectedResult, sliceutils.Union(first, second, third))
}

func TestReverse(t *testing.T) {
	expectedResult := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	asserts.Equal(t, expectedResult, sliceutils.Reverse(numerals))
	// ensure does not modify the original
	asserts.Equal(t, expectedResult, sliceutils.Reverse(numerals))
}
