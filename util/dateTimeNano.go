package util

import (
	"time"
)

// Provides one-way conversion of date
type DateTimeNano time.Time

func NewDateTimeNano(time time.Time) (*DateTimeNano, error) {
	dateTime := DateTimeNano(time)
	return &dateTime, nil
}

func (t *DateTimeNano) UnmarshalText(b []byte) error {
	result, err := time.Parse(time.RFC3339Nano, string(b))
	if err != nil {
		var t2 time.Time
		err = t2.UnmarshalText(b)
		if err != nil {
			return err
		}
		*t = DateTimeNano(t2)
		return nil
	}

	// Save as data
	*t = DateTimeNano(result)
	return nil
}

func (t DateTimeNano) Time() time.Time {
	return (time.Time)(t)
}
