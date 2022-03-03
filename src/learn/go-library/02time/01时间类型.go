package main

import (
	"fmt"
	"time"
)

/*
Go 语言中使用time.Time类型表示时间。我们可以通过time.Now函数获取当前的时间对象，
然后从时间对象中可以获取到年、月、日、时、分、秒等信息。
*/

func main() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}
