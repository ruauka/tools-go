package attrs_go

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAttr(t *testing.T) {
	name := "test"

	TestCases := []struct {
		obj         interface{}
		fieldName   string
		expected    interface{}
		expectedErr error
		testName    string
	}{
		{
			obj:         struct{ Username string }{Username: name},
			fieldName:   "Username",
			expected:    name,
			expectedErr: nil,
			testName:    "OK. Value field",
		},
		{
			obj:         struct{ Username *string }{Username: &name},
			fieldName:   "Username",
			expected:    name,
			expectedErr: nil,
			testName:    "OK. Ptr field",
		},
		{
			obj:         &struct{ Username string }{Username: name},
			fieldName:   "Username",
			expected:    nil,
			expectedErr: ErrPointerStruct,
			testName:    "OK. Ptr struct",
		},
		{
			obj:         "not struct arg",
			fieldName:   "Username",
			expected:    nil,
			expectedErr: ErrNotStruct,
			testName:    "ERR. Arg not struct",
		},
		{
			obj:         struct{ Username string }{Username: name},
			fieldName:   "not in struct field",
			expected:    nil,
			expectedErr: ErrFieldNotInStruct,
			testName:    "ERR. Field not struct",
		},
		{
			obj:         struct{ username string }{username: name},
			fieldName:   "username",
			expected:    nil,
			expectedErr: ErrUnexportedField,
			testName:    "ERR. Unexported field",
		},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.testName, func(t *testing.T) {
			actual, err := GetAttr(testCase.obj, testCase.fieldName)
			if err != nil {
				require.ErrorIs(t, err, testCase.expectedErr)
			}

			require.Equal(t, testCase.expected, actual)
		})
	}
}

func ExampleGetAttr() {
	type User struct {
		Username string
	}

	user := User{Username: "username"}

	value, err := GetAttr(user, "Username")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
	// Output: username
}

func TestSetAttr(t *testing.T) {
	curValue := "test"
	newValue := "new_test"
	fieldUsername := "Username"

	TestCases := []struct {
		obj         interface{}
		fieldName   string
		newValue    interface{}
		expectedErr error
		testName    string
	}{
		{
			obj:         &struct{ Username string }{Username: curValue},
			fieldName:   fieldUsername,
			newValue:    newValue,
			expectedErr: nil,
			testName:    "OK. Value field",
		},
		{
			obj:         &struct{ Username *string }{Username: &curValue},
			fieldName:   fieldUsername,
			newValue:    newValue,
			expectedErr: nil,
			testName:    "OK. Value field",
		},
		{
			obj:         struct{ Username string }{Username: curValue},
			fieldName:   fieldUsername,
			newValue:    newValue,
			expectedErr: ErrNotPointerStruct,
			testName:    "ERR. Struct passed not by pointer",
		},
		{
			obj:         &newValue,
			fieldName:   fieldUsername,
			newValue:    newValue,
			expectedErr: ErrNotStruct,
			testName:    "ERR. Not struct arg",
		},
		{
			obj:         &struct{ Username string }{Username: curValue},
			fieldName:   "Field not in struct",
			newValue:    newValue,
			expectedErr: ErrFieldNotInStruct,
			testName:    "ERR. Field not in struct",
		},
		{
			obj:         &struct{ Username int }{Username: 0},
			fieldName:   fieldUsername,
			newValue:    newValue,
			expectedErr: ErrWrongFieldValueType,
			testName:    "ERR. Wrong field value type",
		},
		{
			obj:         &struct{ username string }{username: curValue},
			fieldName:   strings.ToLower(fieldUsername),
			newValue:    newValue,
			expectedErr: ErrUnexportedField,
			testName:    "ERR. Unexported field",
		},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.testName, func(t *testing.T) {
			err := SetAttr(testCase.obj, testCase.newValue, testCase.fieldName)
			require.ErrorIs(t, testCase.expectedErr, err)
		})
	}
}

func ExampleSetAttr() {
	type User struct {
		Username string
	}

	u := &User{Username: "username"}

	if err := SetAttr(u, "new_username", "Username"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Username)
	// Output: new_username
}

