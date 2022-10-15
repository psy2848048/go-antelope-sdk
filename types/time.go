package types

import (
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	timeFormat := "2006-01-02T15:04:05.999999999"
	timeStr := strings.ReplaceAll(string(b), `"`, "")
	appendedTime := strings.Join([]string{timeStr, "000000"}, "")

	ret, err := time.Parse(timeFormat, appendedTime)
	if err != nil {
		return err
	}

	*t = Time(ret)
	return nil
}

func (t Time) String() string {
	tt := time.Time(t)
	return tt.String()
}
