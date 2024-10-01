package main

import "time"

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day int, month int, year int) string {
	days := 0

	// 1970 年 12 月 31 日是星期四 1968是闰年
	// 年的天数
	days += 365*(year-1971) + (year-1-1968)/4

	// 月份之前的天数
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}

	// 月份中的天数
	days += day
	return week[(days+3)%7]
}

func dayOfTheWeek2(day, month, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}

func main() {

}
