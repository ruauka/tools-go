package rtime

import (
	"math"
	"time"
)

// Days - the number of days between two dates.
func Days(d1, d2 time.Time) int {
	diff := d2.Sub(d1)
	// rounding up the number of months
	return int(math.Ceil(diff.Hours() / 24))
}
