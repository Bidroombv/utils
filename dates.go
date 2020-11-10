package utils

import (
	"encoding/json"
	"time"
)

const DateISOFormat = "2006-01-02"

type DateISO struct {
	time.Time
}

func (d *DateISO) UnmarshalJSON(p []byte) error {
	var val string
	err := json.Unmarshal(p, &val)
	if err != nil {
		return err
	}

	decoded, err := time.Parse(DateISOFormat, val)
	if err != nil {
		return err
	}

	d.Time = decoded

	return nil
}

func (d *DateISO) String() string {
	return d.Format(DateISOFormat)
}
