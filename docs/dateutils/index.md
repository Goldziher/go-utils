# dateutils

Time and date utilities for business logic and date manipulation.

## Overview

The `dateutils` package provides utilities for working with dates and times, including business day calculations, date range operations, and common date manipulations.

## Functions

**Date Boundaries**: Floor, Ceil, StartOfDay, EndOfDay, StartOfWeek, EndOfWeek
**Month Operations**: GetFirstDayOfMonth, GetLastDayOfMonth, DaysInMonth
**Date Ranges**: Overlap, DaysBetween
**Business Logic**: AddBusinessDays, IsWeekend, IsWeekday
**Age Calculation**: Age, AgeAt
**Parsing**: ParseDate, ParseDateWithLayout, MustParseDate, MustParseDateWithLayout
**Comparison**: IsSameDay, IsSameMonth

## Example

```go
import (
    "time"
    "github.com/Goldziher/go-utils/dateutils"
)

now := time.Now()

// Date boundaries
startOfDay := dateutils.StartOfDay(now)      // Today at 00:00:00
endOfDay := dateutils.EndOfDay(now)          // Today at 23:59:59.999999999
startOfWeek := dateutils.StartOfWeek(now)    // This Sunday at 00:00:00

// Business day calculations
nextWeek := dateutils.AddBusinessDays(now, 5)  // 5 business days from now
isWeekend := dateutils.IsWeekend(now)          // true if Saturday/Sunday

// Date range operations
start1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
end1 := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
start2 := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
end2 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)

overlaps := dateutils.Overlap(start1, end1, start2, end2)  // true

// Age calculation
birthdate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
age := dateutils.Age(birthdate)  // Current age in years

// Month operations
firstDay := dateutils.GetFirstDayOfMonth()     // First day of current month
lastDay := dateutils.GetLastDayOfMonth()       // Last day of current month
daysInMonth := dateutils.DaysInMonth(now)      // Number of days in current month

// Date comparison
same := dateutils.IsSameDay(time.Now(), time.Now().Add(time.Hour))  // true
sameMonth := dateutils.IsSameMonth(start1, start2)  // true
```
