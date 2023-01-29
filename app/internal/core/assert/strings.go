package assert

type Str struct {
	value string
}

func String(value string) *Str {
	return &Str{
		value: value,
	}
}

func (s *Str) LengthLess(limit int) (bool, int) {
	actual := len(s.value)
	return actual < limit, actual
}

func (s *Str) LengthMore(limit int) (bool, int) {
	actual := len(s.value)
	return actual > limit, actual
}

func (s *Str) Empty() bool {
	return s.value == ""
}

func (s *Str) NotEmpty() bool {
	return s.value != ""
}

func (s *Str) IsEmail() bool {
	return s.value != ""
}
