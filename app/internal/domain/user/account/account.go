package account

import (
	"folyod/internal/core/values/timestamp"
	"folyod/internal/domain/user/account/uuid"
	"folyod/internal/domain/user/personal"
)

type Account struct {
	uuid      uuid.AccountUuid
	personal  personal.Personal
	createdAt timestamp.Timestamp
}
