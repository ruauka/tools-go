// Package rtime - time functions.
package rtime

import "time"

// Months - The number of months between two dates.
func Months(d1, d2 time.Time) int {
	var absOff = false

	// if date 1 is greater than 2, then swap them
	if d1.After(d2) {
		d2, d1 = d1, d2
		// disabling the calculation modulo
		absOff = true
	}

	d1Year, d1Month, _ := d1.Date()
	d2Year, d2Month, _ := d2.Date()

	result := (d1Year-d2Year)*12 + int(d1Month) - int(d2Month) + extraMonth(d1, d2)

	if absOff {
		return result
	}

	return absInt(result)
}

// ExtraMonth - calculation of the allowance in the form of one month.
func extraMonth(d1, d2 time.Time) int {
	_, d1Month, d1Day := d1.Date()
	_, d2Month, d2Day := d2.Date()

	if d1Day == d2Day {
		return 0
	}

	if isLastDayInMonth(d2) {
		return 0
	}

	if d1Day > d2Day {
		return 1
	}
	if isLastDayInMonth(d1) {
		if d1Month > d2Month {
			return 1
		}
	}

	return 0
}

// IsLastDayInMonth - Checking that the day is the last day of the month.
func isLastDayInMonth(d time.Time) bool {
	return d.Day() == lastDayInMonth(d)
}

// LastDayInMonth - the last day of the month.
func lastDayInMonth(d time.Time) int {
	year, month, _ := d.Date()
	switch month {
	case time.February:
		// checking the leap year for February
		if year%4 == 0 {
			return 29
		}
		return 28
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	default:
		return 30
	}
}

// absInt - module int.
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
