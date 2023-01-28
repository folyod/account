package assert

import "testing"

func TestStringLengthLess(t *testing.T) {
	dataProvider := []struct {
		string      string
		limit       int
		expectation bool
	}{
		{"somestring", 100, true},
		{"somestring", 11, true},
		{"somestring", 10, false},
		{"somestring", 5, false},
		{"somestring", 0, false},
		{"somestring", -1, false},
	}

	for _, testcase := range dataProvider {
		assertion := String(testcase.string)
		if testcase.expectation != assertion.LengthLess(testcase.limit) {
			t.Fatalf(
				"Actual string length (%d) more than expectation (%d) characters",
				len(testcase.string),
				testcase.limit)
		}
	}
}
