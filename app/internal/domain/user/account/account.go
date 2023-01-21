package account

import (
	"folyod/internal/core/values/timestamp"
	"folyod/internal/domain/user/account/values/nickname"
	"folyod/internal/domain/user/account/values/uuid"
)

type Account struct {
	uuid      uuid.AccountUuid
	nickname  nickname.Nickname
	createdAt timestamp.Timestamp
}

func (a *Account) Uuid() uuid.AccountUuid {
	return a.uuid
}
