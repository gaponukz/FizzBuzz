package unit

import (
	"fizzbuzz/internal/application/rules"
	"fizzbuzz/internal/application/strategies"
	"fizzbuzz/internal/domain/value_objects"
	"strconv"
	"testing"
)

func TestRules(t *testing.T) {
	fizzCondition := strategies.NewDivStrategy(3)
	buzzCondition := strategies.NewDivStrategy(5)
	ruleInstance := rules.NewRulesCollection(
		rules.NewRule("FizzBuzz", strategies.NewAndStrategy(fizzCondition, buzzCondition)),
		rules.NewRule("Fizz", fizzCondition),
		rules.NewRule("Buzz", buzzCondition),
	)

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
			actualTag := ruleInstance.Find(testCase.InputNumber, value_objects.Tag(strconv.Itoa(testCase.InputNumber)))
			if actualTag != testCase.ExpectedTag {
				t.Errorf("For input %d, expected %s, but got %s", testCase.InputNumber, testCase.ExpectedTag, actualTag)
			}
		})
	}
}
