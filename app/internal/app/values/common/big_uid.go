package common

type BigUid struct {
	value uint64
}

func (u *BigUid) Get() uint64 {
	return u.value
}

func (u *BigUid) Equal(uid BigUid) bool {
	return u.value == uid.value
}

func New(value uint64) (BigUid, error) {
	str := BigUid{
		value: value,
	}

	return str, nil
}
