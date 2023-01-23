package password

type Password struct {
	value string
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) Equal(password Password) bool {
	return p.Value() == password.Value()
}

func (p *Password) Check(password Password) bool {
	return p.Value() == password.Value()
}
