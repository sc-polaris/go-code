package main

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
	maxDepth, node := 0, -1
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		if depth > maxDepth {
			maxDepth, node = depth, x
		}
		f[x] = fa
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
			}
		}
	}
	dfs(0, -1, 1)
	maxDepth = 0
	dfs(node, -1, 1)

	var path []int
	for node != -1 {
		path = append(path, node)
		node = f[node]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}

func main() {

}
