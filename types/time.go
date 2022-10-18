package types

import (
	"fmt"
	"strings"
	"time"
)

type Time time.Time

func NewTimeFromString(strTime string) (Time, error) {
	timeFormat1 := "2006-01-02T15:04:05"
	timeFormat2 := "2006-01-02T15:04:05.999999999"

	var ret time.Time
	var err error

	if !strings.Contains(strTime, ".") {
		ret, err = time.Parse(timeFormat1, strTime)
	} else {
		splitedTime := strings.Split(strTime, ".")
		concattedTime := ""

		if len(splitedTime[1]) == 3 {
			concattedTime = strings.Join([]string{strTime, "000000"}, "")
		} else if len(splitedTime[1]) == 9 {
			concattedTime = strTime
		}

		ret, err = time.Parse(timeFormat2, concattedTime)
	}

	if err != nil {
		return Time{}, err
	}

	return Time(ret), nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	ret := fmt.Sprintf(`"%s"`, t.String())
	return []byte(ret), nil
}

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
	return tt.Format("2006-01-02T15:04:05")
}
