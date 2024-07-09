// Package asm - asm functions. Build for AMD64.
package asm

import (
	"unsafe"

	"github.com/ruauka/tools-go/internal/asm"
)

// Sum32 sums the values []float32 with optimization in goasm.
func Sum32(x []float32) float32 {
	return asm.Sum32(x)
}

// Sum64 sums the values []float64 with optimization in goasm.
func Sum64(x []float64) float64 {
	var (
		buf  = unsafe.Pointer(&x[0])
		xLen = unsafe.Pointer(uintptr(len(x)))
		res  float64
	)

	asm.Sum64(buf, xLen, unsafe.Pointer(&res))

	return res
}

// Mul32 multiply arguments []float32 element-wise with optimization in goasm.
func Mul32(x []float32, y []float32) {
	asm.Mul32(x, y)
}

// Mul64 multiply arguments []float64 element-wise with optimization in goasm.
func Mul64(x []float64, y []float64) {
	asm.Mul64(x, y)
}

// Mul64Simd multiply arguments []float64 element-wise with optimization simd in goasm.
func Mul64Simd(out []float64, x []float64, y []float64) {
	asm.Mul64Simd(out[:len(x)], x, y)
}

// MulNum32 multiply arguments []float32 element-wise with number with optimization in goasm.
func MulNum32(x []float32, a float32) {
	asm.MulNum32(x, a)
}

// MulNum64 multiply arguments []float64 element-wise with number with optimization in goasm.
func MulNum64(x []float64, a float64) {
	asm.MulNum64(x, a)
}

// Add32 add number from []float32 for each element in []float32 with optimization in goasm.
func Add32(x []float32, y []float32) {
	asm.Add32(x, y)
}

// Add64 add number from []float64 for each element in []float64 with optimization in goasm.
func Add64(x []float64, y []float64) {
	asm.Add64(x, y)
}

// AddNum32 add number float32 for each element in []float32 with optimization in goasm.
func AddNum32(x []float32, a float32) {
	asm.AddNum32(x, a)
}

// AddNum64 add number float64 for each element in []float32 with optimization in goasm.
func AddNum64(x []float64, a float64) {
	asm.AddNum64(x, a)
}

// MaximumNum32 element-wise maximum of []float32 elements with number with optimization in goasm.
func MaximumNum32(x []float32, a float32) {
	asm.MaximumNum32(x, a)
}

// MaximumNum64 element-wise maximum of []float64 elements with number with optimization in goasm.
func MaximumNum64(x []float64, a float64) {
	asm.MaximumNum64(x, a)
}
