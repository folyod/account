package account

import (
	"folyod/internal/domain/account/values/nickname"
)

type Repository interface {
	Add(account *Account) error
	HasNickname(nickname *nickname.Nickname) bool
}
