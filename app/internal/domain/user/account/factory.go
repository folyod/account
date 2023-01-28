package account

import (
	"folyod/internal/core/values/timestamp"
	"folyod/internal/domain/user/account/values/nickname"
	"folyod/internal/domain/user/account/values/uuid"
)

type createAccountDto interface {
	Nickname() string
}

func New(dto createAccountDto) Account {
	return Account{
		uuid:      uuid.New(),
		nickname:  nickname.Make(dto.Nickname()),
		createdAt: timestamp.Now(),
	}
}
