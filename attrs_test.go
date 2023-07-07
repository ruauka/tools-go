package attrs_go

import (
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
			expectedErr: ErrNotPointer,
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
			err := SetAttr(testCase.obj, testCase.fieldName, testCase.newValue)
			require.ErrorIs(t, testCase.expectedErr, err)
		})
	}
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
