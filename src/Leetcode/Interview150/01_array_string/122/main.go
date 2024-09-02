package main

import "math"

// 记忆化搜索 dfs(i,0) 第 i 天没有股票的最大利润 dfs(i,1) 第 i 天有股票的最大利润
func maxProfit(prices []int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i int, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return
		}

		p := &memo[i][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0)
}

// dp 1
func maxProfit1(prices []int) int {
	n := len(prices)
	f := make([][2]int, n+1)
	f[0][1] = math.MinInt
	for i, p := range prices {
		f[i+1][0] = max(f[i][0], f[i][1]+p)
		f[i+1][1] = max(f[i][1], f[i][0]-p)
	}
	return f[n][0]
}

// dp 2
func maxProfit2(prices []int) int {
	f0, f1 := 0, math.MinInt
	for _, p := range prices {
		f0, f1 = max(f0, f1+p), max(f1, f0-p)
	}
	return f0
}
