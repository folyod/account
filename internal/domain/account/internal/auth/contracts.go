package auth

import (
	"folyod/internal/domain/account/internal/auth/values/email"
	"folyod/internal/domain/account/internal/auth/values/external_id"
	"folyod/internal/domain/account/internal/auth/values/password"
	"folyod/internal/domain/account/internal/auth/values/phone"
)

type CreateAuthDto interface {
	Email() *email.Email
	Password() *password.Password
	Phone() *phone.Phone
	ExternalId() *external_id.ExternalId
}

type Repository interface {
	HasEmail(email.Email) bool
	HasPhone(phone.Phone) bool
	HasExternalId(id external_id.ExternalId) bool
}
