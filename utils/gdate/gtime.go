package gdate

import (
	"fmt"
	"time"
)

//getDaysInYearByThisYear 年中的第几天
func getDaysInYearByThisYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1

	} else {
		total = total + d
	}
	return total
}

func CurrentDay(t time.Time) (current time.Time) {
	now := time.Now()
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		t.Hour(), t.Minute(), t.Second())
	timeFmt := "2006-01-02 15:04:05"
	//time.Parse(timeFmt, day)
	current, _ = time.ParseInLocation(timeFmt, day, time.Local)

	return
}

func ZeroDateTime() (current time.Time) {
	now := time.Now()
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		0, 0, 0)
	timeFmt := "2006-01-02 15:04:05"
	//current, _ = time.Parse(timeFmt, day)
	current, _ = time.ParseInLocation(timeFmt, day, time.Local)

	return
}

func DateTimeString(mtime string) (current time.Time) {

	timeFmt := "2006-01-02 15:04:05"
	//current, _ = time.Parse(timeFmt, day)
	tm, _ := time.ParseInLocation(timeFmt, mtime, time.Local)

	now := time.Now()
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		tm.Hour(), tm.Minute(), tm.Second())
	current, _ = time.ParseInLocation(timeFmt, day, time.Local)
	return
}

func TodayZeroDateTime() (current int64) {
	now := time.Now()
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		0, 0, 0)
	timeFmt := "2006-01-02 15:04:05"
	//current, _ = time.Parse(timeFmt, day)
	currentTime, _ := time.ParseInLocation(timeFmt, day, time.Local)
	current = currentTime.Unix()
	return
}
func YesterdayZeroDateTime() (current int64) {
	now := time.Now().AddDate(0, 0, -1)
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		0, 0, 0)
	timeFmt := "2006-01-02 15:04:05"
	//current, _ = time.Parse(timeFmt, day)
	currentTime, _ := time.ParseInLocation(timeFmt, day, time.Local)
	current = currentTime.Unix()
	return
}

func Day24() (current time.Time) {
	now := time.Now()
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		23, 59, 59)
	timeFmt := "2006-01-02 15:04:05"
	//current, _ = time.Parse(timeFmt, day)
	current, _ = time.ParseInLocation(timeFmt, day, time.Local)

	return
}

func Day9() (current int64) {
	now := time.Now().UTC().Add(8 * time.Hour)
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		9, 0, 0)
	timeFmt := "2006-01-02 15:04:05"
	nine, err := time.ParseInLocation(timeFmt, day, time.Local)
	if err != nil {
		return now.Unix()
	}
	nine = nine.UTC().Add(8 * time.Hour)
	return nine.Unix()
}

func Day22() (current int64) {
	now := time.Now().UTC().Add(8 * time.Hour)
	day := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		now.Year(), now.Month(), now.Day(),
		22, 30, 0)
	timeFmt := "2006-01-02 15:04:05"
	nine, err := time.ParseInLocation(timeFmt, day, time.Local)
	if err != nil {
		return now.Unix()
	}
	nine = nine.UTC().Add(8 * time.Hour)
	return nine.Unix()
}

func GetDate() string {
	return time.Now().Format("20060102")
}

func GetYesterDate() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Format("20060102")
}

func GetHourTime() string {
	return time.Now().Format("150405")
}

func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
