package main

import (
	"fmt"
	"math/bits"
)

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ y, w int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]-1
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	const mx = 14 // 2^14 > 10^4

	// fa[i][j]: 节点i的2^j祖父节点
	fa := make([][mx]int, n)
	// 记录每个节点从根节点的当前节点的深度
	depth := make([]int, n)
	// count[i][j] 根节点到节点i权重为j的路径数量
	count := make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, 26)
	}

	// 预处理，计算出每个节点的父节点以及深度
	var dfs func(int, int, int, [26]int)
	dfs = func(x int, f int, d int, cnt [26]int) {
		fa[x][0] = f
		depth[x] = d
		copy(count[x], cnt[:])
		// 倍增更新fa数组
		for i := 1; i < mx && fa[x][i-1] != -1; i++ {
			fa[x][i] = fa[fa[x][i-1]][i-1]
		}
		for _, e := range g[x] {
			if y, w := e.y, e.w; y != f {
				cnt[w]++
				dfs(y, x, d+1, cnt)
				cnt[w]--
			}
		}
	}
	dfs(0, -1, 0, [26]int{})

	// 倍增法 计算 LCA 模版（这里返回最小操作次数）
	var lca func(int, int) int
	lca = func(x int, y int) int {
		if depth[x] > depth[y] {
			x, y = y, x
		}
		// 令 y 和 x 在同一个深度
		dist := depth[y] - depth[x]
		fx, fy := x, y
		for i := 0; i < bits.Len(uint(dist)); i++ {
			if dist>>i&1 == 1 {
				fy = fa[fy][i]
			}
		}
		for i := mx - 1; i >= 0; i-- {
			if px, py := fa[fx][i], fa[fy][i]; px != py {
				fx, fy = px, py
			}
		}
		if fx != fy {
			fx = fa[fx][0]
		}

		maxVal := 0
		for i := 0; i < 26; i++ {
			maxVal = max(maxVal, count[x][i]+count[y][i]-2*count[fx][i])
		}

		return depth[x] + depth[y] - 2*depth[fx] - maxVal
	}

	var ans []int
	for _, q := range queries {
		ans = append(ans, lca(q[0], q[1]))
	}

	return ans
}

func main() {
	n := 8
	edges := [][]int{
		{1, 2, 6},
		{1, 3, 4},
		{2, 4, 6},
		{2, 5, 3},
		{3, 6, 6},
		{3, 0, 8},
		{0, 7, 2},
	}
	queries := [][]int{
		//{4, 6},
		//{0, 4},
		{6, 5},
		{7, 4},
	}
	fmt.Println(minOperationsQueries(n, edges, queries))
}