func TestSetStructAttrs(t *testing.T) {
	type UserValue struct {
		Username string
		Age      int
		Married  bool
		Friends  []string
	}

	type UserPtr struct {
		Username *string
		Age      *int
		Married  *bool
		Friends  []string
	}

	curUsername := "username"
	curAge := 20
	curMarried := true
	curFriends := []string{"fried1", "friend2"}

	newUsername := "new_username"
	newAge := 30
	newMarried := false
	newFriends := []string{"new_fried1", "new_friend2"}

	TestCases := []struct {
		curObj      interface{}
		newObj      interface{}
		expected    interface{}
		expectedErr error
		testName    string
	}{
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: UserValue{
				Username: newUsername,
				Age:      newAge,
				Married:  newMarried,
				Friends:  newFriends,
			},
			expected: &UserValue{
				Username: newUsername,
				Age:      newAge,
				Married:  newMarried,
				Friends:  newFriends,
			},
			expectedErr: nil,
			testName:    "OK. Value fields. Changes all fields of curObj",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: UserValue{
				Username: newUsername,
				Age:      newAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			expected: &UserValue{
				Username: newUsername,
				Age:      newAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			expectedErr: nil,
			testName:    "OK. Value fields. Changes 2 fields of curObj",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: UserPtr{
				Username: &newUsername,
				Age:      &newAge,
				Married:  &newMarried,
				Friends:  newFriends,
			},
			expected: &UserValue{
				Username: newUsername,
				Age:      newAge,
				Married:  newMarried,
				Friends:  newFriends,
			},
			expectedErr: nil,
			testName:    "OK. Ptr fields. Changes all fields of curObj",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: UserPtr{
				Username: &newUsername,
				Age:      &newAge,
				Married:  &curMarried,
				Friends:  curFriends,
			},
			expected: &UserValue{
				Username: newUsername,
				Age:      newAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			expectedErr: nil,
			testName:    "OK. Ptr fields. Changes 2 fields of curObj",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: UserPtr{
				Username: &newUsername,
				Age:      nil,
				Married:  nil,
				Friends:  nil,
			},
			expected: &UserValue{
				Username: newUsername,
				Age:      curAge,
				Married:  curMarried,
			},
			expectedErr: nil,
			testName:    "OK. Ptr fields. Some fields are nil",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: struct {
				username string
			}{
				username: "Field not in struct",
			},
			expected: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			expectedErr: ErrUnexportedField,
			testName:    "ERR. Value fields. Err in GetAttr, field not exported",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: struct {
				Username int
			}{
				Username: 0,
			},
			expected: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			expectedErr: ErrWrongFieldValueType,
			testName:    "ERR. Value fields. Err in SetAttr, Wrong field value type",
		},
		{
			curObj: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			newObj: 0,
			expected: &UserValue{
				Username: curUsername,
				Age:      curAge,
				Married:  curMarried,
				Friends:  curFriends,
			},
			expectedErr: ErrNotStruct,
			testName:    "ERR. Value fields. Arg not struct",
		},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.testName, func(t *testing.T) {
			err := SetStructAttrs(testCase.curObj, testCase.newObj)
			if err != nil {
				require.ErrorIs(t, err, testCase.expectedErr)
			}

			require.Equal(t, testCase.expected, testCase.curObj)
		})
	}
}

func ExampleSetStructAttrs() {
	type User struct {
		Username string // will change by pte
		Age      int    // will change by value
		Married  bool   // will be the same
	}

	type NewUser struct {
		Username *string `json:"username"`
		Age      int     `json:"age"`
		Married  *bool   `json:"married"` // nil
	}

	user := &User{
		Username: "username",
		Age:      30,
		Married:  true,
	}

	newUserName := "new_username"
	newUser := NewUser{
		Username: &newUserName,
		Age:      35,
		Married:  nil,
	}

	fmt.Printf("%s, %d, %v\n", user.Username, user.Age, user.Married)

	if err := SetStructAttrs(user, newUser); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s, %d, %v\n", user.Username, user.Age, user.Married)
	// Output:
	// username, 30, true
	// new_username, 35, true
}

func TestRound(t *testing.T) {
	t.Run("OK float32", func(t *testing.T) {
		result := Round(float32(0.123456789), 4)
		require.Equal(t, float32(0.1235), result)
	})
	t.Run("OK float64", func(t *testing.T) {
		result := Round(0.123456789, 4)
		require.Equal(t, 0.1235, result)
	})
}

func ExampleRound() {
	var (
		val32 float32 = 0.12345
		val64 float64 = 0.12345
	)

	res32 := Round(val32, 3)
	fmt.Println(res32)
	fmt.Println(reflect.TypeOf(res32))

	res64 := Round(val64, 3)
	fmt.Println(res64)
	fmt.Println(reflect.TypeOf(res64))

	// Output:  0.123
	// float32
	// 0.123
	// float64
}

func TestRoundFloatStruct(t *testing.T) {
	notStruct := "arg not struct"

	type Foo struct {
		Field1 float32
		Field2 float64
		Field3 []float32
		Field4 []float64
		Field5 [3]float32
		Field6 [3]float64
		Field7 int
		Field8 string
	}

	foo := Foo{
		Field1: 1.1111,
		Field2: 2.2222,
		Field3: []float32{1.1111, 2.2222, 3.3333},
		Field4: []float64{4.4444, 5.5555, 7.7777},
		Field5: [3]float32{1.1111, 2.2222, 3.3333},
		Field6: [3]float64{4.4444, 5.5555, 7.7777},
		Field7: 7,
		Field8: "field8",
	}

	expected := Foo{
		Field1: 1.111,
		Field2: 2.222,
		Field3: []float32{1.111, 2.222, 3.333},
		Field4: []float64{4.444, 5.556, 7.778},
		Field5: [3]float32{1.111, 2.222, 3.333},
		Field6: [3]float64{4.444, 5.556, 7.778},
		Field7: 7,
		Field8: "field8",
	}

	TestCases := []struct {
		obj         interface{}
		precision   int
		expected    interface{}
		expectedErr error
		testName    string
	}{
		{
			obj:         &foo,
			precision:   3,
			expected:    &expected,
			expectedErr: nil,
			testName:    "OK",
		},
		{
			obj:         foo,
			precision:   3,
			expected:    foo,
			expectedErr: ErrNotPointerStruct,
			testName:    "ERR. Struct passed not by pointer",
		},
		{
			obj:         &notStruct,
			precision:   3,
			expected:    &notStruct,
			expectedErr: ErrNotStruct,
			testName:    "ERR. Arg not struct",
		},
		{
			obj:         &struct{ field1 float64 }{field1: 0},
			precision:   3,
			expected:    &struct{ field1 float64 }{field1: 0},
			expectedErr: ErrUnexportedField,
			testName:    "ERR. Unexported field",
		},
		{
			obj:         &struct{ Field1 []float64 }{Field1: nil},
			precision:   3,
			expected:    &struct{ Field1 []float64 }{Field1: nil},
			expectedErr: nil,
			testName:    "OK. Slice len == 0",
		},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.testName, func(t *testing.T) {
			err := RoundFloatStruct(testCase.obj, testCase.precision)
			if err != nil {
				require.ErrorIs(t, testCase.expectedErr, err)
			}

			require.Equal(t, testCase.expected, testCase.obj)
		})
	}
}

