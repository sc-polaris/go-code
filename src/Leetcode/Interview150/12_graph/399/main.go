package main

// bfs
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := make(map[string]int)
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	// 建图
	type edge struct {
		to     int
		weight float64
	}
	g := make([][]edge, len(id))
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		g[v] = append(g[v], edge{w, values[i]})
		g[w] = append(g[w], edge{v, 1 / values[i]})
	}

	bfs := func(x, y int) float64 {
		ratios := make([]float64, len(g))
		ratios[x] = 1
		q := []int{x}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			if v == y {
				return ratios[v]
			}
			for _, e := range g[v] {
				if w := e.to; ratios[w] == 0 {
					ratios[w] = ratios[v] * e.weight
					q = append(q, w)
				}
			}
		}
		return -1
	}

	var ans []float64
	for _, q := range queries {
		x, okX := id[q[0]]
		y, okY := id[q[1]]
		if !okX || !okY {
			ans = append(ans, -1)
		} else {
			ans = append(ans, bfs(x, y))
		}
	}
	return ans
}

// floyd
func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := make(map[string]int)
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	// 建图
	g := make([][]float64, len(id))
	for i := range g {
		g[i] = make([]float64, len(id))
	}
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		g[v][w] = values[i]
		g[w][v] = 1 / values[i]
	}

	// floyd
	for k := range g {
		for i := range g {
			for j := range g {
				if g[i][k] > 1e-6 && g[k][j] > 1e-6 {
					g[i][j] = g[i][k] * g[k][j]
				}
			}
		}
	}

	var ans []float64
	for _, q := range queries {
		x, okX := id[q[0]]
		y, okY := id[q[1]]
		if !okX || !okY || g[x][y] == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, g[x][y])
		}
	}
	return ans
}

/*
	带权并查集
	我们还可以考虑以并查集的方式存储节点之间的关系。设节点 x 的值（即对应变量的取值）为 v[x]。对于任意两点 x,y，
	假设它们在并查集中具有共同的父亲 f，且 v[x]/v[f]=a,v[y]/v[f]=b，则 v[x]/v[y]=a/b。

	在观察到这一点后，就不难利用并查集的思想解决此题。对于每个节点 x 而言，除了维护其父亲 f[x] 之外，还要维护其权
	值 w，其中「权值」定义为节点 x 的取值与父亲 f[x] 的取值之间的比值。换言之，我们有
					w[x] = v[x] / v[f[x]]

	下面，我们对并查集的两种操作的实现细节做出讨论。当查询节点 x 父亲时，如果 f[x]!=x，我们需要先找到 f[x] 的父亲
	father，并将 f[x] 更新为 father。此时，我们有
					w[x] <- v[x] / v[father]
						 = (v[x] / v[f[x]]) * (v[f[x]] / v[father])
						 = w[x]*w[f[x]]
	也就是说，我们要将 w[x] 更新为 w[x]*w[f[x]]。

	当合并两个节点 x,y 时，我们首先找到两者的父亲 fx,fy，并将 f[fx] 为 fy，此时，我们有
					w[fx] <- v[fx] / v[fy]
						  = (v[x]/w[x])/(v[y]/w[y])
						  = v[x]/v[y] * w[y]/w[x]
	也就是说，当在已有的图中添加一条方程式 v[x]/v[y] = k 时，需要将 w[fx] 更新为 k*w[y]/w[x]。
*/

func calcEquation3(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := make(map[string]int)
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	fa := make([]int, len(id))
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	merge := func(x, y int, val float64) {
		fx, fy := find(x), find(y)
		w[fx] = val * w[y] / w[x]
		fa[fx] = fy
	}

	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	var ans []float64
	for _, q := range queries {
		x, okX := id[q[0]]
		y, okY := id[q[1]]
		if okX && okY && find(x) == find(y) {
			ans = append(ans, w[x]/w[y])
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}
