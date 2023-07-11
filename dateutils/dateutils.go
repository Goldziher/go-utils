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
