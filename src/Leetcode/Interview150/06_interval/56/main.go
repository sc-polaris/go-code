package main

import "slices"

func merge(intervals [][]int) (ans [][]int) {
	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}
	return
}
