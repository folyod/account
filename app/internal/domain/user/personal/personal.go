package personal

import (
	"user/internal/domain/account/uuid"
	"user/internal/domain/personal/name"
)

type Personal struct {
	accountUuid uuid.AccountUuid
	name        name.Name
}
