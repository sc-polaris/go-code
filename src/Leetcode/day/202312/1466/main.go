package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	g := make([][][2]int, n)
	for _, e := range connections {
		x, y := e[0], e[1]
		// 1 标记原方向边，0 标记反方向的边
		g[x] = append(g[x], [2]int{y, 1})
		g[y] = append(g[y], [2]int{x, 0})
	}

	var dfs func(int, int) int
	dfs = func(x int, fa int) (res int) {
		for _, v := range g[x] {
			if y, path := v[0], v[1]; y != fa {
				res += path + dfs(y, x)
			}
		}
		return
	}

	return dfs(0, -1)
}

func main() {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}
