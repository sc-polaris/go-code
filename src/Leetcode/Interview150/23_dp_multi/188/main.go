package main

import "math"

// dfs(i, j, hold) 是一个递归函数，表示在第 i 天，已完成 j 次交易，且当前是否持有股票（由 hold 表示）时的最大利润。

func maxProfit(k int, prices []int) int {
	n := len(prices)
	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 标识没有计算
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, hold int) (res int) {
		if j < 0 {
			return math.MinInt / 2
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2
			}
			return
		}
		p := &memo[i][j][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 { // 继续持有股票 上一天不持有，今天买入
			return max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		}
		// 上一天不持有 上一天持有今天卖出
		return max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i])
	}
	return dfs(n-1, k, 0)
}

func maxProfit2(k int, prices []int) int {
	n := len(prices)
	f := make([][][2]int, n+1)
	for i := range f {
		f[i] = make([][2]int, k+2)
		for j := range f[i] {
			f[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
		}
	}
	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}
	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
		}
	}
	return f[n][k+1][0]
}

func maxProfit3(k int, prices []int) int {
	f := make([][2]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
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
