package account

import (
	"user/internal/domain/account/uuid"
	account "user/public"
)

func CreateAccount(command account.CreateAccountCommand) Account {
	return Account{
		uuid:     uuid.New(),
		personal: personal.,
		createdAt: nil,
	}
}
