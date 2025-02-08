package main

func climbStairs(n int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		p := &memo[i]
		if *p == 0 {
			*p = dfs(i-1) + dfs(i-2)
		}
		return *p
	}
	return dfs(n)
}

func climbStairs2(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func climbStairs3(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}
