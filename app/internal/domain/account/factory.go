package account

import (
	"folyod/internal/core/values/timestamp"
	"folyod/internal/domain/user/values/nickname"
	"folyod/internal/domain/user/values/uuid"
)

type createAccountDto interface {
	Nickname() string
}

func NewAccount(dto createAccountDto) Account {
	return Account{
		uuid:      uuid.New(),
		nickname:  nickname.Make(dto.Nickname()),
		createdAt: timestamp.Now(),
	}
}
