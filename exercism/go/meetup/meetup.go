package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) (day int) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	totalDays := daysIn(month, year)
	lastDay := time.Date(year, month, totalDays, 0, 0, 0, 0, time.UTC)
	firstTeenthDate := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)

	switch wSched {
	case First:
		difference := int(wDay+7-firstDay.Weekday()) % 7
		day = firstDay.Day() + difference
	case Second:
		difference := int(wDay+7-firstDay.Weekday())%7 + 7
		day = firstDay.Day() + difference
	case Third:
		difference := int(wDay+7-firstDay.Weekday())%7 + 2*7
		day = firstDay.Day() + difference
	case Fourth:
		difference := int(wDay+7-firstDay.Weekday())%7 + 3*7
		day = firstDay.Day() + difference
	case Last:
		difference := int(7+lastDay.Weekday()-wDay) % 7
		day = totalDays - difference
	case Teenth:
		difference := int(wDay+7-firstTeenthDate.Weekday()) % 7
		day = firstTeenthDate.Day() + difference
	default:
		panic("Invalid Schedule")
	}

	return day
}
