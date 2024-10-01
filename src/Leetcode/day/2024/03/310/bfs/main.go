package main

/*

	设 dist[x][y] 表示节点 x 到节点 y 的距离，假设树中距离最长的两个节点为 (x,y)，它们之间的距离为
	maxdist=dist[x][y]，则可以推出以任意节点构成的树最小高度一定为 minheight=[maxdist/2]，且最
	小高度的树根节点一定在节点 x 到节点 y 的路径上。

	解决方法：
	1. 以任意节点 p 出发，利用 bfs 或者 dfs 找到以 p 为起点的最长路径的终点 x
	2. 以节点 x 出发，找到以 x 为起点的最长路径的终点 y；
	3. x 到 y 之间的路径即为途中的最长路径，找到路径的中间节点即为根节点。

*/

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	f := make([]int, n)
	bfs := func(start int) (x int) {
		st := make([]bool, n)
		st[start] = true
		q := []int{start}
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range g[x] {
				if !st[y] {
					st[y] = true
					f[y] = x
					q = append(q, y)
				}
			}
		}
		return
	}
	x := bfs(0) // 找到与节点 0 最远的节点 x
	y := bfs(x) // 找到与节点 x 最远的点 y

	var path []int
	f[x] = -1
	for y != -1 {
		path = append(path, y)
		y = f[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}

func main() {

}
