package main

import "slices"

/*
	设h=heights[i]是矩形的高度，那么矩形的宽度最大是多少？我们需要知道：
		在i左侧的小于 h的最近元素的下标left，如果不存在则为−1。求出了left，那么left+1就是在i左侧的大于等于 h的最近元素的下标。
		在i右侧的小于 h的最近元素的下标right，如果不存在则为n。求出了right，那么right−1就是在i右侧的大于等于 h的最近元素的下标。
*/

func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	left := make([]int, n)
	var st []int
	for i, x := range heights {
		for len(st) > 0 && x <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n)
	st = st[:0]
	for i, x := range slices.Backward(heights) {
		for len(st) > 0 && x <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	for i, h := range heights {
		ans = max(ans, h*(right[i]-left[i]-1))
	}
	return ans
}
