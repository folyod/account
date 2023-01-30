package user

import (
	"errors"
	"folyod/internal/domain/user/account"
)

type registerRequest interface {
	Name() string
	Nickname() string
	Email() string
	Password() string
}

type registerViewModel interface {
}

type userRepository interface {
	Add(usr account.User) bool
	HasEmail(email string) bool
	HasNickname(nickname string) bool
}

func Register(request registerRequest, repository userRepository) (string, error) {
	usr := account.New(request)
	err := account.CanCreateUser(usr, repository)
	if err != nil {
		return "", err
	}
	isSaved := repository.Add(usr)
	if !isSaved {
		return "", errors.New("")
	}
	uuid := usr.Account.Uuid()
	return uuid.Value(), nil
}
