package email

import (
	"account/internal/core/tester"
	"testing"
)

func TestCreateCorrectEmail(t *testing.T) {
	_, err := Make("correctemail@email.com")
	tester.AssertNil(t, err)
}

func TestIncorrectEmail(t *testing.T) {
	_, err := Make("incorrectemail")
	tester.AssertNotNil(t, err)
	tester.AssertSame(t, "Value is not correct email", err.Error())
}
