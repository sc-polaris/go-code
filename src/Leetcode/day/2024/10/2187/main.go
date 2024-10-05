package main

import (
	"slices"
	"sort"
)

/*
	给你一个数组 time ，其中 time[i] 表示第 i 辆公交车完成 一趟旅途 所需要花费的时间。

	每辆公交车可以 连续 完成多趟旅途，也就是说，一辆公交车当前旅途完成后，可以 立马开始 下一趟旅途。每辆公交车 独立 运行，也就是说可以同时有多辆公交车在运行且互不影响。

	给你一个整数 totalTrips ，表示所有公交车 总共 需要完成的旅途数目。请你返回完成 至少 totalTrips 趟旅途需要花费的 最少 时间。
*/

func minimumTime(time []int, totalTrips int) int64 {
	minT := slices.Min(time)
	maxT := slices.Max(time)
	avg := (totalTrips-1)/len(time) + 1
	left := minT * avg
	right := min(maxT*avg, minT*totalTrips)
	return int64(left + sort.Search(right-left, func(x int) bool {
		x += left
		sum := 0
		for _, t := range time {
			sum += x / t
			if sum >= totalTrips {
				return true
			}
		}
		return false
	}))
}
