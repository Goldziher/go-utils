package dateutils_test

import (
	"testing"
	"time"

	"github.com/Goldziher/go-utils/dateutils"
	"github.com/stretchr/testify/assert"
)

func TestFloor(t *testing.T) {
	now := time.Now()
	assert.Equal(t, "00:00:00", dateutils.Floor(now).Format("15:04:05"))
}

func TestCeil(t *testing.T) {
	now := time.Now()
	assert.Equal(t, "23:59:59", dateutils.Ceil(now).Format("15:04:05"))
}

func TestOverlap(t *testing.T) {
	s1, _ := time.Parse("2006-01-02", "2022-12-28")
	e1, _ := time.Parse("2006-01-02", "2022-12-31")

	s2, _ := time.Parse("2006-01-02", "2022-12-30")
	e2, _ := time.Parse("2006-01-02", "2023-01-01")

	s3, _ := time.Parse("2006-01-02", "2023-01-02")
	e3, _ := time.Parse("2006-01-02", "2023-01-04")

	assert.Equal(t, true, dateutils.Overlap(s1, e1, s2, e2))
	assert.Equal(t, false, dateutils.Overlap(s1, e1, s3, e3))

	s4, _ := time.Parse("2006-01-02", "2023-07-13")
	e4, _ := time.Parse("2006-01-02", "2023-07-14")

	s5, _ := time.Parse("2006-01-02", "2023-07-10")
	e5, _ := time.Parse("2006-01-02", "2023-07-17")

	assert.Equal(t, true, dateutils.Overlap(s4, e4, s5, e5))
	assert.Equal(t, true, dateutils.Overlap(s5, e5, s4, e4))
}

func TestGetFirstDayOfMonth(t *testing.T) {
	result := dateutils.GetFirstDayOfMonth()

	now := time.Now()
	year, month, _ := now.Date()

	assert.Equal(t, year, result.Year())
	assert.Equal(t, month, result.Month())
	assert.Equal(t, 1, result.Day())
	assert.Equal(t, 0, result.Hour())
	assert.Equal(t, 0, result.Minute())
	assert.Equal(t, 0, result.Second())
}

func TestGetFirstDayOfMonthFor(t *testing.T) {
	date := time.Date(2023, 7, 15, 14, 30, 45, 0, time.UTC)
	result := dateutils.GetFirstDayOfMonthFor(date)

	assert.Equal(t, 2023, result.Year())
	assert.Equal(t, time.July, result.Month())
	assert.Equal(t, 1, result.Day())
	assert.Equal(t, 0, result.Hour())
	assert.Equal(t, 0, result.Minute())
	assert.Equal(t, 0, result.Second())
}

func TestGetLastDayOfMonth(t *testing.T) {
	// This test depends on the current date, so we test the general behavior
	result := dateutils.GetLastDayOfMonth()

	now := time.Now()
	year, month, _ := now.Date()

	assert.Equal(t, year, result.Year())
	assert.Equal(t, month, result.Month())
	// Last day should be >= 28 for any month
	assert.GreaterOrEqual(t, result.Day(), 28)
	assert.Equal(t, 23, result.Hour())
	assert.Equal(t, 59, result.Minute())
	assert.Equal(t, 59, result.Second())
}

func TestGetLastDayOfMonthFor(t *testing.T) {
	date := time.Date(2023, 7, 15, 14, 30, 45, 0, time.UTC)
	result := dateutils.GetLastDayOfMonthFor(date)

	assert.Equal(t, 2023, result.Year())
	assert.Equal(t, time.July, result.Month())
	assert.Equal(t, 31, result.Day())
	assert.Equal(t, 23, result.Hour())
	assert.Equal(t, 59, result.Minute())
	assert.Equal(t, 59, result.Second())

	// Test February in non-leap year
	febDate := time.Date(2023, 2, 15, 12, 0, 0, 0, time.UTC)
	febResult := dateutils.GetLastDayOfMonthFor(febDate)
	assert.Equal(t, 28, febResult.Day())

	// Test February in leap year
	febLeapDate := time.Date(2024, 2, 15, 12, 0, 0, 0, time.UTC)
	febLeapResult := dateutils.GetLastDayOfMonthFor(febLeapDate)
	assert.Equal(t, 29, febLeapResult.Day())
}

