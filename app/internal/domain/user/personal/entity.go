package personal

import (
	"folyod/internal/domain/user/account/values/uuid"
	"folyod/internal/domain/user/personal/values/name"
)

type Personal struct {
	accountUuid uuid.AccountUuid
	name        name.Name
}
