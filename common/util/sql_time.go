package util

import (
	"encoding/binary"
	"strconv"
	"time"
)

type SqlTime time.Time

// MarshalJSON 实现Marshaler
func (t *SqlTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(*t)
	unix := tt.Unix()
	if unix <= 0 {
		unix = 0
	}
	return Int64ToBytes(unix), nil
}

func (t *SqlTime) UnmarshalJSON(data []byte) error {

	if string(data) == "0" || string(data) == "null" {
		return nil
	}

	tmp, _ := strconv.ParseInt(string(data), 10, 64)

	*t = SqlTime(time.Unix(tmp, 0))

	return nil
}

func (t *SqlTime) Unix() int64 {
	if t == nil {
		return 0
	}
	unix := time.Time(*t).Unix()
	if unix <= 0 {
		unix = 0
	}
	return unix
}

func (t *SqlTime) FormatYD() string {
	if t == nil {
		return ""
	}

	tm := time.Time(*t)
	return tm.Format("2006-01-02")
}

func (t *SqlTime) FormatYH() string {
	if t == nil {
		return ""
	}

	tm := time.Time(*t)
	return tm.Format("2006-01-02 15")
}

func (t *SqlTime) FormatYM() string {
	if t == nil {
		return ""
	}

	tm := time.Time(*t)
	return tm.Format("2006-01-02 15:04")
}

func (t *SqlTime) FormatYS() string {
	if t == nil {
		return ""
	}

	tm := time.Time(*t)
	return tm.Format("2006-01-02 15:04:05")
}

func (t *SqlTime) FormatHS() string {
	if t == nil {
		return ""
	}

	tm := time.Time(*t)
	return tm.Format("15:04:05")
}

func (t *SqlTime) FormatHM() string {
	if t == nil {
		return ""
	}

	tm := time.Time(*t)
	return tm.Format("15:04")
}

func (t *SqlTime) Time() *time.Time {
	if t == nil {
		return nil
	}

	tm := time.Time(*t)
	return &tm
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
