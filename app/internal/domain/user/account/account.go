package account

import (
	"github.com/folyod/folyod/tree/master/app/internal/core"
	"user/account/uuid"
	"user/personal"
)

type Account struct {
	uuid      uuid.AccountUuid
	personal  personal.Personal
	createdAt timestamp.Timestamp
}
