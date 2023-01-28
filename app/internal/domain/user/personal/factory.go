package personal

import (
	"folyod/internal/domain/user/personal/values/name"
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
