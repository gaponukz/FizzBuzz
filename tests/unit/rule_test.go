package unit

import (
	"fizzbuzz/internal/usecase"
	"fizzbuzz/internal/value_objects"
	"testing"
)

func TestRules(t *testing.T) {
	ruleInstance := usecase.NewRule()

	cases := []struct {
		InputNumber int
		ExpectedTag value_objects.Tag
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{7, "7"},
		{9, "Fizz"},
		{10, "Buzz"},
		{30, "FizzBuzz"},
		{14, "14"},
		{18, "Fizz"},
		{20, "Buzz"},
		{45, "FizzBuzz"},
		{28, "28"},
		{33, "Fizz"},
		{40, "Buzz"},
		{60, "FizzBuzz"},
		{77, "77"},
		{90, "FizzBuzz"},
		{100, "Buzz"},
		{1000, "Buzz"},
	}

	for _, testCase := range cases {
		t.Run(string(testCase.ExpectedTag), func(t *testing.T) {
			actualTag, err := ruleInstance.Convert(testCase.InputNumber)
			if err != nil {
				t.Error(err.Error())
			}

			if actualTag != testCase.ExpectedTag {
				t.Errorf("For input %d, expected %s, but got %s", testCase.InputNumber, testCase.ExpectedTag, actualTag)
			}
		})
	}
}
