package main

/*
	现有一棵 无向 树，树中包含 n 个节点，按从 0 到 n - 1 标记。树的根节点是节点 0 。给你一个长度为 n - 1 的二维整数数组 edges，其中 edges[i] = [ai, bi] 表示树中节点 ai 与节点 bi 之间存在一条边。

	如果一个节点的所有子节点为根的
	子树包含的节点数相同，则认为该节点是一个 好节点。

	返回给定树中 好节点 的数量。

	子树 指的是一个节点以及它所有后代节点构成的一棵树。
*/

func countGoodNodes(edges [][]int) (ans int) {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		size, sz0, ok := 1, 0, true
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			sz := dfs(y, x)
			if sz0 == 0 {
				sz0 = sz // 记录第一个儿子子树的大小
			} else if sz != sz0 { // 存在大小不一样的儿子子树
				ok = false // 注意不能 break，其他子树 y 仍然要递归
			}
			size += sz
		}
		if ok {
			ans++
		}
		return size
	}
	dfs(0, -1)
	return
}
