package login

import (
	"errors"
)

type Login struct {
	value string
}

func (vo *Login) New(login string) error {
	vo.value = login

	length := len([]rune(login))

	if length < 2 {
		return errors.New("login should be more than 1 symbol")
	}

	if length > 32 {
		return errors.New("text")
	}

	return nil
}

func (vo *Login) Value() string {
	return vo.value
}

func (vo *Login) Equal(login Login) bool {
	return vo.Value() == login.Value()
}
