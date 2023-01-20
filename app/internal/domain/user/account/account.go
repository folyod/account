package account

import (
	"user/account/uuid"
	"user/personal"
	"values/timestamp"
)

type Account struct {
	uuid      uuid.AccountUuid
	personal  personal.Personal
	createdAt timestamp.Timestamp
}
