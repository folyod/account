package phone

type Phone struct {
	value string
}

func (p *Phone) Value() string {
	return p.value
}

func (p *Phone) Equal(phone Phone) bool {
	return p.Value() == phone.Value()
}
