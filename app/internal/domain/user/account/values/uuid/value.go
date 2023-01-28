package uuid

type AccountUuid struct {
	value string
}

func (vo *AccountUuid) Value() string {
	return vo.value
}

func (vo *AccountUuid) Equal(uuid AccountUuid) bool {
	return vo.Value() == uuid.Value()
}
