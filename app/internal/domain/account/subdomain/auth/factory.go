package auth

import (
	"folyod/internal/domain/account/subdomain/auth/values/email"
	"folyod/internal/domain/account/subdomain/auth/values/password"
	"folyod/internal/domain/account/subdomain/auth/values/phone"
	"folyod/internal/domain/account/values/uuid"
)

type createAuthDto interface {
	Email() string
	Phone() string
	Password() string
}

func Create(accountUuid string, dto createAuthDto) Auth {
	errorBag := []error{}
	uuidValue, err := uuid.Make(accountUuid)
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
