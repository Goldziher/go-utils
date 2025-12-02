package sliceutils_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/Goldziher/go-utils/sliceutils"
	"github.com/stretchr/testify/assert"
)

type MyInt int

type Pluckable struct {
	Code  string
	Value string
}

var numerals = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var numeralsWithUserDefinedType = []MyInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var lastNames = []string{"Jacobs", "Vin", "Jacobs", "Smith"}

func TestFilter(t *testing.T) {
	expectedResult := []int{0, 2, 4, 6, 8}
	actualResult := sliceutils.Filter(numerals, func(value int, _ int, _ []int) bool {
		return value%2 == 0
	})
	assert.Equal(t, expectedResult, actualResult)
}

func TestForEach(t *testing.T) {
	result := 0
	sliceutils.ForEach(numerals, func(value int, _ int, _ []int) {
		result += value
	})
	assert.Equal(t, 45, result)
}

func TestMap(t *testing.T) {
	expectedResult := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	actualResult := sliceutils.Map(numerals, func(value int, _ int, _ []int) string {
		return strconv.Itoa(value)
	})
	assert.Equal(t, expectedResult, actualResult)

	assert.Nil(t, sliceutils.Map([]int{}, func(_ int, _ int, _ []int) string {
		return ""
	}))

	assert.Nil(t, sliceutils.Map([]int(nil), func(_ int, _ int, _ []int) string {
		return ""
	}))
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
	assert.Equal(t, expectedResult, actualResult)
}

func TestFind(t *testing.T) {
	expectedResult := "Wednesday"
	actualResult := sliceutils.Find(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})
	assert.Equal(t, expectedResult, *actualResult)
	assert.Nil(t, sliceutils.Find(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Rishon")
	}))
}

func TestFindIndex(t *testing.T) {
	expectedResult := 3
	actualResult := sliceutils.FindIndex(days, func(value string, index int, slice []string) bool {
		return strings.Contains(value, "Wed")
	})
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(
		t,
		-1,
		sliceutils.FindIndex(days, func(value string, index int, slice []string) bool {
			return strings.Contains(value, "Rishon")
		}),
	)
}

func TestFindIndexOf(t *testing.T) {
	expectedResult := 3
	actualResult := sliceutils.FindIndexOf(days, "Wednesday")
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, -1, sliceutils.FindIndexOf(days, "Rishon"))
}

func TestFindLastIndex(t *testing.T) {
	expectedResult := 2
	actualResult := sliceutils.FindLastIndex(
		lastNames,
		func(value string, index int, slice []string) bool {
			return value == "Jacobs"
		},
	)
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(
		t,
		-1,
		sliceutils.FindLastIndex(lastNames, func(value string, index int, slice []string) bool {
			return value == "Hamudi"
		}),
	)

	assert.Equal(
		t,
		0,
		sliceutils.FindLastIndex(days, func(value string, _ int, _ []string) bool {
			return value == "Sunday"
		}),
	)
}

func TestFindLastIndexOf(t *testing.T) {
	expectedResult := 2
	actualResult := sliceutils.FindLastIndexOf(lastNames, "Jacobs")
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, -1, sliceutils.FindLastIndexOf(lastNames, "Hamudi"))
}

func TestFindIndexes(t *testing.T) {
	expectedResult := []int{0, 2}
	actualResult := sliceutils.FindIndexes(
		lastNames,
		func(value string, index int, slice []string) bool {
			return value == "Jacobs"
		},
	)
	assert.Equal(t, expectedResult, actualResult)
	assert.Nil(
		t,
		sliceutils.FindIndexes(lastNames, func(value string, index int, slice []string) bool {
			return value == "Hamudi"
		}),
	)
}

func TestFindIndexesOf(t *testing.T) {
	expectedResult := []int{0, 2}
	actualResult := sliceutils.FindIndexesOf(lastNames, "Jacobs")
	assert.Equal(t, expectedResult, actualResult)
	assert.Nil(t, sliceutils.FindIndexesOf(lastNames, "Hamudi"))
}

