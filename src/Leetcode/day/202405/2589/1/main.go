package main

import "slices"

/*
	你有一台电脑，它可以 同时 运行无数个任务。给你一个二维整数数组 tasks ，其中 tasks[i] = [starti, endi, durationi] 表示
	第 i 个任务需要在 闭区间 时间段 [starti, endi] 内运行 durationi 个整数时间点（但不需要连续）。

	当电脑需要运行任务时，你可以打开电脑，如果空闲时，你可以将电脑关闭。

	请你返回完成所有任务的情况下，电脑最少需要运行多少秒。
*/

/*
	方法一：贪心+暴力

	1. 按照区间右端点从小到大排序。
	2. 排序后，对于区间 task[i] 来说，它右侧的任务要么和它没有交集，要么包含它的一部分后缀。
	3. 遍历排序后的任务，先冲击区间内的已运行的电脑运行时间点，如果个数小于 duration，则需要新增时间点。
	   依据提示 2，尽量把新增的时间点安排在区间 [start,end] 的后缀上，这样下一个区间就能统计到更多已运行的时间点。

*/

func findMinimumTime(tasks [][]int) (ans int) {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	run := make([]bool, tasks[len(tasks)-1][1]+1)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		for _, b := range run[start : end+1] { // 去掉运行中的时间点
			if b {
				d--
			}
		}
		for i := end; d > 0; i-- { // 剩余的 d 填充后缀
			if !run[i] {
				run[i] = true
				d--
				ans++
			}
		}
	}
	return
}
