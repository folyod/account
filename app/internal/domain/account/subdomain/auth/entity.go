package auth

import (
	"folyod/internal/domain/account/subdomain/auth/values/email"
	"folyod/internal/domain/account/subdomain/auth/values/password"
	"folyod/internal/domain/account/subdomain/auth/values/phone"
	"folyod/internal/domain/account/values/uuid"
)

type Auth struct {
	accountUuid uuid.AccountUuid
	email       email.Email
	password    password.Password
	phone       phone.Phone
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
