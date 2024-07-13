package rslices

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"golang.org/x/exp/constraints"
)

func TestSum(t *testing.T) {
	type cases[T constraints.Float | constraints.Integer | constraints.Complex] struct {
		s        []T
		expected T
		testName string
	}

	testCasesInts := []cases[int]{
		{
			s:        []int{1, 2, 3},
			expected: 6,
			testName: "OK. Ints",
		},
	}
	for _, testCase := range testCasesInts {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Sum")
			t.Description("Check func `Sum`. Slice of ints")
			t.WithParameters(
				allure.NewParameter("s", testCase.s),
			)

			actual := Sum(testCase.s)

			t.Assert().Equal(testCase.expected, actual, "Sum of Int slice")
		})
	}

	testCasesFloats := []cases[float64]{
		{
			s:        []float64{1, 2, 3},
			expected: 6,
			testName: "OK. Float",
		},
	}
	for _, testCase := range testCasesFloats {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Sum")
			t.Description("Check func `Sum`. Slice of Float")
			t.WithParameters(
				allure.NewParameter("s", testCase.s),
			)

			actual := Sum(testCase.s)

			t.Assert().Equal(testCase.expected, actual, "Sum of Float slice")
		})
	}
}

func ExampleSum() {
	var (
		s32 = []float32{1, 2, 3}
		s64 = []float64{1, 2, 3}
	)

	resFloats32 := Sum(s32)
	fmt.Println(resFloats32)

	resFloats64 := Sum(s64)
	fmt.Println(resFloats64)

	// Output:
	// 6
	// 6
}

func TestMul(t *testing.T) {
	type args[T constraints.Float] struct {
		s1, s2 []T
	}

	type cases[T constraints.Float] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesFloat32 := []cases[float32]{
		{
			args:     args[float32]{s1: []float32{1, 2, 3}, s2: []float32{2, 2, 2}},
			expected: []float32{2, 4, 6},
			testName: "OK. Float32",
		},
	}
	for _, testCase := range testCasesFloat32 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Mul")
			t.Description("Check func `Mul`. Slices of float32")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			Mul(testCase.args.s1, testCase.args.s2)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "Mul of two Float32 slices")
		})
	}

	testCasesFloat64 := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, s2: []float64{2, 2, 2}},
			expected: []float64{2, 4, 6},
			testName: "OK. Float64",
		},
	}
	for _, testCase := range testCasesFloat64 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Mul")
			t.Description("Check func `Mul`. Slices of float64")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			Mul(testCase.args.s1, testCase.args.s2)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "Mul of two Float64 slices")
		})
	}
}

func ExampleMul() {
	var (
		s32_1 = []float32{1, 2, 3}
		s32_2 = []float32{2, 2, 2}

		s64_1 = []float64{1, 2, 3}
		s64_2 = []float64{2, 2, 2}
	)

	Mul(s32_1, s32_2)
	fmt.Println(s32_1)

	Mul(s64_1, s64_2)
	fmt.Println(s64_1)

	// Output:
	// [2 4 6]
	// [2 4 6]
}

func TestMulNum(t *testing.T) {
	type args[T constraints.Float] struct {
		s1  []T
		num T
	}

	type cases[T constraints.Float] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesFloat32 := []cases[float32]{
		{
			args:     args[float32]{s1: []float32{1, 2, 3}, num: 2},
			expected: []float32{2, 4, 6},
			testName: "OK. Float32",
		},
	}
	for _, testCase := range testCasesFloat32 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("MulNum")
			t.Description("Check func `MulNum`. Slice of float32")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("num", testCase.args.num),
			)

			MulNum(testCase.args.s1, testCase.args.num)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "MulNum on Float32 slice")
		})
	}

	testCasesFloat64 := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, num: 2},
			expected: []float64{2, 4, 6},
			testName: "OK. Float64",
		},
	}
	for _, testCase := range testCasesFloat64 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("MulNum")
			t.Description("Check func `MulNum`. Slice of float64")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("num", testCase.args.num),
			)

			MulNum(testCase.args.s1, testCase.args.num)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "Mul on Float64 slices")
		})
	}
}

func ExampleMulNum() {
	var (
		s32_1          = []float32{1, 2, 3}
		num_32 float32 = 2

		s64_1          = []float64{1, 2, 3}
		num_64 float64 = 2
	)

	MulNum(s32_1, num_32)
	fmt.Println(s32_1)

	MulNum(s64_1, num_64)
	fmt.Println(s64_1)

	// Output:
	// [2 4 6]
	// [2 4 6]
}

