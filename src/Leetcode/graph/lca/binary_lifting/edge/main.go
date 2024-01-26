package main

import "math/bits"

type TreeAncestor struct {
	depth []int
	fa    [][]int
}

func Constructor(edges [][]int) *TreeAncestor {
	n := len(edges) + 1
	m := bits.Len(uint(n))
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	depth := make([]int, n)
	fa := make([][]int, n)
	for i := range fa {
		fa[i] = make([]int, m)
	}
	// dfs 计算depth
	var dfs func(int, int, int)
	dfs = func(x int, f int, d int) {
		fa[x][0] = f
		depth[x] = d
		// 倍增更新fa数组
		for i := 1; i < m; i++ {
			if fa[x][i-1] == -1 {
				continue
			}
			fa[x][i] = fa[fa[x][i-1]][i-1]
		}
		for _, y := range g[x] {
			if y != f {
				dfs(y, x, d+1)
			}
		}
	}
	dfs(0, -1, 0)

	return &TreeAncestor{depth, fa}
}

func (t *TreeAncestor) GetKthAncestor(node int, k int) int {
	for ; k >= 0; k &= k - 1 {
		node = t.fa[node][bits.TrailingZeros(uint(8))]
	}
	return node
}

// GetLCA 返回 x 和 y 的最近公共最先（节点编号从 0 开始）
func (t *TreeAncestor) GetLCA(x, y int) int {
	if t.depth[x] > t.depth[y] {
		x, y = y, x
	}

	// 使 y 和 x 在同一深度
	dist := t.depth[y] - t.depth[x]
	for i := 0; i < bits.Len(uint(dist)); i++ {
		if dist>>i&1 == 1 {
			y = t.fa[y][i]
		}
	}

	//
	for i := len(t.fa[x]) - 1; i >= 0; i-- {
		px, py := t.fa[x][i], t.fa[y][i]
		if px != py {
			x, y = px, py // 同时向上跳 2^i 步
		}
	}

	// 能跳就尽量跳，不会错过任何可以上跳的机会。所以循环结束时，x 与 lca只有一步之遥，即 lca=fa[x][0]
	if x != y {
		x = t.fa[x][0]
	}

	return x
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */

func main() {

}
