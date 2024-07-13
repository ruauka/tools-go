package rtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestDays(t *testing.T) {
	testCases := []struct {
		d1       time.Time
		d2       time.Time
		testName string
		expected int
	}{
		{
			d1:       time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2022, 9, 9, 0, 0, 0, 0, time.UTC),
			expected: 365.0,
			testName: "OK. Not leap",
		},
		{
			d1:       time.Date(2024, 9, 9, 0, 0, 0, 0, time.UTC),
			d2:       time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC),
			expected: 366.0,
			testName: "OK. Leap",
		},
	}

	for _, testCase := range testCases {
		runner.Run(t, testCase.testName, func(t provider.T) {
			// allure id
			t.AllureID(fmt.Sprintf("%s_%s", t.Name(), testCase.testName))

			// allure report info
			t.Epic("rtime")
			t.Story("Days")
			t.Description("Check func `Days`")
			t.WithParameters(
				allure.NewParameter("Date1", testCase.d1),
				allure.NewParameter("Date2", testCase.d2),
			)

			actual := Days(testCase.d2, testCase.d1)

			t.Assert().Equal(testCase.expected, actual, "Checking the calculation of the number of days between two dates")
		})
	}
}

func ExampleDays() {
	var (
		d1 = time.Date(2022, 9, 9, 0, 0, 0, 0, time.UTC)
		d2 = time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC)
	)

	res := Days(d1, d2)
	fmt.Println(res)
	// Output: 365
}
