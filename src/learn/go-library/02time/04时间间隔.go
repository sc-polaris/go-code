package main

import (
	"fmt"
	"time"
)

/*
time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
time.Duration表示一段时间间隔，可表示的最长时间段大约290年。

time 包中定义的时间间隔类型的常量如下：

const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
例如：time.Duration表示1纳秒，time.Second表示1秒。

· Add：增加时间后的时间
	func (t Time) Add(d Duration) Time
· Sub：求两个时间之间的差值：
	func (t Time) Sub(u Time) Duration
	返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，
	将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。
· Equal：
	func (t Time) Equal(u Time) bool
	判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。
	本方法和用t==u不同，这种方法还会比较地点和时区信息。
· Before：
	func (t Time) Before(u Time) bool
	如果t代表的时间点在u之前，返回真；否则返回假。
· After：
	func (t Time) After(u Time) bool
	如果t代表的时间点在u之后，返回真；否则返回假。
*/

func main() {
	// 举个例子，求一个小时之后的时间：
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}
