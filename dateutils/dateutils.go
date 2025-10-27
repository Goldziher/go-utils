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

// StartOfDay returns the start of the day (00:00:00) for the given time in its timezone.
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day (23:59:59.999999999) for the given time in its timezone.
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns the start of the week (Sunday at 00:00:00) for the given time.
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	return StartOfDay(t.AddDate(0, 0, -weekday))
}

// EndOfWeek returns the end of the week (Saturday at 23:59:59.999999999) for the given time.
func EndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	daysUntilSaturday := 6 - weekday
	return EndOfDay(t.AddDate(0, 0, daysUntilSaturday))
}

// DaysBetween returns the number of days between two dates.
// The result is negative if end is before start.
func DaysBetween(start, end time.Time) int {
	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

// IsWeekend returns true if the given time is on a weekend (Saturday or Sunday).
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday returns true if the given time is on a weekday (Monday through Friday).
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// AddBusinessDays adds the specified number of business days to the given time.
// Business days are Monday through Friday, skipping weekends.
// Negative values subtract business days.
func AddBusinessDays(t time.Time, days int) time.Time {
	if days == 0 {
		return t
	}

	direction := 1
	if days < 0 {
		direction = -1
		days = -days
	}

	result := t
	for i := 0; i < days; {
		result = result.AddDate(0, 0, direction)
		if IsWeekday(result) {
			i++
		}
	}
	return result
}

// Age calculates the age in years from the given birthdate to now.
func Age(birthdate time.Time) int {
	return AgeAt(birthdate, time.Now())
}

// AgeAt calculates the age in years from the given birthdate to the specified date.
func AgeAt(birthdate, at time.Time) int {
	age := at.Year() - birthdate.Year()
	if at.Month() < birthdate.Month() ||
		(at.Month() == birthdate.Month() && at.Day() < birthdate.Day()) {
		age--
	}
	return age
}

// DaysInMonth returns the number of days in the month of the given time.
func DaysInMonth(t time.Time) int {
	// Get the first day of next month, then subtract one day
	year, month, _ := t.Date()
	firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())
	lastOfThisMonth := firstOfNextMonth.AddDate(0, 0, -1)
	return lastOfThisMonth.Day()
}

// IsSameDay returns true if both times are on the same day (same year, month, and day).
func IsSameDay(a, b time.Time) bool {
	aYear, aMonth, aDay := a.Date()
	bYear, bMonth, bDay := b.Date()
	return aYear == bYear && aMonth == bMonth && aDay == bDay
}

// IsSameMonth returns true if both times are in the same month (same year and month).
func IsSameMonth(a, b time.Time) bool {
	aYear, aMonth, _ := a.Date()
	bYear, bMonth, _ := b.Date()
	return aYear == bYear && aMonth == bMonth
}

// IsSameYear returns true if both times are in the same year.
func IsSameYear(a, b time.Time) bool {
	return a.Year() == b.Year()
}
