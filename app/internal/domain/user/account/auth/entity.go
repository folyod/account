package auth

import (
	"errors"
	"folyod/internal/domain/user/account/auth/values/email"
	"folyod/internal/domain/user/account/auth/values/password"
	"folyod/internal/domain/user/account/auth/values/phone"
	"folyod/internal/domain/user/account/values/uuid"
)

type Auth struct {
	accountUuid uuid.AccountUuid
	email       email.Email
	password    password.Password
	phone       phone.Phone
}

func New(
	accountUuid uuid.AccountUuid,
	email email.Email,
	password password.Password,
	phone phone.Phone) (*Auth, error) {
	isEmpty := email.Value() != "" ||
		password.Value() != "" ||
		phone.Value() != ""
	if isEmpty {
		return nil, errors.New("ff")
	}
	return Auth{
		accountUuid: accountUuid,
		email:       email,
		password:    password,
		phone:       phone,
	}
}

func (a *Auth) Email() email.Email {
	return a.email
}

func (a *Auth) Phone() phone.Phone {
	return a.phone
}

func (a *Auth) Password() password.Password {
	return a.password
}
