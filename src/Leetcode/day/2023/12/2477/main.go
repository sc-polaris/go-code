package main

func minimumFuelCost(roads [][]int, seats int) (res int64) {
	g := make([][]int, len(roads)+1)
	for _, e := range roads {
		x, y := e[0], e[1]
		g[x] = append(g[x], y) // 记录每个点的邻居
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int
	dfs = func(x int, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y != fa { // 递归子节点
				size += dfs(y, x) // 统计子树大小
			}
		}
		if x > 0 {
			res += int64((size-1)/seats + 1) // ceil(size/seats)
		}
		return size
	}
	dfs(0, -1)
	return
}

func main() {

}
