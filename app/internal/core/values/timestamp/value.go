package timestamp

import "time"

type Timestamp struct {
	value time.Time
}

func (t Timestamp) Value() time.Time {
	return t.value
}

func (t Timestamp) Equal(timestamp Timestamp) bool {
	return t.Value() == timestamp.Value()
}
