package tester

import "testing"

func AssertNil(t *testing.T, v interface{}) {
	if v != nil {
		t.Fatalf("Expection nill, actual %T", v)
	}
}

func AssertNotNil(t *testing.T, v interface{}) {
	if v != nil {
		t.Fatalf("Value is nill")
	}
}

func AssertSame(t *testing.T, v1 interface{}, v2 interface{}) {
	if v1 != v2 {
		t.Fatalf("Values not the same. Expect (%s). Actual (%s)", v1, v2)
	}
}
