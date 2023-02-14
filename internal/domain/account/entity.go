package account

import (
	"folyod/internal/domain/account/values/nickname"
	"folyod/internal/domain/account/values/uuid"
	"time"
)

type Account struct {
	uuid      *uuid.AccountUuid
	nickname  *nickname.Nickname
	createdAt time.Time
}

func newAccount() *Account {
	return &Account{
		uuid:      uuid.Generate(),
		nickname:  nickname.Generate(),
		createdAt: time.Now(),
	}
}

func (a *Account) Uuid() *uuid.AccountUuid {
	return a.uuid
}

func (a *Account) Nickname() *nickname.Nickname {
	return a.nickname
}

func (a *Account) ChangeNickname(nickname *nickname.Nickname) {
	a.nickname = nickname
}

func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}
