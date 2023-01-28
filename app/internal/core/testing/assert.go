package testing

import "testing"

type Assert struct {
	v interface{}
	t testing.T
}

func (a *Assert) Value(value interface{}, t *testing.T) {
	a.v = value
	a.t = t
}

func AssertSame(v1 interface{}, v2 interface{}) {
	test
}
