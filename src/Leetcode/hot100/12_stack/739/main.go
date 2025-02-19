package main

import "slices"

// 从左边往右
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var st []int
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}

// 从右往左
func dailyTemperatures2(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var st []int
	for i, t := range slices.Backward(temperatures) {
		for len(st) > 0 && t >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}
