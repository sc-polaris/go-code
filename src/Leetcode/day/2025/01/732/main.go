package main

import "github.com/emirpasic/gods/v2/trees/redblacktree"

/*
	当 k 个日程存在一些非空交集时（即, k 个日程包含了一些相同时间），就会产生 k 次预订。

	给你一些日程安排 [startTime, endTime) ，请你在每个日程安排添加后，返回一个整数 k ，表示所有先前日程安排会产生的最大 k 次预订。

	实现一个 MyCalendarThree 类来存放你的日程安排，你可以一直添加新的日程安排。
	· MyCalendarThree() 初始化对象。
	· int book(int startTime, int endTime) 返回一个整数 k ，表示日历中存在的 k 次预订的最大值。
*/

type MyCalendarThree struct {
	*redblacktree.Tree[int, int]
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.New[int, int]()}
}

func (c *MyCalendarThree) Add(k, v int) {
	if val, ok := c.Get(k); ok {
		v += val
	}
	c.Put(k, v)
}

func (c *MyCalendarThree) Book(startTime int, endTime int) (ans int) {
	c.Add(startTime, 1)
	c.Add(endTime, -1)

	s := 0
	for _, v := range c.Values() {
		s += v
		if s > ans {
			ans = s
		}
	}
	return
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
