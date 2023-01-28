package account

import (
	"folyod/internal/core/values/timestamp"
	"folyod/internal/domain/user/account/auth"
	"folyod/internal/domain/user/account/auth/values/email"
	"folyod/internal/domain/user/account/auth/values/password"
	"folyod/internal/domain/user/account/values/nickname"
	"folyod/internal/domain/user/account/values/uuid"
)

type Account struct {
	uuid      uuid.AccountUuid
	nickname  nickname.Nickname
	createdAt timestamp.Timestamp
	auth      auth.Auth
}



func (a *Account) Uuid() uuid.AccountUuid {
	return a.uuid
}

func (a *Account) Nickname() nickname.Nickname {
	return a.nickname
}

func (a *Account) ChangeEmail(email email.Email) {
	a.auth = auth.New(email: email, password: a.auth.)
}
