package main

// dfs(i,j)，表示 s3[:i+j+2] 能否由 s1[:i+1] 和 s 2 [:j+1] 交错组成。
// 其中记号 s[:k] 表示 s 的长为 k 的前缀，即 s[0] 到 s[k−1]。
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}

	// 本题 i 和 j 可以是 −1，为避免下标越界，可以把 dfs(i,j) 记录到 memo[i+1][j+1] 中。
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 && j < 0 {
			return true
		}
		p := &memo[i+1][j+1]
		if *p < 0 {
			if i >= 0 && s1[i] == s3[i+j+1] && dfs(i-1, j) ||
				j >= 0 && s2[j] == s3[i+j+1] && dfs(i, j-1) {
				*p = 1
			} else {
				*p = 0
			}
		}
		return *p == 1
	}
	return dfs(n-1, m-1)
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for j := range m {
		f[0][j+1] = s2[j] == s3[j] && f[0][j]
	}
	for i := range n {
		f[i+1][0] = s1[i] == s3[i] && f[i][0]
		for j := range m {
			f[i+1][j+1] = s1[i] == s3[i+j+1] && f[i][j+1] ||
				s2[j] == s3[i+j+1] && f[i+1][j]
		}
	}
	return f[n][m]
}

func isInterleave3(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}

	f := make([]bool, m+1)
	f[0] = true
	for j := range m {
		f[j+1] = s2[j] == s3[j] && f[j]
	}
	for i := range n {
		f[0] = s1[i] == s3[i] && f[0]
		for j := range m {
			f[j+1] = s1[i] == s3[i+j+1] && f[j+1] ||
				s2[j] == s3[i+j+1] && f[j]
		}
	}
	return f[m]
}
