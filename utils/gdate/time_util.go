package utils

import (
	"time"
)

const BASE_TIME = "2006-01-02 15:04:05"

const BASE_TIME_DATE = "2006-01-02"

func Now() *time.Time {
	t := time.Now()
	return &t
}

func NowStr() string {
	return time.Now().Format(BASE_TIME)
}

func NowDateStr() string {
	return time.Now().Format(BASE_TIME_DATE)
}

func NowWithDuration(s string) *time.Time {
	now := time.Now()
	if s != "" {
		duration, _ := time.ParseDuration(s)
		t := now.Add(duration)
		return &t
	}

	return &now
}

func NowWithDurationStr(s string) string {
	now := time.Now()
	if s != "" {
		duration, _ := time.ParseDuration(s)
		t := now.Add(duration)
		return t.Format(BASE_TIME)
	}

	return now.Format(BASE_TIME)
}

func ParseTimeStr(timeStr string) (*time.Time, error) {
	if t, err := time.Parse(BASE_TIME, timeStr); err != nil {
		return nil, err

	} else {
		return &t, nil
	}
}

func IsSameDay(t1, t2 *time.Time) bool {
	t1Str := t1.Format(BASE_TIME_DATE)
	t2Str := t2.Format(BASE_TIME_DATE)
	return t1Str == t2Str
}
