package main

import "math"

func numSquares(n int) int {
	qx := int(math.Sqrt(float64(n)))
	memo := make([][]int, qx+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			if j == 0 {
				return 0
			}
			return math.MaxInt
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if j < i*i {
			*p = dfs(i-1, j)
		} else {
			*p = min(dfs(i-1, j), dfs(i, j-i*i)+1)
		}
		return *p
	}
	return dfs(qx, n)
}

func numSquares2(n int) int {
	qx := int(math.Sqrt(float64(n)))
	f := make([][]int, qx+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		f[0][i] = math.MaxInt
	}
	for i := 1; i*i <= n; i++ {
		for j := 0; j <= n; j++ {
			if j < i*i {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-i*i]+1)
			}
		}
	}
	return f[qx][n]
}

func numSquares3(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			f[j] = min(f[j], f[j-i*i]+1)
		}
	}
	return f[n]
}
