package personal

import (
	"account/internal/domain/personal/values/name"
	"account/internal/domain/user/values/uuid"
)

type Personal struct {
	accountUuid uuid.AccountUuid
	name        name.Name
}
