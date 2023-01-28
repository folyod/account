package email

import "testing"

func testCreateEmail(t *testing.T) {
	email, err := Make("test@example.com")

	if email.Value() != "test@example.com" {
		t.Errorf("")
	}
}
