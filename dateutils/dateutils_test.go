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

func TestBeforeOrEqual(t *testing.T) {
	milestone, _ := time.Parse("2006-01-02", "2023-01-01")

	dBefore, _ := time.Parse("2006-01-02", "2022-12-31")
	dEqual, _ := time.Parse("2006-01-02", "2023-01-01")
	dAfter, _ := time.Parse("2006-01-02", "2023-01-31")

	assert.Equal(t, true, dateutils.BeforeOrEqual(milestone, dBefore))
	assert.Equal(t, true, dateutils.BeforeOrEqual(milestone, dEqual))
	assert.Equal(t, false, dateutils.BeforeOrEqual(milestone, dAfter))
}

func TestAfterOrEqual(t *testing.T) {
	milestone, _ := time.Parse("2006-01-02", "2023-01-01")

	dBefore, _ := time.Parse("2006-01-02", "2022-12-31")
	dEqual, _ := time.Parse("2006-01-02", "2023-01-01")
	dAfter, _ := time.Parse("2006-01-02", "2023-01-31")

	assert.Equal(t, false, dateutils.AfterOrEqual(milestone, dBefore))
	assert.Equal(t, true, dateutils.AfterOrEqual(milestone, dEqual))
	assert.Equal(t, true, dateutils.AfterOrEqual(milestone, dAfter))
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
