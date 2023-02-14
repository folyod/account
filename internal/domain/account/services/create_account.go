package services

import (
	"folyod/internal/domain/account"
	"folyod/internal/domain/account/values/nickname"
)

func SaveCreatedAccount(account *account.Account, repository account.Repository) error {
	if repository.HasNickname(account.Nickname()) {
		nick := generateNickname(repository)
		account.ChangeNickname(nick)
	}
	return repository.Add(account)
}

func generateNickname(repository account.Repository) *nickname.Nickname {
	var nick *nickname.Nickname
	for i := 0; i < 10; i++ {
		nick = nickname.Generate()
		if repository.HasNickname(nick) {
			continue
		}
		break
	}
	if nick == nil {
		panic("Неудалось сформировать уникальное имя аккаунта")
	}
	return nick
}
