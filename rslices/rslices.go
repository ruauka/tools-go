// Package rslices - slices functions.
package rslices

import (
	"cmp"
	"slices"

	"golang.org/x/exp/constraints"
)

// Sum - sums the values in a collection. If collection is empty 0 is returned.
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

func MaximumNum[S ~[]E, E constraints.Float](x S, a E) {
	for i := 0; i < len(x); i++ {
		if x[i] < a {
			x[i] = 0
		}
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

func Concat[S ~[]E, E any](collections ...S) S {
	var size int

	for _, collection := range collections {
		size += len(collection)
		if size < 0 {
			panic("len out of range")
		}
	}

	newSl := slices.Grow[S](nil, size)

	for _, collection := range collections {
		newSl = append(newSl, collection...)
	}

	return newSl
}