func TestIncludes(t *testing.T) {
	assert.True(t, sliceutils.Includes(numerals, 1))
	assert.False(t, sliceutils.Includes(numerals, 11))
}

func TestAny(t *testing.T) {
	assert.True(t, sliceutils.Some(numerals, func(value int, _ int, _ []int) bool {
		return value%5 == 0
	}))
	assert.False(t, sliceutils.Some(numerals, func(value int, _ int, _ []int) bool {
		return value == 11
	}))
}

func TestAll(t *testing.T) {
	assert.True(t, sliceutils.Every([]int{1, 1, 1}, func(value int, _ int, _ []int) bool {
		return value == 1
	}))
	assert.False(t, sliceutils.Every([]int{1, 1, 1, 2}, func(value int, _ int, _ []int) bool {
		return value == 1
	}))
}

func TestMerge(t *testing.T) {
	result := sliceutils.Merge(numerals[:5], numerals[5:])
	assert.Equal(t, numerals, result)

	assert.Nil(t, sliceutils.Merge([]int(nil), []int(nil)))
	assert.Nil(t, sliceutils.Merge([]int{}, []int{}))
}

func TestSum(t *testing.T) {
	result := sliceutils.Sum(numerals)
	assert.Equal(t, 45, result)
}

func TestSum2(t *testing.T) {
	result := sliceutils.Sum(numeralsWithUserDefinedType)
	assert.Equal(t, MyInt(45), result)
}

func TestRemove(t *testing.T) {
	testSlice := []int{1, 2, 3}
	result := sliceutils.Remove(testSlice, 1)
	assert.Equal(t, []int{1, 3}, result)
	assert.Equal(t, []int{1, 2, 3}, testSlice)
	result = sliceutils.Remove(result, 1)
	assert.Equal(t, []int{1}, result)
	result = sliceutils.Remove(result, 3)
	assert.Equal(t, []int{1}, result)
	result = sliceutils.Remove(result, 0)
	assert.Equal(t, []int{}, result)
	result = sliceutils.Remove(result, 1)
	assert.Equal(t, []int{}, result)
}

func TestCopy(t *testing.T) {
	testSlice := []int{1, 2, 3}
	copiedSlice := sliceutils.Copy(testSlice)
	copiedSlice[0] = 2
	assert.NotEqual(t, testSlice, copiedSlice)
}

func TestInsert(t *testing.T) {
	testSlice := []int{1, 2}
	result := sliceutils.Insert(testSlice, 0, 3)
	assert.Equal(t, []int{3, 1, 2}, result)
	assert.NotEqual(t, testSlice, result)
	assert.Equal(t, []int{1, 3, 2}, sliceutils.Insert(testSlice, 1, 3))
	assert.Equal(t, []int{1, 2, 3}, sliceutils.Insert(testSlice, 2, 3))
}

func TestIntersection(t *testing.T) {
	expectedResult := []int{3, 4, 5}

	first := []int{1, 2, 3, 4, 5}
	second := []int{2, 3, 4, 5, 6}
	third := []int{3, 4, 5, 6, 7}

	assert.Equal(t, expectedResult, sliceutils.Intersection(first, second, third))
}

func TestDifference(t *testing.T) {
	expectedResult := []int{1, 7}

	first := []int{1, 2, 3, 4, 5}
	second := []int{2, 3, 4, 5, 6}
	third := []int{3, 4, 5, 6, 7}

	assert.Equal(t, expectedResult, sliceutils.Difference(first, second, third))
}

func TestUnion(t *testing.T) {
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7}

	first := []int{1, 2, 3, 4, 5}
	second := []int{2, 3, 4, 5, 6}
	third := []int{3, 4, 5, 6, 7}

	assert.Equal(t, expectedResult, sliceutils.Union(first, second, third))
}

func TestReverse(t *testing.T) {
	expectedResult := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	assert.Equal(t, expectedResult, sliceutils.Reverse(numerals))
	// ensure does not modify the original
	assert.Equal(t, expectedResult, sliceutils.Reverse(numerals))

	// test basic odd length case
	expectedResult = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.Equal(t, expectedResult, sliceutils.Reverse(numerals[1:]))
}

