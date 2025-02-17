package main

import "slices"

// 输入的视角（逗号选或不选）
func partition(s string) (ans [][]string) {
	check := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	n := len(s)
	var path []string

	// start 表示当前这段回文子串的开始位置
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 不选
		if i < n-1 {
			dfs(i+1, start)
		}

		// 选
		if check(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)             // 下一个子串从 i+1 开始
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0, 0)
	return
}

// 答案的视角（枚举子串结束位置）
func partition2(s string) (ans [][]string) {
	check := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	n := len(s)
	var path []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := i; j < n; j++ { // 枚举子串的结束位置
			if check(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
