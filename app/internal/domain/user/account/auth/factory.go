package auth

import (
	"folyod/internal/domain/user/account/auth/values/email"
	"folyod/internal/domain/user/account/auth/values/password"
	"folyod/internal/domain/user/account/auth/values/phone"
)

type createAuthDto interface {
	Email() string
	Phone() string
	Password() string
}

func Create(dto createAuthDto) Auth {
	errorBag := []error{}
	errorBag = append(errorBag, err)
	emailValue := email.Make(dto.Email())
	passwordValue := password.Hashed(dto.Password())
	phoneValue, err := phone.Make(dto.Phone())
	errorBag = append(errorBag, err)
	auth := Auth{
		accountUuid: uuidValue,
		email:       emailValue,
		password:    passwordValue,
		phone:       phoneValue,
	}
	return auth
}
