package asm

import "unsafe"

//go:noescape
func Sum32(x []float32) float32

//go:noescape
func Sum64(buf, len, res unsafe.Pointer)

//go:noescape
func Mul32(x []float32, y []float32)

//go:noescape
func Mul64(x []float64, y []float64)

//go:noescape
func Mul64Simd(out []float64, x []float64, y []float64)

//go:noescape
func MulNum32(x []float32, a float32)

//go:noescape
func MulNum64(x []float64, a float64)

//go:noescape
func Add32(x []float32, y []float32)

//go:noescape
func Add64(x []float64, y []float64)

//go:noescape
func AddNum32(x []float32, a float32)

//go:noescape
func AddNum64(x []float64, a float64)

//go:noescape
func MaximumNum32(x []float32, a float32)

//go:noescape
func MaximumNum64(x []float64, a float64)
