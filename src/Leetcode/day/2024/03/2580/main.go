package main

import "slices"

func countWays(ranges [][]int) int {
	slices.SortFunc(ranges, func(a, b []int) int { return a[0] - b[0] })
	ans, maxR := 1, -1
	for _, p := range ranges {
		if p[0] > maxR { // 无法合并
			ans = ans * 2 % 1_000_000_007 // 新区间
		}
		maxR = max(maxR, p[1]) // 合并
	}
	return ans
}