func TestUnique(t *testing.T) {
	duplicates := []int{6, 6, 6, 9, 0, 0, 0}
	expectedResult := []int{6, 9, 0}
	assert.Equal(t, expectedResult, sliceutils.Unique(duplicates))
	// Ensure original is unaltered
	assert.NotEqual(t, expectedResult, duplicates)
}

func TestChunk(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}, sliceutils.Chunk(numbers, 2))
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}, sliceutils.Chunk(numbers, 3))
}

func TestPluck(t *testing.T) {
	items := []Pluckable{
		{
			Code:  "azer",
			Value: "Azer",
		},
		{
			Code:  "tyuio",
			Value: "Tyuio",
		},
	}

	assert.Equal(
		t,
		[]string{"azer", "tyuio"},
		sliceutils.Pluck(items, func(item Pluckable) *string {
			return &item.Code
		}),
	)
	assert.Equal(
		t,
		[]string{"Azer", "Tyuio"},
		sliceutils.Pluck(items, func(item Pluckable) *string {
			return &item.Value
		}),
	)
}

func TestFlatten(t *testing.T) {
	items := [][]int{
		{1, 2, 3, 4},
		{5, 6},
		{7, 8},
		{9, 10, 11},
	}

	flattened := sliceutils.Flatten(items)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, flattened)

	assert.Nil(t, sliceutils.Flatten([][]int{}))
	assert.Nil(t, sliceutils.Flatten([][]int(nil)))

	assert.Nil(t, sliceutils.Flatten([][]int{{}, {}}))
	assert.Nil(t, sliceutils.Flatten([][]int{nil, nil}))

	assert.Nil(t, sliceutils.Flatten([][]int{{}, nil}))
}

func TestEnsureUniqueAndAppend(t *testing.T) {
	slice := []string{}
	item := "go-utils"

	slice = sliceutils.EnsureUniqueAndAppend(slice, item)
	assert.Equal(t, 1, len(slice))

	slice = sliceutils.EnsureUniqueAndAppend(slice, item)
	assert.Equal(t, 1, len(slice))
}

func TestFlatMap(t *testing.T) {
	items := []int{1, 2, 3, 4}

	flatMapped := sliceutils.FlatMap(items, func(value int, index int, slice []int) []int {
		return []int{value, value * 2}
	})

	assert.Equal(t, []int{1, 2, 2, 4, 3, 6, 4, 8}, flatMapped)
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
		City string
	}

	t.Run("should group by simple key", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6}
		grouped := sliceutils.GroupBy(items, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})

		assert.Len(t, grouped, 2)
		assert.ElementsMatch(t, []int{2, 4, 6}, grouped["even"])
		assert.ElementsMatch(t, []int{1, 3, 5}, grouped["odd"])
	})

	t.Run("should group structs by field", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30, City: "NYC"},
			{Name: "Bob", Age: 25, City: "LA"},
			{Name: "Charlie", Age: 30, City: "NYC"},
			{Name: "David", Age: 25, City: "SF"},
		}

		byAge := sliceutils.GroupBy(people, func(p Person) int {
			return p.Age
		})

		assert.Len(t, byAge, 2)
		assert.Len(t, byAge[30], 2)
		assert.Len(t, byAge[25], 2)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var items []int
		grouped := sliceutils.GroupBy(items, func(n int) string {
			return "key"
		})

		assert.Empty(t, grouped)
	})
}

func TestPartition(t *testing.T) {
	t.Run("should partition by predicate", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evens, odds := sliceutils.Partition(items, func(n int) bool {
			return n%2 == 0
		})

		assert.ElementsMatch(t, []int{2, 4, 6, 8, 10}, evens)
		assert.ElementsMatch(t, []int{1, 3, 5, 7, 9}, odds)
	})

	t.Run("should handle all true", func(t *testing.T) {
		items := []int{2, 4, 6, 8}
		evens, odds := sliceutils.Partition(items, func(n int) bool {
			return n%2 == 0
		})

		assert.ElementsMatch(t, []int{2, 4, 6, 8}, evens)
		assert.Empty(t, odds)
	})

	t.Run("should handle all false", func(t *testing.T) {
		items := []int{1, 3, 5, 7}
		evens, odds := sliceutils.Partition(items, func(n int) bool {
			return n%2 == 0
		})

		assert.Empty(t, evens)
		assert.ElementsMatch(t, []int{1, 3, 5, 7}, odds)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var items []int
		evens, odds := sliceutils.Partition(items, func(n int) bool {
			return n%2 == 0
		})

		assert.Empty(t, evens)
		assert.Empty(t, odds)
	})
}

