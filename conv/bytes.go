package conv

import "unsafe"

// BytesToString convert []byte to string without allocation.
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// BytesToFloat64BestEffort parses floating-point number s.
//
// It is equivalent to strconv.ParseFloat(string(b), 64), but is faster.
//
// 0 is returned if the number cannot be parsed.
func BytesToFloat64BestEffort(b []byte) float64 {
	return StringToFloat64BestEffort(BytesToString(b))
}
