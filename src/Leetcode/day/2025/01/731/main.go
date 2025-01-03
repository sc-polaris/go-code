package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

/*
	实现一个程序来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。

	当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生 三重预订。

	事件能够用一对整数 startTime 和 endTime 表示，在一个半开区间的时间 [startTime, endTime) 上预定。实数 x 的范围为  startTime <= x < endTime。

	实现 MyCalendarTwo 类：
	· MyCalendarTwo() 初始化日历对象。
	· boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。
*/

/*
	在 start 计数 cnt[start] 加 1，表示从 start 预定的数目加 1；从 end 计数 cnt[end] 减 1，表示从 end 开始预定的数目减 1。
	我们可以利用差分的思想，将每个时间点的预定情况记录下来，然后遍历所有时间点，统计当前时间点的预定情况，
	如果预定次数超过 2 次，则返回 false。否则，返回 true。
*/

type MyCalendarTwo struct {
	*redblacktree.Tree[int, int]
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{redblacktree.New[int, int]()}
}

func (c *MyCalendarTwo) add(key, value int) {
	if v, ok := c.Get(key); ok {
		c.Put(key, v+value)
	} else {
		c.Put(key, value)
	}
}

func (c *MyCalendarTwo) Book(startTime int, endTime int) bool {
	c.add(startTime, 1)
	c.add(endTime, -1)

	s := 0
	for _, v := range c.Values() {
		s += v
		if s > 2 {
			c.add(startTime, -1)
			c.add(endTime, 1)
			return false
		}
	}
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