func TestDistinctBy(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	t.Run("should get distinct by key", func(t *testing.T) {
		items := []int{1, 2, 3, 2, 4, 3, 5}
		distinct := sliceutils.DistinctBy(items, func(n int) int {
			return n
		})

		assert.Equal(t, []int{1, 2, 3, 4, 5}, distinct)
	})

	t.Run("should get distinct structs by field", func(t *testing.T) {
		people := []Person{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 1, Name: "Alice Duplicate"},
			{ID: 3, Name: "Charlie"},
			{ID: 2, Name: "Bob Duplicate"},
		}

		distinct := sliceutils.DistinctBy(people, func(p Person) int {
			return p.ID
		})

		assert.Len(t, distinct, 3)
		assert.Equal(t, 1, distinct[0].ID)
		assert.Equal(t, 2, distinct[1].ID)
		assert.Equal(t, 3, distinct[2].ID)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var items []int
		distinct := sliceutils.DistinctBy(items, func(n int) int {
			return n
		})

		assert.Empty(t, distinct)
	})

	t.Run("should preserve order of first occurrence", func(t *testing.T) {
		items := []string{"a", "b", "c", "a", "d", "b"}
		distinct := sliceutils.DistinctBy(items, func(s string) string {
			return s
		})

		assert.Equal(t, []string{"a", "b", "c", "d"}, distinct)
	})
}

func TestCountBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("should count by key", func(t *testing.T) {
		items := []int{1, 2, 3, 2, 4, 3, 5, 3}
		counts := sliceutils.CountBy(items, func(n int) int {
			return n
		})

		assert.Equal(t, 1, counts[1])
		assert.Equal(t, 2, counts[2])
		assert.Equal(t, 3, counts[3])
		assert.Equal(t, 1, counts[4])
		assert.Equal(t, 1, counts[5])
	})

	t.Run("should count structs by field", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 30},
			{Name: "David", Age: 25},
			{Name: "Eve", Age: 30},
		}

		counts := sliceutils.CountBy(people, func(p Person) int {
			return p.Age
		})

		assert.Equal(t, 3, counts[30])
		assert.Equal(t, 2, counts[25])
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var items []int
		counts := sliceutils.CountBy(items, func(n int) int {
			return n
		})

		assert.Empty(t, counts)
	})
}

