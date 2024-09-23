package main

/*
	给你一个正整数数组 values，其中 values[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的 距离 为 j - i。
	一对景点（i < j）组成的观光组合的得分为 values[i] + values[j] + i - j ，也就是景点的评分之和 减去 它们两者之间
	的距离。
	返回一对观光景点能取得的最高分。
*/

func maxScoreSightseeingPair(values []int) (ans int) {
	mx := 0 // j 左边的 values[i] + i 的最大值
	for j, v := range values {
		ans = max(ans, mx+v-j)
		mx = max(mx, v+j)
	}
	return
}
