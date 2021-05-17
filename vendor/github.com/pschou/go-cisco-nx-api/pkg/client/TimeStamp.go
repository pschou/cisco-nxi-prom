package client

import (
	"time"
)

type TimeStamp uint64

func ParseTimeStamp(text string) (t TimeStamp, err error) {
	var val time.Time
	if len(text) > 0 && text[0] == ' ' {
		text = "0" + text[1:]
	}
	switch len(text) {
	case 23, 24:
		val, err = time.ParseInLocation("Mon Jan 2 15:04:05 2006", text, time.UTC)
	case 19:
		val, err = time.ParseInLocation("01/02/2006 15:04:05", text, time.UTC)
	case 10:
		val, err = time.ParseInLocation("01/02/2006", text, time.UTC)
	}
	if err == nil {
		t = TimeStamp(val.Unix())
	}
	return
}

func (t *TimeStamp) UnmarshalText(text []byte) (err error) {
	*t, err = ParseTimeStamp(string(text))
	return
}
func (t TimeStamp) Time() time.Time {
	return time.Unix(int64(t), 0).UTC()
}
func FromTime(val time.Time) TimeStamp {
	return TimeStamp(val.Unix())
}
func (t TimeStamp) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}
func (t TimeStamp) String() string {
	if t == 0 {
		return ""
	}
	val := time.Unix(int64(t), 0).UTC()
	return val.Format("01/02/2006 15:04:05")
}
