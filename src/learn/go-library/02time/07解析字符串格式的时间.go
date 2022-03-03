package main

import (
	"fmt"
	"time"
)

/*
对于从文本的时间表示中解析出时间对象，time包中提供了time.Parse和time.ParseInLocation两个函数。

其中time.Parse在解析时不需要额外指定时区信息。
time.ParseInLocation函数需要在解析时额外指定时区信息。
*/

// parseDemo 指定时区解析时间
func parseDemo() {
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2022/10/05 11:25:20")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0800 CST
}

// parseDemo2 解析时间
func parseDemo2() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	parseDemo2()
}
