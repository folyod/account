package personal

import (
	"account/internal/domain/personal/values/name"
)

type createPersonalDto interface {
	Name() string
}

func New(account account.Account, dto createPersonalDto) Personal {
	return Personal{
		accountUuid: account.Uuid(),
		name:        name.Make(dto.Name()),
	}
}
