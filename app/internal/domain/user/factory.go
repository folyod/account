package user

import (
	"folyod/internal/domain/user/account"
	"folyod/internal/domain/user/personal"
)

type newUserDto interface {
	Nickname() string
	Name() string
	Email() string
	Password() string
}

func New(dto newUserDto) User {
	user := User{}
	user.Account = account.NewAccount(dto)
	user.Personal = personal.New(user.Account, dto)

	return user
}
