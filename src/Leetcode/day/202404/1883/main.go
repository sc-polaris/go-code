package main

func minSkips(dist []int, speed, hoursBefore int) int {
	sumDist := 0
	for _, d := range dist {
		sumDist += d
	}
	if sumDist > speed*hoursBefore {
		return -1
	}

	n := len(dist)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j < 0 { // 递归边界
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := (dfs(i, j-1) + dist[j] + speed - 1) / speed * speed
		if i > 0 {
			res = min(res, dfs(i-1, j-1)+dist[j])
		}
		*p = res // 记忆化
		return res
	}
	for i := 0; ; i++ {
		if dfs(i, n-2)+dist[n-1] <= speed*hoursBefore {
			return i
		}
	}
}
