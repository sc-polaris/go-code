package main

import "slices"

/*
	给你一个整数数组 salary ，数组里每个数都是 唯一 的，其中 salary[i] 是第 i 个员工的工资。
	请你返回去掉最低工资和最高工资以后，剩下员工工资的平均值。

*/

func average(salary []int) float64 {
	s := 0
	for _, x := range salary {
		s += x
	}
	return float64(s-slices.Min(salary)-slices.Max(salary)) / float64(len(salary)-2)
}
