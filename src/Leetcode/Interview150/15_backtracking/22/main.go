package main

import "bytes"

// 方法一：枚举填左括号还是右括号 generateParenthesis
func generateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	// i = 目前填了多少个括号
	// open = 左括号个数，i-open = 右括号个数
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}

// 方法二：枚举下一个左括号的位置
func generateParenthesis2(n int) (ans []string) {
	var path []int
	// i = 目前填了多少个括号
	// balance = 左括号个数 - 右括号个数
	var dfs func(int, int)
	dfs = func(i, balance int) {
		if len(path) == n {
			s := bytes.Repeat([]byte{')'}, n*2)
			for _, j := range path {
				s[j] = '('
			}
			ans = append(ans, string(s))
			return
		}
		// 枚举填 c=0,1,2,...,balance 个右括号
		for c := range balance + 1 {
			// 先填 c 个右括号，然后填 1 个左括号，记录左括号的下标 i+c
			path = append(path, i+c)
			dfs(i+c+1, balance-c+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}
