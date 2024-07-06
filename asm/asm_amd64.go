package asm

import "github.com/ruauka/tools-go/internal/asm"

// Sum32 sums the values []float32 with optimization in goasm.
func Sum32(x []float32) float32 {
	return asm.Sum32(x)
}
