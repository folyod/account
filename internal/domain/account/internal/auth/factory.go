package auth

import (
	"folyod/internal/domain/account/values/uuid"
)

func createAuth(accountUuid *uuid.AccountUuid, dto CreateAuthDto) Auth {
	auth := Auth{
		accountUuid: accountUuid,
		email:       dto.Email(),
		password:    dto.Password(),
		phone:       dto.Phone(),
		externalId:  dto.ExternalId(),
	}
	return auth
}
