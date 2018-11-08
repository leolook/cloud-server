package util

import (
	log "lib/xlog"
	"time"
)

var Time *FetchTime

type FetchTime struct{}

//获取时区
func (f *FetchTime) Location() (*time.Location, error) {
	return time.LoadLocation("Asia/Chongqing")
}

//获取当前时间
func (f *FetchTime) Now() time.Time {
	loc, err := f.Location()
	if err != nil {
		log.Errorf("Failed to fetch current time,[err=%v]", err)
		return time.Now()
	}
	current := time.Now().In(loc)
	return current
}

//获取当前时间戳
func (f *FetchTime) Unix() int64 {
	return f.Now().Unix()
}

//字符串时间 2006-01-02 15:04
func (f *FetchTime) Unix1() (int64, error) {
	str := f.Format1()
	t, err := f.ParseStr(str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//字符串时间 2006-01-02 15:04:05
func (f *FetchTime) Unix2() (int64, error) {
	str := f.Format2()
	t, err := f.ParseStr2(str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//字符串时间 2006-01-02
func (f *FetchTime) Unix3() (int64, error) {
	str := f.Format3()
	t, err := f.ParseStr3(str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//字符串时间 2006-01-02 15:04
func (f *FetchTime) Format1() string {
	str := f.Now().Format("2006-01-02 15:04")
	return str
}

//字符串时间 2006-01-02 15:04:05
func (f *FetchTime) Format2() string {
	str := f.Now().Format("2006-01-02 15:04:05")
	return str
}

//字符串时间 2006-01-02
func (f *FetchTime) Format3() string {
	str := f.Now().Format("2006-01-02")
	return str
}

//解析字符串时间 2006-01-02 15:04
func (f *FetchTime) ParseStr(str string) (time.Time, error) {
	loc, _ := f.Location()
	return time.ParseInLocation("2006-01-02 15:04", str, loc)
}

//解析字符串时间 2006-01-02 15:04:05
func (f *FetchTime) ParseStr2(str string) (time.Time, error) {
	loc, _ := f.Location()
	return time.ParseInLocation("2006-01-02 15:04:05", str, loc)
}

//解析字符串时间 2006-01-02
func (f *FetchTime) ParseStr3(str string) (time.Time, error) {
	loc, _ := f.Location()
	return time.ParseInLocation("2006-01-02", str, loc)
}

//解析字符串时间 15:04:05
func (f *FetchTime) ParseStr4(str string) (time.Time, error) {
	loc, _ := f.Location()
	return time.ParseInLocation("15:04:05", str, loc)
}
