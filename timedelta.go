package attrs_go //nolint:revive,stylecheck

import "time"

const (
	secondsPerMinute = 60
	minutesPerHour   = 60
	hoursPerDay      = 24
	daysPerYear      = 365
	daysPerLeapYear  = 366
	daysPerMonth     = 30
	monthsPerYear    = 12
)

// TimeDelta - the difference between the two dates for each value.
type TimeDelta struct {
	Years, Months, Days, Hours, Minutes, Seconds, Nanoseconds      int
	TotalMonths, TotalDays, TotalHours, TotalMinutes, TotalSeconds int
}

// isLeapYear - check for leap year.
func isLeapYear(year int) int {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return daysPerLeapYear
	}

	return daysPerYear
}

// daysIn - return a day in a certain month of the year.
func daysIn(year int, month time.Month) int {
	return time.Date(year, month, 0, 0, 0, 0, 0, time.UTC).Day()
}

// Elapsed - Get the time difference from 'from' to 'to'.
func Elapsed(from, to time.Time) *TimeDelta { //nolint:funlen
	if from.Location() != to.Location() {
		to = to.In(to.Location())
	}
	if from.After(to) {
		from, to = to, from
	}

	var (
		y1, M1, d1  = from.Date()
		y2, M2, d2  = to.Date()
		h1, m1, s1  = from.Clock()
		h2, m2, s2  = to.Clock()
		ns1, ns2    = from.Nanosecond(), to.Nanosecond()
		years       = y2 - y1
		months      = int(M2 - M1)
		days        = d2 - d1
		hours       = h2 - h1
		minutes     = m2 - m1
		seconds     = s2 - s1
		nanoseconds = ns2 - ns1
	)

	if nanoseconds < 0 {
		nanoseconds += 1e9
		seconds--
	}
	if seconds < 0 {
		seconds += 60
		minutes--
	}
	if minutes < 0 {
		minutes += 60
		hours--
	}
	if hours < 0 {
		hours += 24
		days--
	}
	if days < 0 {
		days += daysIn(y2, M2-1)
		months--
	}
	if days < 0 {
		days += daysIn(y2, M2)
		months--
	}
	if months < 0 {
		months += 12
		years--
	}

	var (
		totalMonths = years*monthsPerYear + months
		totalDays   int
	)

	if totalMonths%12 == 11 {
		totalDays = d2 - d1
		if years == 0 {
			totalDays += isLeapYear(y1)
		}
	} else {
		totalDays = (totalMonths-years*monthsPerYear)*daysPerMonth + days
	}

	for i := 1; i < years+1; i++ {
		totalDays += isLeapYear(y1 + i)
	}

	var (
		totalHours   = totalDays*hoursPerDay + hours
		totalMinutes = totalHours*minutesPerHour + minutes
		totalSeconds = totalMinutes*secondsPerMinute + seconds
	)

	return &TimeDelta{
		Years:        years,
		Months:       months,
		Days:         days,
		Hours:        hours,
		Minutes:      minutes,
		Seconds:      seconds,
		Nanoseconds:  nanoseconds,
		TotalMonths:  totalMonths,
		TotalDays:    totalDays,
		TotalHours:   totalHours,
		TotalMinutes: totalMinutes,
		TotalSeconds: totalSeconds,
	}
}
