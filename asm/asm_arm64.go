////go:build !amd64

package asm

import "github.com/ruauka/tools-go/rslices"

// Sum32 sums the values []float32 without optimization in goasm.
func Sum32(x []float32) float32 {
	return rslices.Sum(x)
}

// Sum64 sums the values []float64 without optimization in goasm.
func Sum64(x []float64) float64 {
	return rslices.Sum(x)
}
