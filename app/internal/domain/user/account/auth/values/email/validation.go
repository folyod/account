package email

import "net/mail"

func (e *Email) Validate() error {
	_, err := mail.ParseAddress(e.value)
	if err != nil {
		return err
	}
	return nil
}
