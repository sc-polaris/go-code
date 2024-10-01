package main

import (
	"strconv"
	"time"
)

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	return t.YearDay()
}

func dayOfYear2(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		days[1]++
	}

	ans := day
	for _, d := range days[:month-1] {
		ans += d
	}
	return ans
}

func main() {

}
