package user

import (
	"folyod/internal/domain/user/account"
	"folyod/internal/domain/user/auth"
	"folyod/internal/domain/user/personal"
)

type User struct {
	Account  account.Account
	Personal personal.Personal
	Auth     auth.Auth
}
