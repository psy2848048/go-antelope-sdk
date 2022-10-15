package types

import (
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	timeFormat := "2006-01-02T15:04:05.999999999"
	timeStr := strings.ReplaceAll(string(b), `"`, "")

	if strings.Contains(timeStr, ".") {
		// 2022-10-15T07:13:42.000
		timeStr = strings.Join([]string{timeStr, "000000"}, "")
	} else {
		// 2022-10-15T07:16:04
		timeStr = strings.Join([]string{timeStr, ".000000000"}, "")
	}

	ret, err := time.Parse(timeFormat, timeStr)
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