func TestAdd(t *testing.T) {
	type args[T constraints.Float] struct {
		s1, s2 []T
	}

	type cases[T constraints.Float] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesFloat32 := []cases[float32]{
		{
			args:     args[float32]{s1: []float32{1, 2, 3}, s2: []float32{2, 2, 2}},
			expected: []float32{3, 4, 5},
			testName: "OK. Float32",
		},
	}
	for _, testCase := range testCasesFloat32 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Add")
			t.Description("Check func `Add`. Slices of float32")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			Add(testCase.args.s1, testCase.args.s2)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "Add of two Float32 slices")
		})
	}

	testCasesFloat64 := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, s2: []float64{2, 2, 2}},
			expected: []float64{3, 4, 5},
			testName: "OK. Float64",
		},
	}
	for _, testCase := range testCasesFloat64 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Add")
			t.Description("Check func `Add`. Slices of float32")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			Add(testCase.args.s1, testCase.args.s2)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "Add of two Float32 slices")
		})
	}
}

func ExampleAdd() {
	var (
		s32_1 = []float32{1, 2, 3}
		s32_2 = []float32{2, 2, 2}

		s64_1 = []float64{1, 2, 3}
		s64_2 = []float64{2, 2, 2}
	)

	Add(s32_1, s32_2)
	fmt.Println(s32_1)

	Add(s64_1, s64_2)
	fmt.Println(s64_1)

	// Output:
	// [3 4 5]
	// [3 4 5]
}

func TestAddNum(t *testing.T) {
	type args[T constraints.Float] struct {
		s1  []T
		num T
	}

	type cases[T constraints.Float] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesFloat32 := []cases[float32]{
		{
			args:     args[float32]{s1: []float32{1, 2, 3}, num: 2},
			expected: []float32{3, 4, 5},
			testName: "OK. Float32",
		},
	}
	for _, testCase := range testCasesFloat32 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("AddNum")
			t.Description("Check func `AddNum`. Slices of float32")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("num", testCase.args.num),
			)

			AddNum(testCase.args.s1, testCase.args.num)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "AddNum on Float32 slice")
		})
	}

	testCasesFloat64 := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, num: 2},
			expected: []float64{3, 4, 5},
			testName: "OK. Float64",
		},
	}
	for _, testCase := range testCasesFloat64 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("AddNum")
			t.Description("Check func `AddNum`. Slice of float64")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("num", testCase.args.num),
			)

			AddNum(testCase.args.s1, testCase.args.num)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "AddNum on Float64 slice")
		})
	}
}

func ExampleAddNum() {
	var (
		s32_1          = []float32{1, 2, 3}
		num_32 float32 = 2

		s64_1          = []float64{1, 2, 3}
		num_64 float64 = 2
	)

	AddNum(s32_1, num_32)
	fmt.Println(s32_1)

	AddNum(s64_1, num_64)
	fmt.Println(s64_1)

	// Output:
	// [3 4 5]
	// [3 4 5]
}

func TestMaximumNum(t *testing.T) {
	type args[T constraints.Float] struct {
		s1  []T
		num T
	}

	type cases[T constraints.Float] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesFloat32 := []cases[float32]{
		{
			args:     args[float32]{s1: []float32{1, 2, 3}, num: 2},
			expected: []float32{0, 2, 3},
			testName: "OK. Float32",
		},
	}
	for _, testCase := range testCasesFloat32 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("MaximumNum")
			t.Description("Check func `MaximumNum`. Slice of float32")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("num", testCase.args.num),
			)

			MaximumNum(testCase.args.s1, testCase.args.num)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "MaximumNum on Float32 slice")
		})
	}

	testCasesFloat64 := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, num: 2},
			expected: []float64{0, 2, 3},
			testName: "OK. Float64",
		},
	}
	for _, testCase := range testCasesFloat64 {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("MaximumNum")
			t.Description("Check func `MaximumNum`. Slice of float64")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("num", testCase.args.num),
			)

			MaximumNum(testCase.args.s1, testCase.args.num)

			actual := testCase.args.s1

			t.Assert().Equal(testCase.expected, actual, "MaximumNum on Float64 slice")
		})
	}
}

func ExampleMaximumNum() {
	var (
		s32_1          = []float32{1, 2, 3}
		num_32 float32 = 2

		s64_1          = []float64{1, 2, 3}
		num_64 float64 = 2
	)

	MaximumNum(s32_1, num_32)
	fmt.Println(s32_1)

	MaximumNum(s64_1, num_64)
	fmt.Println(s64_1)

	// Output:
	// [0 2 3]
	// [0 2 3]
}

func TestIsIntersect(t *testing.T) {
	type args[T any] struct {
		s1, s2 []T
	}

	type cases[T any] struct {
		args     args[T]
		expected bool
		testName string
	}

	testCasesInts := []cases[int]{
		{
			args:     args[int]{s1: []int{1, 2, 3}, s2: []int{3, 4, 5}},
			expected: true,
			testName: "OK. Ints",
		},
	}
	for _, testCase := range testCasesInts {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("IsIntersect")
			t.Description("Check func `IsIntersect`. Slices of ints")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			actual := IsIntersect(testCase.args.s1, testCase.args.s2)

			t.Assert().Equal(testCase.expected, actual, "Intersections of two Int slices")
		})
	}

	testCasesFloats := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, s2: []float64{3, 4, 5}},
			expected: true,
			testName: "OK. Floats",
		},
	}
	for _, testCase := range testCasesFloats {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("IsIntersect")
			t.Description("Check func `IsIntersect`. Slices of float64")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			actual := IsIntersect(testCase.args.s1, testCase.args.s2)

			t.Assert().Equal(testCase.expected, actual, "Intersections of two Float64 slices")
		})
	}

	testCasesStrs := []cases[string]{
		{
			args:     args[string]{s1: []string{"aaa", "bbb", "ccc"}, s2: []string{"aaa", "bbb", "ddd"}},
			expected: true,
			testName: "OK. Strings",
		},
	}
	for _, testCase := range testCasesStrs {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("IsIntersect")
			t.Description("Check func `IsIntersect`. Slices of strings")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			actual := IsIntersect(testCase.args.s1, testCase.args.s2)

			t.Assert().Equal(testCase.expected, actual, "Intersections of two Strings slices")
		})
	}
}

