package main

// 枚举左括号数量
func generateParenthesis(n int) (ans []string) {
	m := n * 2
	var path []byte
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 填左括号
			path = append(path, '(')
			dfs(i+1, open+1)
			path = path[:len(path)-1]
		}
		if i-open < open { // 填右括号
			path = append(path, ')')
			dfs(i+1, open)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}
