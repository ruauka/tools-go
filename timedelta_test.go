package attrs_go

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var date = time.Date(2022, 10, 7, 0, 0, 0, 0, time.UTC)

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

type testTimedDelta struct {
	dateTo   time.Time
	name     string
	expected int
}

func TestElapsedYears(t *testing.T) {
	testCases := []testTimedDelta{
		{NewDate(2023, 10, 7), "Год ровно", 1},
		{NewDate(2023, 9, 7), "Меньше года", 0},
		{NewDate(2024, 11, 7), "Больше года", 2},
		{NewDate(2025, 11, 7), "Несколько лет", 3},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Elapsed(date, testCase.dateTo).Years
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestElapsedMonths(t *testing.T) {
	testCases := []testTimedDelta{
		{NewDate(2022, 10, 7), "Такая же дата", 0},
		{NewDate(2022, 11, 7), "Ровно месяц", 1},
		{NewDate(2022, 11, 8), "Больше месяца", 1},
		{NewDate(2023, 11, 7), "Несколько месяцев", 13},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Elapsed(date, testCase.dateTo).TotalMonths
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestElapsedDays(t *testing.T) {
	testCases := []testTimedDelta{
		{NewDate(2022, 10, 7), "Такая же дата", 0},
		{NewDate(2022, 10, 8), "Ровно день", 1},
		{NewDate(2022, 11, 7), "месяц", daysPerMonth},
		{NewDate(2022, 12, 7), "Несколько месяцев", 2 * daysPerMonth},
		{NewDate(2023, 10, 1), "Близко к полному году -5", 359},
		{NewDate(2023, 10, 6), "Близко к полному году -1", 364},
		{NewDate(2023, 11, 6), "Год с небольшим", 394},
		{NewDate(2023, 10, 7), "Ровно год", daysPerYear},
		{NewDate(2023, 10, 8), "Больше года", daysPerYear + 1},
		{NewDate(2025, 10, 7), "Несколько лет, включая вискосный", 2*daysPerYear + daysPerLeapYear},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Elapsed(date, testCase.dateTo).TotalDays
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestElapsedHours(t *testing.T) {
	testCases := []testTimedDelta{
		{NewDate(2022, 10, 7), "Такая же дата", 0},
		{NewDate(2022, 10, 8), "Ровно день", hoursPerDay},
		{NewDate(2022, 11, 7), "месяц", daysPerMonth * hoursPerDay},
		{NewDate(2022, 12, 7), "Несколько месяцев", 2 * daysPerMonth * hoursPerDay},
		{NewDate(2023, 10, 7), "Ровно год", daysPerYear * hoursPerDay},
		{NewDate(2023, 10, 8), "Больше года", (daysPerYear + 1) * hoursPerDay},
		{
			time.Date(2022, 10, 7, 0, 59, 59, 0, time.UTC),
			"Неполный час",
			0,
		},
		{
			time.Date(2022, 10, 7, 3, 1, 1, 1, time.UTC),
			"Часы с небольшим",
			3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Elapsed(date, testCase.dateTo).TotalHours
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestElapsedMinutes(t *testing.T) {
	testCases := []testTimedDelta{
		{NewDate(2022, 10, 7), "Такая же дата", 0},
		{NewDate(2022, 10, 8), "Ровно день", hoursPerDay * minutesPerHour},
		{NewDate(2022, 11, 7), "месяц", daysPerMonth * hoursPerDay * minutesPerHour},
		{NewDate(2022, 12, 7), "Несколько месяцев", 2 * daysPerMonth * hoursPerDay * minutesPerHour},
		{NewDate(2023, 10, 7), "Ровно год", daysPerYear * hoursPerDay * minutesPerHour},
		{NewDate(2023, 10, 8), "Больше года", (daysPerYear + 1) * hoursPerDay * minutesPerHour},
		{
			time.Date(2022, 10, 7, 0, 0, 59, 0, time.UTC),
			"Неполная минута",
			0,
		},
		{
			time.Date(2022, 10, 7, 0, 4, 4, 0, time.UTC),
			"Минуты с небольшим",
			4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Elapsed(date, testCase.dateTo).TotalMinutes
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestElapsedSeconds(t *testing.T) {
	testCases := []testTimedDelta{
		{NewDate(2022, 10, 7), "Такая же дата", 0},
		{NewDate(2022, 10, 8), "Ровно день", hoursPerDay * minutesPerHour * secondsPerMinute},
		{NewDate(2022,
			11, 7), "месяц", daysPerMonth * hoursPerDay * minutesPerHour * secondsPerMinute},
		{NewDate(2022, 12, 7), "Несколько месяцев", 2 * daysPerMonth * hoursPerDay * minutesPerHour * secondsPerMinute},
		{NewDate(2023, 10, 7), "Ровно год", daysPerYear * hoursPerDay * minutesPerHour * secondsPerMinute},
		{NewDate(2023, 10, 8), "Больше года", (daysPerYear + 1) * hoursPerDay * minutesPerHour * secondsPerMinute},
		{
			time.Date(2022, 10, 7, 0, 0, 0, 100, time.UTC),
			"Неполная секунда",
			0,
		},
		{
			time.Date(2022, 10, 7, 0, 0, 4, 100, time.UTC),
			"Секунды с небольшим",
			4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Elapsed(date, testCase.dateTo).TotalSeconds
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		name     string
		year     int
		expected int
	}{
		{"2024", 2024, daysPerLeapYear},
		{"2023", 2023, daysPerYear},
		{"2000", 2000, daysPerLeapYear},
		{"400", 400, daysPerLeapYear},
		{"300", 300, daysPerYear},
		{"200", 200, daysPerYear},
		{"100", 100, daysPerYear},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := isLeapYear(testCase.year)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestElapsed(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}

	testCases := []struct {
		name     string
		args     args
		expected *TimeDelta
	}{
		{
			name: "Разные локали + from > to",
			args: args{
				from: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			expected: &TimeDelta{},
		},
		{
			name: "Ветвления",
			args: args{
				from: time.Date(2022, 3, 31, 2, 2, 2, 2, time.UTC),
				to:   time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: &TimeDelta{
				Years:        0,
				Months:       8,
				Days:         30,
				Hours:        22,
				Minutes:      58,
				Seconds:      58,
				Nanoseconds:  999999999,
				TotalMonths:  8,
				TotalDays:    270,
				TotalHours:   6502,
				TotalMinutes: 390178,
				TotalSeconds: 23410738,
			},
		},
	}
	for _, testCase := range testCases {
		actual := Elapsed(testCase.args.from, testCase.args.to)
		require.Equal(t, testCase.expected, actual)
	}
}

func ExampleElapsed() {
	var (
		from = time.Date(2022, 5, 25, 1, 1, 1, 1, time.UTC)
		to   = time.Date(2023, 5, 25, 1, 1, 1, 1, time.UTC)
	)

	res := Elapsed(from, to)
	fmt.Println(res.TotalDays)

	// Output:
	// 365
}
