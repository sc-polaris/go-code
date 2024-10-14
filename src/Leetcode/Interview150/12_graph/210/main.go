package main

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	g := make([][]int, numCourses)
	in := make([]int, numCourses) // 入度
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
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		ans = append(ans, i)
		for _, j := range g[i] {
			in[j]--
			if in[j] == 0 {
				q = append(q, j)
			}
		}
	}
	if len(ans) == numCourses {
		return ans
	}
	return nil
}
