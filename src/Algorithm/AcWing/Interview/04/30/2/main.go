package main

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)

	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		return p[j-1] == '.' || p[j-1] == s[i-1]
	}

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] != '*' {
				if matches(i, j) {
					f[i][j] = f[i][j] || f[i-1][j-1]
				}
			} else {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			}
		}
	}

	return f[n][m]
}
