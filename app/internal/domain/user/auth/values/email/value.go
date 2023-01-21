package email

type Email struct {
	value string
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equal(email Email) bool {
	return e.Value() == email.Value()
}
