package rslices

import "golang.org/x/exp/constraints"

// Sum sums the values in a collection. If collection is empty 0 is returned.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	var sum T = 0
	for _, val := range collection {
		sum += val
	}
	return sum
}

func Mul[T constraints.Float](x, y []T) {
	for i := 0; i < len(x); i++ {
		x[i] *= y[i]
	}
}

func MulNum[S ~[]E, E constraints.Float](x S, a E) {
	for i := 0; i < len(x); i++ {
		x[i] *= a
	}
}

func Add[T constraints.Float](x, y []T) {
	for i := 0; i < len(x); i++ {
		x[i] += y[i]
	}
}

func AddNum[S ~[]E, E constraints.Float](x S, a E) {
	for i := 0; i < len(x); i++ {
		x[i] += a
	}
}

func MaximumNum[S ~[]E, E constraints.Float](x S, a E) {
	for i := 0; i < len(x); i++ {
		if x[i] < a {
			x[i] = 0
		}
	}
}
