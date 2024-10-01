package main

func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
	}

	ans := make([][]int, n)
	vis := make([]int, n)
	start := 0
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = start + 1 // 避免重复访问
		for _, y := range g[x] {
			if vis[y] != start+1 {
				ans[y] = append(ans[y], start) // start 是访问到的点的祖先
				dfs(y)                         // 只递归没有访问过的点
			}
		}
	}

	for ; start < n; start++ {
		dfs(start) // 从 start 开始 dfs
	}
	return ans
}
