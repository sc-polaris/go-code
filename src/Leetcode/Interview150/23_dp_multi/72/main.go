package main

/*
	s删除一个元素 dfs(i-1,j)
	t删除一个元素	dfs(i,j-1)

	t添加一个元素相当于 s 删除一个元素 操作数是一样的

	dfs(i, j)，它计算 s[0...i] 和 t[0...j] 的最小编辑距离。
*/

func minDistance(s string, t string) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // 还没计算过
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		// 当 i < 0 时，表示字符串 s 已经为空了。此时，转换操作就是将 t[0...j] 的所有字符插入到 s 中，操作数是 j + 1。
		if i < 0 {
			return j + 1
		}
		// 当 j < 0 时，表示字符串 t 已经为空了。此时，转换操作就是将 s[0...i] 的所有字符删除，操作数是 i + 1。
		if j < 0 {
			return i + 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if s[i] == t[j] {
			return dfs(i-1, j-1)
		}
		return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
	}
	return dfs(n-1, m-1)
}

func minDistance2(s string, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = j
	}
	for i, x := range s {
		f[i+1][0] = i + 1
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}

func minDistance3(s string, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = j
	}
	for _, x := range s {
		pre := f[0]
		f[0]++ // f[0] = i + 1
		for j, y := range t {
			if x == y {
				f[j+1], pre = pre, f[j+1]
			} else {
				f[j+1], pre = min(f[j+1], f[j], pre)+1, f[j+1]
			}
		}
	}
	return f[m]
}
