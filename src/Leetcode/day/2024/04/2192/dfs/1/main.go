package main

// 逆向 dfs 把边都反向

func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[y] = append(g[y], x) // 反向建图
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true // 避免重复访问
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y) // 只递归访问没有访问过的点
			}
		}
	}
	ans := make([][]int, n)
	for i := range ans {
		clear(vis)
		dfs(i)         // 从 i 开始 dfs
		vis[i] = false // ans[i] 不含 i
		for j, b := range vis {
			if b {
				ans[i] = append(ans[i], j)
			}
		}
	}
	return ans
}
