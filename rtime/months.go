package rtime

import "time"

// Months - Возвращает кол-во месяцев между двумя датами.
func Months(d1, d2 time.Time) int {
	var absOff = false

	// если дата 1 больше 2, то меняем их местами
	if d1.After(d2) {
		d2, d1 = d1, d2
		// Отключаем расчет по модулю
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

// ExtraMonth - расчет надбавки в виде одного месяца.
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

// IsLastDayInMonth - Проверка, что день является последним днем месяца.
func isLastDayInMonth(d time.Time) bool {
	return d.Day() == lastDayInMonth(d)
}

// LastDayInMonth - возвращает последний день месяца.
func lastDayInMonth(d time.Time) int {
	year, month, _ := d.Date()
	switch month {
	case time.February:
		// Проверка высокосного года для февраля
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

// absInt - модуль int.
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
