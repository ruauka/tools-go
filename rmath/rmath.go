package rmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

//var (
//	formatFloat = []string{"%.0f", "%.1f", "%.2f", "%.3f", "%.4f", "%.5f", "%.6f", "%.7f", "%.8f", "%.9f", "%.10f", "%.11f", "%.12f", "%.13f"}
//	poolBuf     = sync.Pool{
//		New: func() any {
//			buf := make([]byte, 0, 10)
//			return &buf
//		},
//	}
//)
//

// Round - float64 | float32 rounder to certain precision.
func Round[T constraints.Float](value T, prec int) T {
	return T(math.Round(float64(value)*(math.Pow10(prec))) / math.Pow10(prec))
}

// RoundUp returns up float64 with precision.
func RoundUp[T constraints.Float](value T, prec int) T {
	return T(math.Ceil(float64(value)*(math.Pow10(prec))) / math.Pow10(prec))
}

//// RoundFloor returns floor float64 with rounding value.
//func RoundFloor(value float64, rounding float64) float64 {
//	return math.Floor(value/rounding) * rounding
//}
//
//// RoundPy returns float64 with precision (python version).
//func RoundPy(x float64, prec int) float64 {
//	if prec > len(formatFloat)-1 {
//		return RoundPy(x, len(formatFloat)-1)
//	}
//
//	bufPtr := poolBuf.Get().(*[]byte)
//	buf := *bufPtr
//	buf = fmt.Appendf(buf, formatFloat[prec], x)
//
//	y := conv.BytesToFloat64BestEffort(buf)
//
//	*bufPtr = buf[:0]
//	poolBuf.Put(bufPtr)
//
//	return y
//}