func TestParseDate(t *testing.T) {
	fallback := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("should parse valid RFC3339 date", func(t *testing.T) {
		result := dateutils.ParseDate("2023-07-15T14:30:45Z", fallback)
		assert.Equal(t, 2023, result.Year())
		assert.Equal(t, time.July, result.Month())
		assert.Equal(t, 15, result.Day())
	})

	t.Run("should return fallback for invalid date", func(t *testing.T) {
		result := dateutils.ParseDate("invalid-date", fallback)
		assert.Equal(t, fallback, result)
	})

	t.Run("should return fallback for empty string", func(t *testing.T) {
		result := dateutils.ParseDate("", fallback)
		assert.Equal(t, fallback, result)
	})
}

func TestParseDateWithLayout(t *testing.T) {
	fallback := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("should parse valid date with custom layout", func(t *testing.T) {
		result := dateutils.ParseDateWithLayout("2023-07-15", "2006-01-02", fallback)
		assert.Equal(t, 2023, result.Year())
		assert.Equal(t, time.July, result.Month())
		assert.Equal(t, 15, result.Day())
	})

	t.Run("should return fallback for invalid date", func(t *testing.T) {
		result := dateutils.ParseDateWithLayout("invalid", "2006-01-02", fallback)
		assert.Equal(t, fallback, result)
	})

	t.Run("should return fallback for wrong layout", func(t *testing.T) {
		result := dateutils.ParseDateWithLayout("2023-07-15", "02/01/2006", fallback)
		assert.Equal(t, fallback, result)
	})
}

func TestMustParseDate(t *testing.T) {
	t.Run("should parse valid RFC3339 date", func(t *testing.T) {
		result := dateutils.MustParseDate("2023-07-15T14:30:45Z")
		assert.Equal(t, 2023, result.Year())
		assert.Equal(t, time.July, result.Month())
		assert.Equal(t, 15, result.Day())
	})

	t.Run("should panic for invalid date", func(t *testing.T) {
		assert.Panics(t, func() {
			dateutils.MustParseDate("invalid-date")
		})
	})
}

func TestMustParseDateWithLayout(t *testing.T) {
	t.Run("should parse valid date with custom layout", func(t *testing.T) {
		result := dateutils.MustParseDateWithLayout("2023-07-15", "2006-01-02")
		assert.Equal(t, 2023, result.Year())
		assert.Equal(t, time.July, result.Month())
		assert.Equal(t, 15, result.Day())
	})

	t.Run("should panic for invalid date", func(t *testing.T) {
		assert.Panics(t, func() {
			dateutils.MustParseDateWithLayout("invalid", "2006-01-02")
		})
	})
}

func TestStartOfDay(t *testing.T) {
	input := time.Date(2023, 7, 15, 14, 30, 45, 123, time.UTC)
	result := dateutils.StartOfDay(input)
	assert.Equal(t, 2023, result.Year())
	assert.Equal(t, time.July, result.Month())
	assert.Equal(t, 15, result.Day())
	assert.Equal(t, 0, result.Hour())
	assert.Equal(t, 0, result.Minute())
	assert.Equal(t, 0, result.Second())
	assert.Equal(t, 0, result.Nanosecond())
}

func TestEndOfDay(t *testing.T) {
	input := time.Date(2023, 7, 15, 14, 30, 45, 123, time.UTC)
	result := dateutils.EndOfDay(input)
	assert.Equal(t, 2023, result.Year())
	assert.Equal(t, time.July, result.Month())
	assert.Equal(t, 15, result.Day())
	assert.Equal(t, 23, result.Hour())
	assert.Equal(t, 59, result.Minute())
	assert.Equal(t, 59, result.Second())
	assert.Equal(t, 999999999, result.Nanosecond())
}

func TestStartOfWeek(t *testing.T) {
	// Wednesday, July 19, 2023
	wednesday := time.Date(2023, 7, 19, 14, 30, 45, 0, time.UTC)
	result := dateutils.StartOfWeek(wednesday)
	// Should be Sunday, July 16, 2023 at 00:00:00
	assert.Equal(t, time.Sunday, result.Weekday())
	assert.Equal(t, 16, result.Day())
	assert.Equal(t, 0, result.Hour())
}

func TestEndOfWeek(t *testing.T) {
	// Wednesday, July 19, 2023
	wednesday := time.Date(2023, 7, 19, 14, 30, 45, 0, time.UTC)
	result := dateutils.EndOfWeek(wednesday)
	// Should be Saturday, July 22, 2023 at 23:59:59
	assert.Equal(t, time.Saturday, result.Weekday())
	assert.Equal(t, 22, result.Day())
	assert.Equal(t, 23, result.Hour())
	assert.Equal(t, 59, result.Minute())
	assert.Equal(t, 59, result.Second())
}

