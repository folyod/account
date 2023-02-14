package email

import (
	"net/mail"
)

func (e *Email) Validate() error {
	_, err := mail.ParseAddress(e.value)
	if err != nil {
		return err
	}
	return nil
}

func (e *Email) filterNotEmpty() error {
	if !assertions.StringIsNotEmpty(e.value) {

	}
}

func (e *Email) filterMaxLength() error {

}

func (e *Email) filterEmail() error {

}
