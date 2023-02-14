package uuid

import (
	"github.com/google/uuid"
)

func New(uuid string) (*AccountUuid, error) {
	value := &AccountUuid{
		value: uuid,
	}
	return value, nil
}

func Generate() *AccountUuid {
	return &AccountUuid{
		value: uuid.NewString(),
	}
}
