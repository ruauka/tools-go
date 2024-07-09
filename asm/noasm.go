//go:build !amd64

// Package asm - asm functions. Build for ARM64.
package asm

import "github.com/ruauka/tools-go/rslices"

// Sum32 sums the values []float32 if not support optimization in goasm.
func Sum32(x []float32) float32 {
	return rslices.Sum(x)
}

// Sum64 sums the values []float64 if not support optimization in goasm.
func Sum64(x []float64) float64 {
	return rslices.Sum(x)
}

// Mul32 multiply arguments []float32 if not support optimization in goasm.
func Mul32(x []float32, y []float32) {
	rslices.Mul(x, y)
}

// Mul64 multiply arguments []float64 if not support optimization in goasm.
func Mul64(x []float64, y []float64) {
	rslices.Mul(x, y)
}

// Mul64Simd multiply arguments []float64 element-wise with optimization simd in goasm.
func Mul64Simd(out []float64, x []float64, y []float64) {
	copy(out, x)
	rslices.Mul(out, y)
}

// MulNum32 multiply arguments []float32 if not support optimization in goasm.
func MulNum32(x []float32, a float32) {
	rslices.MulNum(x, a)
}

// MulNum64 multiply arguments []float64 if not support optimization in goasm.
func MulNum64(x []float64, a float64) {
	rslices.MulNum(x, a)
}

// Add32 add number from []float32 if not support optimization in goasm.
func Add32(x []float32, y []float32) {
	rslices.Add(x, y)
}

// Add64 add number from []float64 if not support optimization in goasm.
func Add64(x []float64, y []float64) {
	rslices.Add(x, y)
}

// AddNum32 add number float32 for each element in []float32 if not support optimization in goasm.
func AddNum32(x []float32, a float32) {
	rslices.AddNum(x, a)
}

// AddNum64 add number float64 for each element in []float32 if not support optimization in goasm.
func AddNum64(x []float64, a float64) {
	rslices.AddNum(x, a)
}

// MaximumNum32 element-wise maximum of []float32 elements with number if not support optimization in goasm.
func MaximumNum32(x []float32, a float32) {
	rslices.MaximumNum(x, a)
}

// MaximumNum64 element-wise maximum of []float64 elements with number if not support optimization in goasm.
func MaximumNum64(x []float64, a float64) {
	rslices.MaximumNum(x, a)
}
