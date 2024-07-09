package rslices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersection(t *testing.T) {
	type args[T any] struct {
		left, right []T
	}

	type cases[T any] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesInts := []cases[int]{
		{
			args:     args[int]{left: []int{1, 2, 3}, right: []int{1, 2, 4}},
			expected: []int{1, 2},
			testName: "Ints. OK",
		},
	}
	for _, testCase := range testCasesInts {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Intersection(testCase.args.left, testCase.args.right)
			require.Equal(t, testCase.expected, actual)
		})
	}

	testCasesFloats := []cases[float64]{
		{
			args:     args[float64]{left: []float64{1, 2, 3}, right: []float64{1, 2, 4}},
			expected: []float64{1, 2},
			testName: "Floats. OK",
		},
	}
	for _, testCase := range testCasesFloats {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Intersection(testCase.args.left, testCase.args.right)
			require.Equal(t, testCase.expected, actual)
		})
	}

	testCasesStrs := []cases[string]{
		{
			args:     args[string]{left: []string{"aaa", "bbb", "ccc"}, right: []string{"aaa", "bbb", "ddd"}},
			expected: []string{"aaa", "bbb"},
			testName: "Strings. OK",
		},
	}
	for _, testCase := range testCasesStrs {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Intersection(testCase.args.left, testCase.args.right)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func ExampleIntersection() {
	var (
		intsL = []int{1, 2, 3}
		intsR = []int{1, 2, 4}

		floatsL = []float64{1.1, 2.2, 3.3}
		floatsR = []float64{1.1, 2.2, 4.4}

		strL = []string{"aaa", "bbb", "ccc"}
		strR = []string{"aaa", "bbb", "ddd"}
	)

	resInts := Intersection(intsL, intsR)
	fmt.Println(resInts)

	resFloats := Intersection(floatsL, floatsR)
	fmt.Println(resFloats)

	resStrs := Intersection(strL, strR)
	fmt.Println(resStrs)

	// Output:
	// [1 2]
	// [1.1 2.2]
	// [aaa bbb]
}

func TestConcat(t *testing.T) {
	testCasesInts := []struct {
		sl1, sl2, sl3 []int
		expected      []int
		testName      string
	}{
		{
			sl1:      []int{1, 2, 3},
			sl2:      []int{4, 5, 6},
			sl3:      []int{7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			testName: "OK. Ints",
		},
	}

	for _, testCase := range testCasesInts {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Concat(testCase.sl1, testCase.sl2, testCase.sl3)
			require.Equal(t, testCase.expected, actual)
		})
	}

	testCasesStrs := []struct {
		sl1, sl2, sl3 []string
		expected      []string
		testName      string
	}{
		{
			sl1:      []string{"1", "2", "3"},
			sl2:      []string{"4", "5", "6"},
			sl3:      []string{"7", "8", "9"},
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			testName: "OK. Strs",
		},
	}

	for _, testCase := range testCasesStrs {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Concat(testCase.sl1, testCase.sl2, testCase.sl3)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func ExampleConcat() {
	var (
		ints1 = []int{1, 2, 3}
		ints2 = []int{4, 5, 6}
		ints3 = []int{7, 8, 9}

		strs1 = []string{"1", "2", "3"}
		strs2 = []string{"4", "5", "6"}
		strs3 = []string{"7", "8", "9"}
	)

	ints := Concat(ints1, ints2, ints3)
	fmt.Println(ints)

	strs := Concat(strs1, strs2, strs3)
	fmt.Println(strs)

	// Output:
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
}
