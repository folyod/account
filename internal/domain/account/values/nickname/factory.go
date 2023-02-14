package nickname

import (
	"github.com/google/uuid"
)

func New(nickname string) *Nickname {
	return &Nickname{
		value: nickname,
	}
}

func Generate() *Nickname {
	return New(uuid.NewString())
}
