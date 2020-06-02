package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	TimeLocation            = "Asia/Shanghai"
	TimeLayoutString        = "2006-01-02 15:04:05"
	TimeDateLayoutString    = "2006-01-02" //only date layout
	TimeDateIntLayoutString = "20060102"   //only date layout

)

func checkDateLocation(date1, date2 time.Time) bool {
	s := date1.Location().String()
	s2 := date2.Location().String()
	//return strings.EqualFold(date1.Location().String(), date2.Location().String())
	return strings.EqualFold(s, s2)

}

//結束日期距離開始日期差幾天
//只考慮日期，不考慮時間
//注意應相同時區，不同時區回傳錯誤
func SubDate(begin, end time.Time) (int, error) {

	if !checkDateLocation(begin, end) {
		return -1, errors.New("time zone not match")
	}

	return int(end.Sub(begin).Hours() / 24), nil
}

//SubSecond returns end 減去 begin 是多少秒 ，有可能是負的
func SubSecond(begin, end time.Time) (int64, error) {

	if !checkDateLocation(begin, end) {
		return -1, errors.New("time zone not match")
	}

	return int64(end.Sub(begin).Seconds()), nil
}

//比較日期是否相同
//注意應相同時區，不同時區回傳錯誤
func EqualDate(date1, date2 time.Time) (bool, error) {
	if !checkDateLocation(date1, date2) {
		return false, errors.New("time zone not match")
	}

	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2, nil
}

//IsBetween returns if target time is between begin and end.
// 如果 target time == begin 或 target time == end , 也算
func IsBetween(targetTime, begin, end time.Time) (bool,error) {

	//begin 必須<= end
	if begin.After(end){
		return false,fmt.Errorf("begin time must <= end time")
	}

	if targetTime.After(begin) && targetTime.Before(end) {
		return true,nil
	}

	if targetTime.Equal(begin) || targetTime.Equal(end) {
		return true,nil
	}
	return false,nil
}

func GetDefaultLocation() *time.Location {
	location, _ := time.LoadLocation(TimeLocation)
	return location
}

//GetNow return now with default location
func GetNowWithDefaultLocation() time.Time {
	loc := GetDefaultLocation()
	return time.Now().In(loc)
}

//GetTimeWithDefaultLocation returns time with default location
//value pattern "2019-06-01 00:00:00"
func GetTimeWithDefaultLocation(value string) (time.Time, error) {
	location := GetDefaultLocation()
	return time.ParseInLocation(TimeLayoutString, value, location)
}

//GetTimeByHourString returns time with hour string ex:"01:01:01"
//return pattern 2019-01-01 00:00:00
//hourStr pattern "01:01:01"
func GetTimeByHourString(t time.Time, hourStr string) (time.Time, error) {

	dateStr := t.Format(TimeDateLayoutString) //取得日期字串

	fullStr := dateStr + " " + hourStr

	loc := GetDefaultLocation()

	return time.ParseInLocation(TimeLayoutString, fullStr, loc)

}
func GetTimeByTimeString(timeStr string) (time.Time, error) {

	loc := GetDefaultLocation()

	return time.ParseInLocation(TimeLayoutString, timeStr, loc)

}

func GetTimeByParam(t time.Time, hour, min, sec string) (time.Time, error) {

	dateStr := t.Format(TimeDateLayoutString) //取得日期字串

	fullStr := dateStr + " " + hour + ":" + min + ":" + sec

	loc := GetDefaultLocation()

	return time.ParseInLocation(TimeLayoutString, fullStr, loc)

}

func ParseWithLayoutString(t time.Time) string {

	return t.Format(TimeLayoutString)
}
