package date

import (
	"errors"
	"time"
)

// ErrInvalidDate is returned when year, month, and day do not represent
// a valid calendar date.
var ErrInvalidDate = errors.New("invalid calendar date")

// Date represents a calendar date without time-of-day or timezone semantics.
type Date struct {
	value time.Time
}

// New creates a valid calendar date.
func New(year int, month time.Month, day int) (Date, error) {
	if month < time.January || month > time.December {
		return Date{}, ErrInvalidDate
	}

	if day < 1 || day > 31 {
		return Date{}, ErrInvalidDate
	}

	value := time.Date(
		year,
		month,
		day,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	if value.Year() != year ||
		value.Month() != month ||
		value.Day() != day {
		return Date{}, ErrInvalidDate
	}

	return Date{value: value}, nil
}

// Year returns the year component.
func (d Date) Year() int {
	return d.value.Year()
}

// Month returns the month component.
func (d Date) Month() time.Month {
	return d.value.Month()
}

// Day returns the day component.
func (d Date) Day() int {
	return d.value.Day()
}

// Before reports whether d occurs before other.
func (d Date) Before(other Date) bool {
	return d.value.Before(other.value)
}

// After reports whether d occurs after other.
func (d Date) After(other Date) bool {
	return d.value.After(other.value)
}
