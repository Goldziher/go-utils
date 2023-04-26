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
