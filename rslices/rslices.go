package rslices

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

// Sum sums the values in a collection. If collection is empty 0 is returned.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	var sum T

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

// IsIntersect return ok or not for intersection slice.
func IsIntersect[E comparable](s1, s2 []E) bool {
	for _, i := range s1 {
		for _, j := range s2 {
			if i == j {
				return true
			}
		}
	}

	return false
}

// Intersection - intersection of two arrays. Returns new slice.
func Intersection[T cmp.Ordered](s1, s2 []T) []T {
	var (
		minimum = min(len(s1), len(s2))
		out     = make([]T, 0, minimum)
		check   = make(map[T]struct{}, minimum)
	)

	for _, i := range s1 {
		for _, j := range s2 {
			if i == j && check[i] == struct{}{} {
				out = append(out, i)
				check[i] = struct{}{}
			}
		}
	}

	return out
}

// SlicesConcat - concatenation of multiple slices.
func SlicesConcat[T any](slices ...[]T) []T {
	var length, idx int

	for _, slice := range slices {
		length += len(slice)
	}

	res := make([]T, length)

	for _, s := range slices {
		idx += copy(res[idx:], s)
	}

	return res
}
