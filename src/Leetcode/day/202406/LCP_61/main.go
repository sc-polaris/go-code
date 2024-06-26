package main

import "cmp"

/*
	力扣城计划在两地设立「力扣嘉年华」的分会场，气象小组正在分析两地区的气温变化趋势，对于第 i ~ (i+1) 天的气温变化趋势，将根据以下规则判断：
	· 若第 i+1 天的气温 高于 第 i 天，为 上升 趋势
	· 若第 i+1 天的气温 等于 第 i 天，为 平稳 趋势
	· 若第 i+1 天的气温 低于 第 i 天，为 下降 趋势

	已知 temperatureA[i] 和 temperatureB[i] 分别表示第 i 天两地区的气温。 组委会希望找到一段天数尽可能多，且两地气温变化趋势相同的时间
	举办嘉年华活动。请分析并返回两地气温变化趋势相同的最大连续天数。

	即最大的 n，使得第 i~i+n 天之间，两地气温变化趋势相同
*/

func temperatureTrend(a []int, b []int) (ans int) {
	same := 0
	for i := 1; i < len(a); i++ {
		if cmp.Compare(a[i-1], a[i]) == cmp.Compare(b[i-1], b[i]) {
			same++
			ans = max(ans, same)
		} else {
			same = 0
		}
	}
	return
}
