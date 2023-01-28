package account

import (
	"errors"
)

type userRepository interface {
	HasEmail(email string) bool
	HasNickname(nickname string) bool
}

func CanCreateUser(usr User, repository userRepository) error {
	nickname := usr.Account.Nickname()
	email := usr.Auth.Email()
	if !repository.HasNickname(nickname.Value()) {
		return errors.New("ff")
	}
	if !repository.HasEmail(email.Value()) {
		return errors.New("dd")
	}
	return nil
}
