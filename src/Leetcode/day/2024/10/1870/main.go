package main

import (
	"math"
	"slices"
	"sort"
)

/*
	给你一个浮点数 hour ，表示你到达办公室可用的总通勤时间。要到达办公室，你必须按给定次序乘坐 n 趟列车。另给你一个长度为 n 的整数数组 dist ，其中 dist[i] 表示第 i 趟列车的行驶距离（单位是千米）。

	每趟列车均只能在整点发车，所以你可能需要在两趟列车之间等待一段时间。
	· 例如，第 1 趟列车需要 1.5 小时，那你必须再等待 0.5 小时，搭乘在第 2 小时发车的第 2 趟列车。
	返回能满足你准时到达办公室所要求全部列车的 最小正整数 时速（单位：千米每小时），如果无法准时到达，则返回 -1 。

	生成的测试用例保证答案不超过 107 ，且 hour 的 小数点后最多存在两位数字 。
*/

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	h100 := int(math.Round(hour * 100)) // 下面不会用到任何浮点数
	delta := h100 - (n-1)*100
	if delta <= 0 { // 无法到达终点
		return -1
	}

	maxDist := slices.Max(dist)
	if h100 <= n*100 { // 特判
		return max(maxDist, (dist[n-1]*100-1)/delta+1)
	}

	h := h100 / (n * 100)
	return 1 + sort.Search((maxDist-1)/h, func(v int) bool {
		v++
		t := 0
		for _, d := range dist[:n-1] {
			t += (d-1)/v + 1
		}
		return (t*v+dist[n-1])*100 <= h100*v
	})
}
