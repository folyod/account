package assert

import "testing"

func TestStringLengthLess(t *testing.T) {
	dataProvider := []struct {
		string            string
		limit             int
		expectation       bool
		expectationLength int
	}{
		{"somestring", 100, true, 10},
		{"somestring", 11, true, 10},
		{"somestring", 10, false, 10},
		{"somestring", 5, false, 10},
		{"somestring", 0, false, 10},
		{"somestring", -1, false, 10},
	}

	for _, testcase := range dataProvider {
		assertion := String(testcase.string)
		result, actualLength := assertion.LengthLess(testcase.limit)
		if testcase.expectation != result {
			t.Fatalf(
				"Actual string length (%d) more than expectation (%d) characters",
				len(testcase.string),
				testcase.limit)
		}
		if actualLength != testcase.expectationLength {
			t.Fatalf(
				"Expectaion string length (%d) characters - actual (%d)",
				testcase.expectationLength,
				len(testcase.string))
		}
	}
}
