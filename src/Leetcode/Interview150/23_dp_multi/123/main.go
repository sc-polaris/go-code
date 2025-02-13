package main

import "math"

func maxProfit(prices []int) int {
	const k = 2
	var f [k + 2][2]int
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2 // 防止溢出
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
		}
	}
	return f[k+1][0]
}
