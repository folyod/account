package uuid

import (
	uuid "github.com/google/uuid"
)

func Make(uuid string) (*AccountUuid, error) {
	value := &AccountUuid{
		value: uuid,
	}

	return value, nil
}

func New() AccountUuid {
	return AccountUuid{
		value: uuid.NewString(),
	}
}
