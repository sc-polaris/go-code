package main

import "math"

func coinChange(coins []int, amount int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		p := &memo[i][c]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans := dfs(n-1, amount)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func coinChange2(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0
	for i, x := range coins {
		for c := 0; c <= amount; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-x]+1)
			}
		}
	}
	ans := f[n][amount]
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func coinChange3(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, x := range coins {
		for c := x; c <= amount; c++ {
			f[c] = min(f[c], f[c-x]+1)
		}
	}
	ans := f[amount]
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}
