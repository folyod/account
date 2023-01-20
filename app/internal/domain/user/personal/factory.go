package personal

import (
	"user/internal/domain/account/uuid"
	"user/internal/domain/personal/name"
)

type createPersonalDto interface {
	Name() string
}

func Make(dto createPersonalDto) Personal {
	return Personal{
		accountUuid: uuid.AccountUuid{},
		name:        name.Name{},
	}
}
