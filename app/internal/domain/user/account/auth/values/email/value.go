package email

type Email struct {
	value string
}

func (e *Email) Value() string {
	err := e.Validate()
	if err != nil {
		panic(err.Error())
	}
	return e.value
}

func (e *Email) Equal(email Email) bool {
	return e.Value() == email.Value()
}
