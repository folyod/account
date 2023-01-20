package timestamp

import "time"

func MakeFromTime(timestamp time.Time) *Timestamp {
	return &Timestamp{
		value: timestamp,
	}
}

func MakeFromString(timestamp string) (*Timestamp, error) {
	formattedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return nil, err
	}
	value := &Timestamp{
		value: formattedTime,
	}
	return value, nil
}

func Now() *Timestamp {
	return &Timestamp{
		value: time.Now(),
	}
}
