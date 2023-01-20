package account

import (
	"time"
)

func CreateAccount(command CreateAccountCommand) Account {

	return Account{
		uuid:      "",
		login:     "",
		createdAt: time.Time{},
		personal:  Personal{},
		auth:      Auth{},
	}
}