func ExampleRoundFloatStruct() {
	type Foo struct {
		Field1 float32
		Field2 float64
		Field3 []float32
		Field4 []float64
		Field5 [3]float32
		Field6 [3]float64
		Field7 int    // will be the same
		Field8 string // will be the same
	}

	foo := &Foo{
		Field1: 1.1111,
		Field2: 2.2222,
		Field3: []float32{1.1111, 2.2222, 3.3333},
		Field4: []float64{4.4444, 5.5555, 7.7777},
		Field5: [3]float32{1.1111, 2.2222, 3.3333},
		Field6: [3]float64{4.4444, 5.5555, 7.7777},
		Field7: 7,
		Field8: "field8",
	}

	fmt.Printf("%+v\n", *foo)

	if err := RoundFloatStruct(foo, 3); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", *foo)
	// Output: {Field1:1.1111 Field2:2.2222 Field3:[1.1111 2.2222 3.3333] Field4:[4.4444 5.5555 7.7777] Field5:[1.1111 2.2222 3.3333] Field6:[4.4444 5.5555 7.7777] Field7:7 Field8:field8}
	// {Field1:1.111 Field2:2.222 Field3:[1.111 2.222 3.333] Field4:[4.444 5.556 7.778] Field5:[1.111 2.222 3.333] Field6:[4.444 5.556 7.778] Field7:7 Field8:field8}
}

