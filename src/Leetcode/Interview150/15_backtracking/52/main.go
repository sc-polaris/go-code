package main

func totalNQueens(n int) (ans int) {
	col := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			ans++
			return
		}
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diag1[r+c] && !diag2[rc] {
				col[c], diag1[r+c], diag2[rc] = true, true, true
				dfs(r + 1)
				col[c], diag1[r+c], diag2[rc] = false, false, false
			}
		}
	}
	dfs(0)
	return
}
