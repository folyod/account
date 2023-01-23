package auth

import (
	"folyod/internal/domain/user/auth/values/email"
	"folyod/internal/domain/user/auth/values/password"
	"folyod/internal/domain/user/values/uuid"
)

type Auth struct {
	accountUuid uuid.AccountUuid
	email       email.Email
	password    password.Password
}

func (a *Auth) Email() email.Email {
	return a.email
}
