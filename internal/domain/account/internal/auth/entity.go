package auth

import (
	"folyod/internal/domain/account/internal/auth/values/email"
	"folyod/internal/domain/account/internal/auth/values/external_id"
	"folyod/internal/domain/account/internal/auth/values/password"
	"folyod/internal/domain/account/internal/auth/values/phone"
	"folyod/internal/domain/account/values/uuid"
)

type Auth struct {
	accountUuid *uuid.AccountUuid
	externalId  *external_id.ExternalId
	email       *email.Email
	phone       *phone.Phone
	password    *password.Password
}

func (a *Auth) Email() *email.Email {
	return a.email
}

func (a *Auth) Phone() *phone.Phone {
	return a.phone
}

func (a *Auth) Password() *password.Password {
	return a.password
}
