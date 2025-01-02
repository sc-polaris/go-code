package main

import "github.com/emirpasic/gods/trees/redblacktree"

/*
	实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。

	当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。

	日程可以用一对整数 startTime 和 endTime 表示，这里的时间是半开区间，即 [startTime, endTime), 实数 x 的范围为，  startTime <= x < endTime 。

	实现 MyCalendar 类：
	· MyCalendar() 初始化日历对象。
	· boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回 false 并且不要将该日程安排添加到日历中。
*/

type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(-1, -1) // 哨兵
	return MyCalendar{t}
}

func (c *MyCalendar) Book(startTime int, endTime int) bool {
	floor, _ := c.Floor(startTime)
	if floor.Value.(int) > startTime { // [start,end) 左侧区间的右端点超过了 start
		return false
	}
	if it := c.IteratorAt(floor); it.Next() && it.Key().(int) < endTime { // [start,end) 右侧区间的左端点小于 end
		return false
	}
	c.Put(startTime, endTime) // 可以插入区间 [start,end)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
