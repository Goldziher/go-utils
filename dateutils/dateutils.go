package dateutils

import "time"

// Floor - takes a datetime and return a datetime from the same day at 00:00:00 (UTC).
func Floor(t time.Time) time.Time {
	return t.UTC().Truncate(time.Hour * 24)
}

// Ceil - takes a datetime and return a datetime from the same day at 23:59:59 (UTC).
func Ceil(t time.Time) time.Time {
	// add 24 hours so that we are dealing with tomorrow's datetime
	// Floor
	// Substract one second and we have today at 23:59:59
	return Floor(t.Add(time.Hour * 24)).Add(time.Second * -1)
}

func BeforeOrEqual(milestone time.Time, date time.Time) bool {
	return date.UTC().Before(milestone) || date.UTC().Equal(milestone)
}

func AfterOrEqual(milestone time.Time, date time.Time) bool {
	return date.UTC().After(milestone) || date.UTC().Equal(milestone)
}

// Overlap - returns true if two date intervals overlap.
func Overlap(start1 time.Time, end1 time.Time, start2 time.Time, end2 time.Time) bool {
	return (AfterOrEqual(start2, start1) && BeforeOrEqual(end2, start1)) ||
		(AfterOrEqual(start2, end1) && BeforeOrEqual(end2, end1)) ||
		(AfterOrEqual(start1, start2) && BeforeOrEqual(end1, start2)) ||
		(AfterOrEqual(start1, end2) && BeforeOrEqual(end1, end2))
}

// GetFirstDayOfMonth returns the first day of the current month at 00:00:00 in the local timezone.
func GetFirstDayOfMonth() time.Time {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
}

// GetFirstDayOfMonthFor returns the first day of the month for the given date at 00:00:00 in the date's timezone.
func GetFirstDayOfMonthFor(date time.Time) time.Time {
	year, month, _ := date.Date()
	location := date.Location()

	return time.Date(year, month, 1, 0, 0, 0, 0, location)
}

// GetLastDayOfMonth returns the last day of the current month at 23:59:59 in the local timezone.
func GetLastDayOfMonth() time.Time {
	firstDay := GetFirstDayOfMonth()
	// Add one month and subtract one second
	return firstDay.AddDate(0, 1, 0).Add(-time.Second)
}

// GetLastDayOfMonthFor returns the last day of the month for the given date at 23:59:59 in the date's timezone.
func GetLastDayOfMonthFor(date time.Time) time.Time {
	firstDay := GetFirstDayOfMonthFor(date)
	// Add one month and subtract one second
	return firstDay.AddDate(0, 1, 0).Add(-time.Second)
}

// ParseDate parses a date string in RFC3339 format.
// If parsing fails, returns the fallback value.
func ParseDate(date string, fallback time.Time) time.Time {
	parsedDate, parseErr := time.Parse(time.RFC3339, date)
	if parseErr != nil {
		return fallback
	}
	return parsedDate
}

// ParseDateWithLayout parses a date string using the specified layout.
// If parsing fails, returns the fallback value.
func ParseDateWithLayout(date string, layout string, fallback time.Time) time.Time {
	parsedDate, parseErr := time.Parse(layout, date)
	if parseErr != nil {
		return fallback
	}
	return parsedDate
}

// MustParseDate parses a date string in RFC3339 format.
// Panics if parsing fails.
func MustParseDate(date string) time.Time {
	parsedDate, parseErr := time.Parse(time.RFC3339, date)
	if parseErr != nil {
		panic(parseErr)
	}
	return parsedDate
}

// MustParseDateWithLayout parses a date string using the specified layout.
// Panics if parsing fails.
func MustParseDateWithLayout(date string, layout string) time.Time {
	parsedDate, parseErr := time.Parse(layout, date)
	if parseErr != nil {
		panic(parseErr)
	}
	return parsedDate
}