func TestIntersection(t *testing.T) {
	type args[T any] struct {
		left, right []T
	}

	type cases[T any] struct {
		args        args[T]
		expected    []T
		expectedErr error
		testName    string
	}

	testCasesInts := []cases[int]{
		{
			args:     args[int]{left: []int{1, 2, 3}, right: []int{1, 2, 4}},
			expected: []int{1, 2},
			testName: "Ints. OK",
		},
		{
			args:        args[int]{left: []int{1}, right: []int{1, 2}},
			expectedErr: ErrLenSlices,
			testName:    "Ints. Error length",
		},
	}
	for _, testCase := range testCasesInts {
		t.Run(testCase.testName, func(t *testing.T) {
			actual, err := Intersection(testCase.args.left, testCase.args.right)
			if err != nil {
				require.ErrorIs(t, testCase.expectedErr, err)
			}

			require.Equal(t, testCase.expected, actual)
		})
	}

	testCasesFloats := []cases[float64]{
		{
			args:     args[float64]{left: []float64{1, 2, 3}, right: []float64{1, 2, 4}},
			expected: []float64{1, 2},
			testName: "Floats. OK",
		},
		{
			args:        args[float64]{left: []float64{1}, right: []float64{1, 2}},
			expectedErr: ErrLenSlices,
			testName:    "Floats. Error length",
		},
	}
	for _, testCase := range testCasesFloats {
		t.Run(testCase.testName, func(t *testing.T) {
			actual, err := Intersection(testCase.args.left, testCase.args.right)
			if err != nil {
				require.ErrorIs(t, testCase.expectedErr, err)
			}

			require.Equal(t, testCase.expected, actual)
		})
	}

	testCasesStrs := []cases[string]{
		{
			args:     args[string]{left: []string{"aaa", "bbb", "ccc"}, right: []string{"aaa", "bbb", "ddd"}},
			expected: []string{"aaa", "bbb"},
			testName: "Strings. OK",
		},
		{
			args:        args[string]{left: []string{"aaa"}, right: []string{"aaa", "bbb"}},
			expectedErr: ErrLenSlices,
			testName:    "Strings. Error length",
		},
	}
	for _, testCase := range testCasesStrs {
		t.Run(testCase.testName, func(t *testing.T) {
			actual, err := Intersection(testCase.args.left, testCase.args.right)
			if err != nil {
				require.ErrorIs(t, testCase.expectedErr, err)
			}

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

	resInts, _ := Intersection(intsL, intsR)
	fmt.Println(resInts)

	resFloats, _ := Intersection(floatsL, floatsR)
	fmt.Println(resFloats)

	resStrs, _ := Intersection(strL, strR)
	fmt.Println(resStrs)

	// Output:
	// [1 2]
	// [1.1 2.2]
	// [aaa bbb]
}

func TestSlicesConcat(t *testing.T) {
	TestCasesInts := []struct {
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

	for _, testCase := range TestCasesInts {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := SlicesConcat(testCase.sl1, testCase.sl2, testCase.sl3)
			require.Equal(t, testCase.expected, actual)
		})
	}

	TestCases := []struct {
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

	for _, testCase := range TestCases {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := SlicesConcat(testCase.sl1, testCase.sl2, testCase.sl3)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func ExampleSlicesConcat() {
	var (
		ints1 = []int{1, 2, 3}
		ints2 = []int{4, 5, 6}
		ints3 = []int{7, 8, 9}

		strs1 = []string{"1", "2", "3"}
		strs2 = []string{"4", "5", "6"}
		strs3 = []string{"7", "8", "9"}
	)

	ints := SlicesConcat(ints1, ints2, ints3)
	fmt.Println(ints)

	strs := SlicesConcat(strs1, strs2, strs3)
	fmt.Println(strs)

	// Output:
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
}
