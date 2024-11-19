package main

/*
	给你一个整数 n 和一个二维整数数组 queries。

	有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。

	queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。

	返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。
*/

// bfs 为避免反复创建 vis 数组，可以在 vis 中保存当前节点是第几次询问访问的。
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	g := make([][]int, n-1)
	for i := range g {
		g[i] = append(g[i], i+1)
	}

	vis := make([]int, n-1)
	bfs := func(i int) int {
		q := []int{0}
		for step := 1; ; step++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				for _, y := range g[x] {
					if y == n-1 {
						return step
					}
					if vis[y] != i {
						vis[y] = i
						q = append(q, y)
					}
				}
			}
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		g[q[0]] = append(g[q[0]], q[1])
		ans[i] = bfs(i + 1)
	}
	return ans
}

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		g[i] = append(g[i], i+1)
	}

	bfs := func(i int) int {
		dist := make([]int, n)
		for i := 1; i < n; i++ {
			dist[i] = -1
		}
		q := []int{0}
		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			for _, y := range g[x] {
				if dist[y] >= 0 {
					continue
				}
				q = append(q, y)
				dist[y] = dist[x] + 1
			}
		}
		return dist[n-1]
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		g[q[0]] = append(g[q[0]], q[1])
		ans[i] = bfs(i + 1)
	}
	return ans
}

/*
	dp
	定义 f[i] 为从 0 到 i 的最短路。

	用 from[i] 记录额外添加的边的终点是 i，起点列表是 from[i]。

	我们可以从 i−1 到 i，也可以从 from[i][j] 到 i，这些位置作为转移来源，用其 f 值 +1 更新 f[i] 的最小值。

	初始值：f[i]=i。

	答案：f[n−1]。
*/

func shortestDistanceAfterQueries3(n int, queries [][]int) []int {
	from := make([][]int, n)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = i
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		from[r] = append(from[r], l)
		if f[l]+1 < f[r] {
			f[r] = f[l] + 1
			for i := r + 1; i < n; i++ {
				f[i] = min(f[i], f[i-1]+1)
				for _, j := range from[i] {
					f[i] = min(f[i], f[j]+1)
				}
			}
		}
		ans[qi] = f[n-1]
	}
	return ans
}
