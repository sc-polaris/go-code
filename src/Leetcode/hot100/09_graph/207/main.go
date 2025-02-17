package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		g[b] = append(g[b], a)
		in[a]++
	}
	var q []int
	for i, x := range in {
		if x == 0 {
			q = append(q, i)
		}
	}
	cnt := 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		cnt++
		for _, j := range g[i] {
			in[j]--
			if in[j] == 0 {
				q = append(q, j)
			}
		}
	}
	return cnt == numCourses
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	colors := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(x int) bool {
		colors[x] = 1 // x 正在访问
		for _, y := range g[x] {
			if colors[y] == 1 || colors[y] == 0 && dfs(y) {
				return true
			}
		}
		colors[x] = 2 // x 完全访问完毕
		return false
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false // 有环
		}
	}
	return true
}