func TestDaysBetween(t *testing.T) {
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 9, dateutils.DaysBetween(start, end))

	// Negative result when end is before start
	assert.Equal(t, -9, dateutils.DaysBetween(end, start))

	// Same day
	assert.Equal(t, 0, dateutils.DaysBetween(start, start))
}

func TestIsWeekend(t *testing.T) {
	saturday := time.Date(2023, 7, 15, 12, 0, 0, 0, time.UTC) // Saturday
	sunday := time.Date(2023, 7, 16, 12, 0, 0, 0, time.UTC)   // Sunday
	monday := time.Date(2023, 7, 17, 12, 0, 0, 0, time.UTC)   // Monday

	assert.True(t, dateutils.IsWeekend(saturday))
	assert.True(t, dateutils.IsWeekend(sunday))
	assert.False(t, dateutils.IsWeekend(monday))
}

func TestIsWeekday(t *testing.T) {
	saturday := time.Date(2023, 7, 15, 12, 0, 0, 0, time.UTC) // Saturday
	monday := time.Date(2023, 7, 17, 12, 0, 0, 0, time.UTC)   // Monday
	friday := time.Date(2023, 7, 21, 12, 0, 0, 0, time.UTC)   // Friday

	assert.False(t, dateutils.IsWeekday(saturday))
	assert.True(t, dateutils.IsWeekday(monday))
	assert.True(t, dateutils.IsWeekday(friday))
}

func TestAddBusinessDays(t *testing.T) {
	// Starting on Friday, July 14, 2023
	friday := time.Date(2023, 7, 14, 12, 0, 0, 0, time.UTC)

	// Add 1 business day -> Monday, July 17
	result := dateutils.AddBusinessDays(friday, 1)
	assert.Equal(t, time.Monday, result.Weekday())
	assert.Equal(t, 17, result.Day())

	// Add 5 business days -> Friday, July 21
	result = dateutils.AddBusinessDays(friday, 5)
	assert.Equal(t, time.Friday, result.Weekday())
	assert.Equal(t, 21, result.Day())

	// Add 0 business days
	result = dateutils.AddBusinessDays(friday, 0)
	assert.Equal(t, friday, result)

	// Subtract business days
	monday := time.Date(2023, 7, 17, 12, 0, 0, 0, time.UTC)
	result = dateutils.AddBusinessDays(monday, -1)
	assert.Equal(t, time.Friday, result.Weekday())
	assert.Equal(t, 14, result.Day())
}

func TestAge(t *testing.T) {
	// Test that Age returns a reasonable value
	birthdate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	age := dateutils.Age(birthdate)

	// Age should be at least 24 (as of 2024) and less than 100 (reasonable upper bound)
	assert.GreaterOrEqual(t, age, 24)
	assert.Less(t, age, 100)
}

func TestAgeAt(t *testing.T) {
	birthdate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)

	// Age on exact birthday
	exactBirthday := time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 33, dateutils.AgeAt(birthdate, exactBirthday))

	// Age before birthday this year
	beforeBirthday := time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 32, dateutils.AgeAt(birthdate, beforeBirthday))

	// Age after birthday this year
	afterBirthday := time.Date(2023, 7, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 33, dateutils.AgeAt(birthdate, afterBirthday))

	// Edge case: same month, day before birthday
	dayBefore := time.Date(2023, 5, 14, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 32, dateutils.AgeAt(birthdate, dayBefore))
}

func TestDaysInMonth(t *testing.T) {
	// January has 31 days
	january := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 31, dateutils.DaysInMonth(january))

	// February has 28 days (non-leap year)
	february := time.Date(2023, 2, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 28, dateutils.DaysInMonth(february))

	// February has 29 days (leap year)
	februaryLeap := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 29, dateutils.DaysInMonth(februaryLeap))

	// April has 30 days
	april := time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 30, dateutils.DaysInMonth(april))
}

func TestIsSameDay(t *testing.T) {
	day1Morning := time.Date(2023, 7, 15, 8, 0, 0, 0, time.UTC)
	day1Evening := time.Date(2023, 7, 15, 20, 0, 0, 0, time.UTC)
	day2 := time.Date(2023, 7, 16, 8, 0, 0, 0, time.UTC)

	assert.True(t, dateutils.IsSameDay(day1Morning, day1Evening))
	assert.False(t, dateutils.IsSameDay(day1Morning, day2))
}

func TestIsSameMonth(t *testing.T) {
	july1 := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	july31 := time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC)
	august1 := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
	julyDifferentYear := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)

	assert.True(t, dateutils.IsSameMonth(july1, july31))
	assert.False(t, dateutils.IsSameMonth(july31, august1))
	assert.False(t, dateutils.IsSameMonth(july1, julyDifferentYear))
}
