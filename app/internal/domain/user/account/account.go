package account

import (
	"core/values/timestamp"
	"user/account/uuid"
	"user/personal"
)

type Account struct {
	uuid      uuid.AccountUuid
	personal  personal.Personal
	createdAt timestamp.Timestamp
}
