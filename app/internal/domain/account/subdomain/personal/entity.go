package personal

import (
	"folyod/internal/domain/user/personal/values/name"
	"folyod/internal/domain/user/values/uuid"
)

type Personal struct {
	accountUuid uuid.AccountUuid
	name        name.Name
}
