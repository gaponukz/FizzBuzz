package unit

import (
	"fizzbuzz/src/application/rules"
	"fizzbuzz/src/application/strategies"
	"testing"
)

func TestRules(t *testing.T) {
	rule := rules.NewRule("Test", strategies.NewDivStrategy(2))

	if rule.Tag() != "Test" {
		t.Errorf("Expected Test, got %s", rule.Tag())
	}

	cases := []struct {
		number int
		result bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, true},
	}

	for _, testCase := range cases {
		success := rule.IsSuccess(testCase.number)
		if success != testCase.result {
			t.Errorf("Expected %t, got %t", testCase.result, success)
		}
	}
}
