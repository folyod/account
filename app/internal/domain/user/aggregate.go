package user

import (
	"folyod/internal/domain/user/account"
	"folyod/internal/domain/user/auth"
	"folyod/internal/domain/user/personal"
)

type User struct {
	account.Account
	personal.Personal
	auth.Auth
}