func ExampleIsIntersect() {
	var (
		ints1 = []int{1, 2, 3}
		ints2 = []int{3, 4, 5}

		floats1 = []float64{1.1, 2.2, 3.3}
		floats2 = []float64{11.1, 22.2, 44.4}

		str1 = []string{"aaa", "bbb", "ccc"}
		str2 = []string{"aaa", "bbb", "ddd"}
	)

	resInts := IsIntersect(ints1, ints2)
	fmt.Println(resInts)

	resFloats := IsIntersect(floats1, floats2)
	fmt.Println(resFloats)

	resStrs := IsIntersect(str1, str2)
	fmt.Println(resStrs)

	// Output:
	// true
	// false
	// true
}

func TestIntersection(t *testing.T) {
	type args[T any] struct {
		s1, s2 []T
	}

	type cases[T any] struct {
		args     args[T]
		expected []T
		testName string
	}

	testCasesInts := []cases[int]{
		{
			args:     args[int]{s1: []int{1, 2, 3}, s2: []int{1, 2, 4}},
			expected: []int{1, 2},
			testName: "OK. Ints",
		},
	}
	for _, testCase := range testCasesInts {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Intersection")
			t.Description("Check func `Intersection`. Slices of ints")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			actual := Intersection(testCase.args.s1, testCase.args.s2)

			t.Assert().Equal(testCase.expected, actual, "Intersections of two Int slices")
		})
	}

	testCasesFloats := []cases[float64]{
		{
			args:     args[float64]{s1: []float64{1, 2, 3}, s2: []float64{1, 2, 4}},
			expected: []float64{1, 2},
			testName: "OK. Floats",
		},
	}
	for _, testCase := range testCasesFloats {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Intersection")
			t.Description("Check func `Intersection`. Slices of float64")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			actual := Intersection(testCase.args.s1, testCase.args.s2)

			t.Assert().Equal(testCase.expected, actual, "Intersections of two Float64 slices")
		})
	}

	testCasesStrs := []cases[string]{
		{
			args:     args[string]{s1: []string{"aaa", "bbb", "ccc"}, s2: []string{"aaa", "bbb", "ddd"}},
			expected: []string{"aaa", "bbb"},
			testName: "OK. Strings",
		},
	}
	for _, testCase := range testCasesStrs {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Intersection")
			t.Description("Check func `Intersection`. Slices of strings")
			t.WithParameters(
				allure.NewParameter("s1", testCase.args.s1),
				allure.NewParameter("s2", testCase.args.s2),
			)

			actual := Intersection(testCase.args.s1, testCase.args.s2)

			t.Assert().Equal(testCase.expected, actual, "Intersections of two Strings slices")
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
		s1, s2, s3 []int
		expected   []int
		testName   string
	}{
		{
			s1:       []int{1, 2, 3},
			s2:       []int{4, 5, 6},
			s3:       []int{7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			testName: "OK. Ints",
		},
	}
	for _, testCase := range testCasesInts {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Concat")
			t.Description("Check func `Concat`. Slices of Ints")
			t.WithParameters(
				allure.NewParameter("s1", testCase.s1),
				allure.NewParameter("s2", testCase.s2),
				allure.NewParameter("s3", testCase.s3),
			)

			actual := Concat(testCase.s1, testCase.s2, testCase.s3)

			t.Assert().Equal(testCase.expected, actual, "Concat two Int slices")
		})
	}

	testCasesStrs := []struct {
		s1, s2, s3 []string
		expected   []string
		testName   string
	}{
		{
			s1:       []string{"1", "2", "3"},
			s2:       []string{"4", "5", "6"},
			s3:       []string{"7", "8", "9"},
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			testName: "OK. Strs",
		},
	}
	for _, testCase := range testCasesStrs {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rslices")
			t.Story("Concat")
			t.Description("Check func `Concat`. Slices of Strings")
			t.WithParameters(
				allure.NewParameter("s1", testCase.s1),
				allure.NewParameter("s2", testCase.s2),
				allure.NewParameter("s3", testCase.s3),
			)

			actual := Concat(testCase.s1, testCase.s2, testCase.s3)

			t.Assert().Equal(testCase.expected, actual, "Concat two Strings slices")
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