func TestTake(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, sliceutils.Take([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, sliceutils.Take([]int{1, 2, 3, 4, 5}, 10))
	assert.Equal(t, []int{}, sliceutils.Take([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{}, sliceutils.Take([]int{1, 2, 3}, -1))
	assert.Equal(t, []int{}, sliceutils.Take([]int{}, 5))
}

func TestSkip(t *testing.T) {
	assert.Equal(t, []int{4, 5}, sliceutils.Skip([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, []int{}, sliceutils.Skip([]int{1, 2, 3, 4, 5}, 10))
	assert.Equal(t, []int{1, 2, 3}, sliceutils.Skip([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{1, 2, 3}, sliceutils.Skip([]int{1, 2, 3}, -1))
	assert.Equal(t, []int{}, sliceutils.Skip([]int{}, 5))
}

func TestTakeLast(t *testing.T) {
	assert.Equal(t, []int{4, 5}, sliceutils.TakeLast([]int{1, 2, 3, 4, 5}, 2))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, sliceutils.TakeLast([]int{1, 2, 3, 4, 5}, 10))
	assert.Equal(t, []int{}, sliceutils.TakeLast([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{}, sliceutils.TakeLast([]int{1, 2, 3}, -1))
	assert.Equal(t, []int{}, sliceutils.TakeLast([]int{}, 5))
}

func TestSkipLast(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, sliceutils.SkipLast([]int{1, 2, 3, 4, 5}, 2))
	assert.Equal(t, []int{}, sliceutils.SkipLast([]int{1, 2, 3, 4, 5}, 10))
	assert.Equal(t, []int{1, 2, 3}, sliceutils.SkipLast([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{1, 2, 3}, sliceutils.SkipLast([]int{1, 2, 3}, -1))
	assert.Equal(t, []int{}, sliceutils.SkipLast([]int{}, 5))
}

func TestTakeWhile(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, sliceutils.TakeWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n < 4
	}))
	assert.Equal(t, []int{}, sliceutils.TakeWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 10
	}))
	assert.Equal(
		t,
		[]int{1, 2, 3, 4, 5},
		sliceutils.TakeWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
			return n < 10
		}),
	)
}

func TestSkipWhile(t *testing.T) {
	assert.Equal(t, []int{4, 5}, sliceutils.SkipWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n < 4
	}))
	assert.Equal(t, []int{}, sliceutils.SkipWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n < 10
	}))
	assert.Equal(
		t,
		[]int{1, 2, 3, 4, 5},
		sliceutils.SkipWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
			return n > 10
		}),
	)
}

func TestMinBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	min := sliceutils.MinBy(people, func(p Person) int {
		return p.Age
	})

	assert.NotNil(t, min)
	assert.Equal(t, "Bob", min.Name)
	assert.Equal(t, 25, min.Age)

	var emptyPeople []Person
	assert.Nil(t, sliceutils.MinBy(emptyPeople, func(p Person) int {
		return p.Age
	}))
}

func TestMaxBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	max := sliceutils.MaxBy(people, func(p Person) int {
		return p.Age
	})

	assert.NotNil(t, max)
	assert.Equal(t, "Charlie", max.Name)
	assert.Equal(t, 35, max.Age)

	var emptyPeople []Person
	assert.Nil(t, sliceutils.MaxBy(emptyPeople, func(p Person) int {
		return p.Age
	}))
}

func TestHead(t *testing.T) {
	head := sliceutils.Head([]int{1, 2, 3, 4, 5})
	assert.NotNil(t, head)
	assert.Equal(t, 1, *head)

	assert.Nil(t, sliceutils.Head([]int{}))
}

func TestTail(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4, 5}, sliceutils.Tail([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{}, sliceutils.Tail([]int{1}))
	assert.Equal(t, []int{}, sliceutils.Tail([]int{}))
}

func TestLast(t *testing.T) {
	last := sliceutils.Last([]int{1, 2, 3, 4, 5})
	assert.NotNil(t, last)
	assert.Equal(t, 5, *last)

	assert.Nil(t, sliceutils.Last([]int{}))
}

func TestInitial(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, sliceutils.Initial([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{}, sliceutils.Initial([]int{1}))
	assert.Equal(t, []int{}, sliceutils.Initial([]int{}))
}

func TestCompact(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, sliceutils.Compact([]int{0, 1, 0, 2, 0, 3, 0}))
	assert.Equal(t, []string{"a", "b"}, sliceutils.Compact([]string{"", "a", "", "b", ""}))
	assert.Equal(t, []int{}, sliceutils.Compact([]int{0, 0, 0}))
	assert.Equal(t, []int{1, 2, 3}, sliceutils.Compact([]int{1, 2, 3}))
}

func TestWindows(t *testing.T) {
	windows := sliceutils.Windows([]int{1, 2, 3, 4, 5}, 3)
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}, windows)

	windows2 := sliceutils.Windows([]int{1, 2, 3}, 2)
	assert.Equal(t, [][]int{
		{1, 2},
		{2, 3},
	}, windows2)

	assert.Equal(t, [][]int{}, sliceutils.Windows([]int{1, 2, 3}, 5))
	assert.Equal(t, [][]int{}, sliceutils.Windows([]int{1, 2, 3}, 0))
	assert.Equal(t, [][]int{}, sliceutils.Windows([]int{1, 2, 3}, -1))
}
