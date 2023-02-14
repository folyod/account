package account

import (
	"folyod/internal/domain/account/services"
)

func CreateAccount(repository Repository) (*Account, error) {
	acc := newAccount()
	err := services.SaveCreatedAccount(acc, repository)
	return acc, err
}

func Login() {

}

func Logout() {

}

func ChangeNickname() {

}

func ChangePassword() {

}

func ChangeEmail() {

}

func ChangeExternalId() {

}
