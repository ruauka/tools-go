package rtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestMonths(t *testing.T) {
	testCases := []struct {
		d1       time.Time
		d2       time.Time
		testName string
		expected int
	}{
		{
			d1:       time.Date(2022, 9, 9, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC),
			expected: 12,
			testName: "Check 1. d2 > d1. Days in the middle of the month",
		},
		{
			d1:       time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2022, 9, 9, 0, 0, 0, 0, time.UTC),
			expected: -12,
			testName: "Check 2. d1 > d2. Days in the middle of the month",
		},
		{
			d1:       time.Date(2023, 3, 31, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2022, 4, 29, 0, 0, 0, 0, time.UTC),
			expected: -11,
			testName: "Check 3. d1 > d2. d1 has the last day of the month",
		},
		{
			d1:       time.Date(2022, 4, 30, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2023, 5, 31, 0, 0, 0, 0, time.UTC),
			expected: 13,
			testName: "Check 4. d2 > d1. d1 and d2 have the last day of the month",
		},
		{
			d1:       time.Date(2022, 5, 31, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2023, 4, 29, 0, 0, 0, 0, time.UTC),
			expected: 10,
			testName: "Check 5. d2 > d1. 1 has the last day of the month and the day is longer than the day in d2",
		},
		{
			d1:       time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2023, 1, 30, 0, 0, 0, 0, time.UTC),
			expected: 10,
			testName: "Check 6. d2 > d1. d1 has the last day of the month and a day less than a day in d2",
		},
	}

	for _, testCase := range testCases {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rtime")
			t.Story("Months")
			t.Description("Check func `Months`")
			t.WithParameters(
				allure.NewParameter("Date1", testCase.d1),
				allure.NewParameter("Date2", testCase.d2),
			)

			actual := Months(testCase.d1, testCase.d2)

			t.Assert().Equal(testCase.expected, actual, "Checking the calculation of the difference between two dates in months")
		})
	}
}
